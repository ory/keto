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
	"strings"

	"github.com/hashicorp/golang-lru"
	"github.com/ory/ladon/compiler"
	"github.com/pkg/errors"
)

func NewRegexpMatcher(size int) *RegexpMatcher {
	if size <= 0 {
		size = 512
	}

	// golang-lru only returns an error if the cache's size is 0. This, we can safely ignore this error.
	cache, _ := lru.New(size)
	return &RegexpMatcher{
		Cache: cache,
	}
}

type RegexpMatcher struct {
	*lru.Cache

	C map[string]*regexp.Regexp
}

func (m *RegexpMatcher) get(pattern string) *regexp.Regexp {
	if val, ok := m.Cache.Get(pattern); !ok {
		return nil
	} else if reg, ok := val.(*regexp.Regexp); !ok {
		return nil
	} else {
		return reg
	}
}

func (m *RegexpMatcher) set(pattern string, reg *regexp.Regexp) {
	m.Cache.Add(pattern, reg)
}

// Matches a needle with an array of regular expressions and returns true if a match was found.
func (m *RegexpMatcher) Matches(p Policy, haystack []string, needle string) (bool, error) {
	var reg *regexp.Regexp
	var err error
	for _, h := range haystack {

		// This means that the current haystack item does not contain a regular expression
		if strings.Count(h, string(p.GetStartDelimiter())) == 0 {
			// If we have a simple string match, we've got a match!
			if h == needle {
				return true, nil
			}

			// Not string match, but also no regexp, continue with next haystack item
			continue
		}

		if reg = m.get(h); reg != nil {
			if reg.MatchString(needle) {
				return true, nil
			}
			continue
		}

		reg, err = compiler.CompileRegex(h, p.GetStartDelimiter(), p.GetEndDelimiter())
		if err != nil {
			return false, errors.WithStack(err)
		}

		m.set(h, reg)
		if reg.MatchString(needle) {
			return true, nil
		}
	}
	return false, nil
}
