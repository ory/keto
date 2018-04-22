/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
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
 * @Copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 *
 */

package warden

import (
	"strings"

	"github.com/ory/ladon"
	"github.com/sirupsen/logrus"
)

// AuditLoggerLogrus outputs information about granting or rejecting policies.
type AuditLoggerLogrus struct {
	Logger logrus.FieldLogger
}

func (a *AuditLoggerLogrus) logger() logrus.FieldLogger {
	if a.Logger == nil {
		a.Logger = logrus.New()
	}
	return a.Logger
}

func (a *AuditLoggerLogrus) LogRejectedAccessRequest(r *ladon.Request, p ladon.Policies, d ladon.Policies) {
	if len(d) > 1 {
		allowed := joinPoliciesNames(d[0 : len(d)-1])
		denied := d[len(d)-1].GetID()
		a.logger().
			WithField("action", r.Action).
			WithField("resource", r.Resource).
			WithField("subject", r.Subject).
			WithField("allowed_by", allowed).
			WithField("denied_by", denied).
			Print("Some policies allow this request, but one forcefully rejected it")
	} else if len(d) == 1 {
		denied := d[len(d)-1].GetID()
		a.logger().
			WithField("action", r.Action).
			WithField("resource", r.Resource).
			WithField("subject", r.Subject).
			WithField("denied_by", denied).
			Print("A policy forcefully rejected this request")
	} else {
		a.logger().
			WithField("action", r.Action).
			WithField("resource", r.Resource).
			WithField("subject", r.Subject).
			Print("Because no policy was found for this request, it is rejected")
	}
}

func (a *AuditLoggerLogrus) LogGrantedAccessRequest(r *ladon.Request, p ladon.Policies, d ladon.Policies) {
	a.logger().
		WithField("action", r.Action).
		WithField("resource", r.Resource).
		WithField("subject", r.Subject).
		WithField("allowed_by", joinPoliciesNames(d)).
		Print("One or more policies granted this request.")
}

func joinPoliciesNames(policies ladon.Policies) string {
	var names []string
	for _, policy := range policies {
		names = append(names, policy.GetID())
	}
	return strings.Join(names, ", ")
}
