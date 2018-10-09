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

// StringPairsEqualCondition is a condition which is fulfilled if the given
// array of pairs contains two-element string arrays where both elements
// in the string array are equal
type StringPairsEqualCondition struct{}

// Fulfills returns true if the given value is an array of string arrays and
// each string array has exactly two values which are equal
func (c *StringPairsEqualCondition) Fulfills(value interface{}, _ *Request) bool {
	pairs, PairsOk := value.([]interface{})
	if !PairsOk {
		return false
	}

	for _, v := range pairs {
		pair, PairOk := v.([]interface{})
		if !PairOk || (len(pair) != 2) {
			return false
		}

		a, AOk := pair[0].(string)
		b, BOk := pair[1].(string)

		if !AOk || !BOk || (a != b) {
			return false
		}
	}

	return true
}

// GetName returns the condition's name.
func (c *StringPairsEqualCondition) GetName() string {
	return "StringPairsEqualCondition"
}
