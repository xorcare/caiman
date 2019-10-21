// Copyright © 2019-2020 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package converter

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"strings"
	"sync"
	"text/template"

	"github.com/emersion/go-vcard"
	"github.com/go-ldap/ldif"
	"github.com/spf13/cast"

	"github.com/xorcare/caiman/internal/config"
)

var _funcMap = map[string]interface{}{
	"Base64Encode": base64Encode,
}

// Converter structure of the converter.
type Converter struct {
	Config    config.Config
	templates map[string]*template.Template
	once      sync.Once
}

// LDIF2vCARD4 converts LDIF to vCard 4.
func (c *Converter) LDIF2vCARD4(l ldif.LDIF) (cards Cards, err error) {
	skipped, successful, nullable := 0, 0, 0
	defer func() {
		log.Println(fmt.Sprintf("total entries %d", len(l.Entries)))
		log.Println(fmt.Sprintf("skipped %d entries because it is nil", nullable))
		log.Println(fmt.Sprintf("skipped %d entries because bad count of filled fields", skipped))
		log.Println(fmt.Sprintf("successfully processed %d entries", successful))
	}()
	for _, v := range l.Entries {
		if v == nil || v.Entry == nil {
			nullable++
			continue
		}
		entry := adapter{entry: *v.Entry}

		card := vcard.Card{}
		for fieldName, fields := range c.Config.Fields {
			fieldName = strings.ToUpper(fieldName)
			for _, field := range fields {
				tpl, err := c.parseTemplate(field.Template)
				if err != nil {
					return nil, err
				}
				buf := bytes.NewBuffer([]byte{})
				err = tpl.Execute(buf, map[string]interface{}{
					"entry": entry,
				})
				if err != nil {
					return nil, err
				}
				add(card, fieldName, vcard.Field{
					Value:  buf.String(),
					Params: field.Params,
					Group:  field.Group,
				})
			}
		}
		if len(card) <= 3 {
			skipped++
			continue
		}
		vcard.ToV4(card)
		cards = append(cards, card)
		successful++
	}

	return cards, nil
}

func (c *Converter) parseTemplate(tpl string) (tmpl *template.Template, err error) {
	c.once.Do(func() {
		c.templates = make(map[string]*template.Template)
	})

	if tmpl, ok := c.templates[tpl]; ok {
		return tmpl.Clone()
	}

	tmpl, err = template.New("").Funcs(_funcMap).Parse(tpl)
	if err != nil {
		return nil, err
	}

	c.templates[tpl] = tmpl

	return tmpl.Clone()
}

func add(card vcard.Card, k string, f vcard.Field) {
	if k != "" && strings.TrimSpace(f.Value) != "" {
		card.Add(k, &f)
	}
}

func base64Encode(content interface{}) (string, error) {
	str, err := cast.ToStringE(content)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString([]byte(str)), nil
}
