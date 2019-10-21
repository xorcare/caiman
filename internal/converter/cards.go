// Copyright © 2020 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package converter

import (
	"io"

	"github.com/emersion/go-vcard"
)

// Cards its data vCard set with helper methods.
type Cards []vcard.Card

// Encode its encodes data into a vCard text representation.
func (c Cards) Encode(w io.Writer) error {
	encoder := vcard.NewEncoder(w)
	for _, card := range c {
		if err := encoder.Encode(card); err != nil {
			return err
		}
	}

	return nil
}
