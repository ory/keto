/*
 * Copyright © 2016-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package ladon

import (
	"github.com/pkg/errors"
)

// Ladon is an implementation of Warden.
type Ladon struct {
	Manager     Manager
	Matcher     matcher
	AuditLogger AuditLogger
}

func (l *Ladon) matcher() matcher {
	if l.Matcher == nil {
		l.Matcher = DefaultMatcher
	}
	return l.Matcher
}

func (l *Ladon) auditLogger() AuditLogger {
	if l.AuditLogger == nil {
		l.AuditLogger = DefaultAuditLogger
	}
	return l.AuditLogger
}

// IsAllowed returns nil if subject s has permission p on resource r with context c or an error otherwise.
func (l *Ladon) IsAllowed(r *Request) (err error) {
	policies, err := l.Manager.FindRequestCandidates(r)
	if err != nil {
		return err
	}

	// Although the manager is responsible of matching the policies, it might decide to just scan for
	// subjects, it might return all policies, or it might have a different pattern matching than Golang.
	// Thus, we need to make sure that we actually matched the right policies.
	return l.DoPoliciesAllow(r, policies)
}

// DoPoliciesAllow returns nil if subject s has permission p on resource r with context c for a given policy list or an error otherwise.
// The IsAllowed interface should be preferred since it uses the manager directly. This is a lower level interface for when you don't want to use the ladon manager.
func (l *Ladon) DoPoliciesAllow(r *Request, policies []Policy) (err error) {
	var allowed = false
	var deciders = Policies{}

	// Iterate through all policies
	for _, p := range policies {
		// Does the action match with one of the policies?
		// This is the first check because usually actions are a superset of get|update|delete|set
		// and thus match faster.
		if pm, err := l.matcher().Matches(p, p.GetActions(), r.Action); err != nil {
			return errors.WithStack(err)
		} else if !pm {
			// no, continue to next policy
			continue
		}

		// Does the subject match with one of the policies?
		// There are usually less subjects than resources which is why this is checked
		// before checking for resources.
		if sm, err := l.matcher().Matches(p, p.GetSubjects(), r.Subject); err != nil {
			return err
		} else if !sm {
			// no, continue to next policy
			continue
		}

		// Does the resource match with one of the policies?
		if rm, err := l.matcher().Matches(p, p.GetResources(), r.Resource); err != nil {
			return errors.WithStack(err)
		} else if !rm {
			// no, continue to next policy
			continue
		}

		// Are the policies conditions met?
		// This is checked first because it usually has a small complexity.
		if !l.passesConditions(p, r) {
			// no, continue to next policy
			continue
		}

		// Is the policies effect deny? If yes, this overrides all allow policies -> access denied.
		if !p.AllowAccess() {
			deciders = append(deciders, p)
			l.auditLogger().LogRejectedAccessRequest(r, policies, deciders)
			return errors.WithStack(ErrRequestForcefullyDenied)
		}

		allowed = true
		deciders = append(deciders, p)
	}

	if !allowed {
		l.auditLogger().LogRejectedAccessRequest(r, policies, deciders)
		return errors.WithStack(ErrRequestDenied)
	}

	l.auditLogger().LogGrantedAccessRequest(r, policies, deciders)
	return nil
}

func (l *Ladon) passesConditions(p Policy, r *Request) bool {
	for key, condition := range p.GetConditions() {
		if pass := condition.Fulfills(r.Context[key], r); !pass {
			return false
		}
	}
	return true
}
