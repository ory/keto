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

import "strings"

// ResourceContainsCondition is fulfilled if the context matches a substring within the resource name
type ResourceContainsCondition struct{}

// Fulfills returns true if the request's resouce contains the given value string
func (c *ResourceContainsCondition) Fulfills(value interface{}, r *Request) bool {

	filter, ok := value.(map[string]interface{})
	if !ok {
		return false
	}

	valueString, ok := filter["value"].(string)
	if !ok || len(valueString) < 1 {
		return false
	}

	//If no delimiter provided default to "equals" check
	delimiterString, ok := filter["delimiter"].(string)
	if !ok || len(delimiterString) < 1 {
		delimiterString = ""
	}

	// Append delimiter to strings to prevent delim+1 being interpreted as delim+10 being present
	filterValue := delimiterString + valueString + delimiterString
	resourceString := delimiterString + r.Resource + delimiterString

	matches := strings.Contains(resourceString, filterValue)
	return matches

}

// GetName returns the condition's name.
func (c *ResourceContainsCondition) GetName() string {
	return "ResourceContainsCondition"
}
