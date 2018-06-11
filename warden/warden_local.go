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
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package warden

import (
	"context"

	"github.com/ory/fosite"
	"github.com/ory/keto/role"
	"github.com/ory/ladon"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func NewWarden(
	warden ladon.Warden,
	roles role.Manager,
	l logrus.FieldLogger) *Warden {
	return &Warden{
		Warden: warden,
		Roles:  roles,
		L:      l,
	}
}

type Warden struct {
	Warden ladon.Warden
	Roles  role.Manager
	L      logrus.FieldLogger
}

func (w *Warden) IsAllowed(ctx context.Context, a *AccessRequest) error {
	if err := w.isAllowed(ctx, &ladon.Request{
		Resource: a.Resource,
		Action:   a.Action,
		Subject:  a.Subject,
		Context:  a.Context,
	}); err != nil {
		w.L.WithFields(logrus.Fields{
			"subject": a.Subject,
			"request": a,
			"reason":  "The policy decision point denied the request",
		}).WithError(err).Infof("Access denied")
		return err
	}

	w.L.WithFields(logrus.Fields{
		"subject": a.Subject,
		"request": a,
		"reason":  "The policy decision point allowed the request",
	}).Infof("Access allowed")
	return nil
}

func (w *Warden) isAllowed(ctx context.Context, a *ladon.Request) error {
	roles, err := w.Roles.FindRolesByMember(a.Subject, 10000, 0)
	if err != nil {
		return err
	}

	errs := make([]error, len(roles)+1)
	errs[0] = w.Warden.IsAllowed(&ladon.Request{
		Resource: a.Resource,
		Action:   a.Action,
		Subject:  a.Subject,
		Context:  a.Context,
	})

	for k, g := range roles {
		errs[k+1] = w.Warden.IsAllowed(&ladon.Request{
			Resource: a.Resource,
			Action:   a.Action,
			Subject:  g.ID,
			Context:  a.Context,
		})
	}

	for _, err := range errs {
		if errors.Cause(err) == ladon.ErrRequestForcefullyDenied {
			return errors.Wrap(fosite.ErrRequestForbidden, err.Error())
		}
	}

	// If no one explicitly denies the access request (e.g. some group), it's ok to return with "access granted"
	// if at least one of the decisions is positive (no error)
	for _, err := range errs {
		if err == nil {
			return nil
		}
	}

	return errors.Wrap(fosite.ErrRequestForbidden, ladon.ErrRequestDenied.Error())
}
