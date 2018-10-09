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
	"encoding/json"

	"github.com/pkg/errors"
)

// Condition either do or do not fulfill an access request.
type Condition interface {
	// GetName returns the condition's name.
	GetName() string

	// Fulfills returns true if the request is fulfilled by the condition.
	Fulfills(interface{}, *Request) bool
}

// Conditions is a collection of conditions.
type Conditions map[string]Condition

// AddCondition adds a condition to the collection.
func (cs Conditions) AddCondition(key string, c Condition) {
	cs[key] = c
}

// MarshalJSON marshals a list of conditions to json.
func (cs Conditions) MarshalJSON() ([]byte, error) {
	out := make(map[string]*jsonCondition, len(cs))
	for k, c := range cs {
		raw, err := json.Marshal(c)
		if err != nil {
			return []byte{}, errors.WithStack(err)
		}

		out[k] = &jsonCondition{
			Type:    c.GetName(),
			Options: json.RawMessage(raw),
		}
	}

	return json.Marshal(out)
}

// UnmarshalJSON unmarshals a list of conditions from json.
func (cs Conditions) UnmarshalJSON(data []byte) error {
	if cs == nil {
		return errors.New("Can not be nil")
	}

	var jcs map[string]jsonCondition
	var dc Condition

	if err := json.Unmarshal(data, &jcs); err != nil {
		return errors.WithStack(err)
	}

	for k, jc := range jcs {
		var found bool
		for name, c := range ConditionFactories {
			if name == jc.Type {
				found = true
				dc = c()

				if len(jc.Options) == 0 {
					cs[k] = dc
					break
				}

				if err := json.Unmarshal(jc.Options, dc); err != nil {
					return errors.WithStack(err)
				}

				cs[k] = dc
				break
			}
		}

		if !found {
			return errors.Errorf("Could not find condition type %s", jc.Type)
		}
	}

	return nil
}

type jsonCondition struct {
	Type    string          `json:"type"`
	Options json.RawMessage `json:"options"`
}

// ConditionFactories is where you can add custom conditions
var ConditionFactories = map[string]func() Condition{
	new(StringEqualCondition).GetName(): func() Condition {
		return new(StringEqualCondition)
	},
	new(CIDRCondition).GetName(): func() Condition {
		return new(CIDRCondition)
	},
	new(EqualsSubjectCondition).GetName(): func() Condition {
		return new(EqualsSubjectCondition)
	},
	new(StringPairsEqualCondition).GetName(): func() Condition {
		return new(StringPairsEqualCondition)
	},
	new(StringMatchCondition).GetName(): func() Condition {
		return new(StringMatchCondition)
	},
	new(ResourceContainsCondition).GetName(): func() Condition {
		return new(ResourceContainsCondition)
	},
}
