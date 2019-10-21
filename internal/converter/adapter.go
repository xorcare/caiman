// Copyright © 2019-2020 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package converter

import "gopkg.in/ldap.v2"

// adaptor an entry representation for go templates.
type adapter struct {
	entry ldap.Entry
}

// Attr returns the first value for the named attribute, or "".
func (a adapter) Attr(key string) string {
	return a.entry.GetAttributeValue(key)
}

// Attrs returns the values for the named attribute, or an empty list.
func (a adapter) Attrs(key string) []string {
	return a.entry.GetAttributeValues(key)
}
