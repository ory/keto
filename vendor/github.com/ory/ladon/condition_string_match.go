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
	"regexp"
)

// StringMatchCondition is a condition which is fulfilled if the given
// string value matches the regex pattern specified in StringMatchCondition
type StringMatchCondition struct {
	Matches string `json:"matches"`
}

// Fulfills returns true if the given value is a string and matches the regex
// pattern in StringMatchCondition.Matches
func (c *StringMatchCondition) Fulfills(value interface{}, _ *Request) bool {
	s, ok := value.(string)

	matches, _ := regexp.MatchString(c.Matches, s)

	return ok && matches
}

// GetName returns the condition's name.
func (c *StringMatchCondition) GetName() string {
	return "StringMatchCondition"
}
