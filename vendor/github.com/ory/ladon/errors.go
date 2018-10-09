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
	"net/http"

	"github.com/pkg/errors"
)

var (
	// ErrRequestDenied is returned when an access request can not be satisfied by any policy.
	ErrRequestDenied = &errorWithContext{
		error:  errors.New("Request was denied by default"),
		code:   http.StatusForbidden,
		status: http.StatusText(http.StatusForbidden),
		reason: "The request was denied because no matching policy was found.",
	}

	// ErrRequestForcefullyDenied is returned when an access request is explicitly denied by a policy.
	ErrRequestForcefullyDenied = &errorWithContext{
		error:  errors.New("Request was forcefully denied"),
		code:   http.StatusForbidden,
		status: http.StatusText(http.StatusForbidden),
		reason: "The request was denied because a policy denied request.",
	}

	// ErrNotFound is returned when a resource can not be found.
	ErrNotFound = &errorWithContext{
		error:  errors.New("Resource could not be found"),
		code:   http.StatusNotFound,
		status: http.StatusText(http.StatusNotFound),
	}
)

func NewErrResourceNotFound(err error) error {
	if err == nil {
		err = errors.New("not found")
	}

	return errors.WithStack(&errorWithContext{
		error:  err,
		code:   http.StatusNotFound,
		status: http.StatusText(http.StatusNotFound),
		reason: "The requested resource could not be found.",
	})
}

type errorWithContext struct {
	code   int
	reason string
	status string
	error
}

// StatusCode returns the status code of this error.
func (e *errorWithContext) StatusCode() int {
	return e.code
}

// RequestID returns the ID of the request that caused the error, if applicable.
func (e *errorWithContext) RequestID() string {
	return ""
}

// Reason returns the reason for the error, if applicable.
func (e *errorWithContext) Reason() string {
	return e.reason
}

// ID returns the error id, if applicable.
func (e *errorWithContext) Status() string {
	return e.status
}

// Details returns details on the error, if applicable.
func (e *errorWithContext) Details() []map[string]interface{} {
	return []map[string]interface{}{}
}
