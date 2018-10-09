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
	"net"
)

// CIDRCondition makes sure that the warden requests' IP address is in the given CIDR.
type CIDRCondition struct {
	CIDR string `json:"cidr"`
}

// Fulfills returns true if the the request is fulfilled by the condition.
func (c *CIDRCondition) Fulfills(value interface{}, _ *Request) bool {
	ips, ok := value.(string)
	if !ok {
		return false
	}

	_, cidrnet, err := net.ParseCIDR(c.CIDR)
	if err != nil {
		return false
	}

	ip := net.ParseIP(ips)
	if ip == nil {
		return false
	}

	return cidrnet.Contains(ip)
}

// GetName returns the condition's name.
func (c *CIDRCondition) GetName() string {
	return "CIDRCondition"
}
