// Copyright © 2019-2020 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/xorcare/caiman/cmd/caiman/cli"
)

// Useful links:

// RFC	2849	https://tools.ietf.org/html/rfc2849	The LDAP Data Interchange Format (LDIF) - Technical Specification
// RFC	4525	https://tools.ietf.org/html/rfc4525	Lightweight Directory Access Protocol (LDAP)
// RFC	6350	https://tools.ietf.org/html/rfc6350	vCard Format Specification
// RFC	6351	https://tools.ietf.org/html/rfc6351	xCard: vCard XML Representation
// RFC	7095	https://tools.ietf.org/html/rfc7095	jCard: The JSON Format for vCard

// https://en.wikipedia.org/wiki/VCard#vCard_4.0				vCard
// https://en.wikipedia.org/wiki/LDAP_Data_Interchange_Format	LDAP Data Interchange Format

func main() {
	cli.Execute()
}
