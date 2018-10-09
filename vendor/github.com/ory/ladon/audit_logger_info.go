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
	"log"
	"os"
	"strings"
)

// AuditLoggerInfo outputs information about granting or rejecting policies.
type AuditLoggerInfo struct {
	Logger *log.Logger
}

func (a *AuditLoggerInfo) logger() *log.Logger {
	if a.Logger == nil {
		a.Logger = log.New(os.Stderr, "", log.LstdFlags)
	}
	return a.Logger
}

func (a *AuditLoggerInfo) LogRejectedAccessRequest(r *Request, p Policies, d Policies) {
	if len(d) > 1 {
		allowed := joinPoliciesNames(d[0 : len(d)-1])
		denied := d[len(d)-1].GetID()
		a.logger().Printf("policies %s allow access, but policy %s forcefully denied it", allowed, denied)
	} else if len(d) == 1 {
		denied := d[len(d)-1].GetID()
		a.logger().Printf("policy %s forcefully denied the access", denied)
	} else {
		a.logger().Printf("no policy allowed access")
	}
}

func (a *AuditLoggerInfo) LogGrantedAccessRequest(r *Request, p Policies, d Policies) {
	a.logger().Printf("policies %s allow access", joinPoliciesNames(d))
}

func joinPoliciesNames(policies Policies) string {
	names := []string{}
	for _, policy := range policies {
		names = append(names, policy.GetID())
	}
	return strings.Join(names, ", ")
}
