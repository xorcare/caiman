// Copyright © 2019-2020 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"io"

	"github.com/emersion/go-vcard"
	"gopkg.in/yaml.v2"
)

var _default = Config{
	Fields: map[string][]Field{
		vcard.FieldEmail: {
			{
				Template: `{{.entry.Attr "mail"}}`,
				Params: vcard.Params{
					vcard.ParamType:      []string{"INTERNET", vcard.TypeWork},
					vcard.ParamPreferred: []string{"1"},
				},
			},
		},
		vcard.FieldName: {
			{
				Template: `{{.entry.Attr "sn"}}{{$givenName := .entry.Attr "givenName"}}{{if ne $givenName ""}};{{$givenName}}{{end}}`,
			},
		},
		vcard.FieldTelephone: {
			{
				Template: `{{.entry.Attr "mobile"}}`,
				Params: vcard.Params{
					vcard.ParamType: []string{
						vcard.TypeText, vcard.TypeVoice, vcard.TypeHome,
					},
					vcard.ParamPreferred: []string{"1"},
				},
			},
			{
				Template: `{{.entry.Attr "telephoneNumber"}}`,
				Params: vcard.Params{
					vcard.ParamType: []string{vcard.TypeWork},
				},
			},
		},
		vcard.FieldTitle: {
			{
				Template: `{{.entry.Attr "title"}}`,
			},
		},
		vcard.FieldOrganization: {
			{
				Template: `{{if .entry.Attr "department" }};{{.entry.Attr "department"}}{{else}}{{end}}`,
			},
		},
		vcard.FieldPhoto: {
			{
				Template: `{{.entry.Attr "jpegPhoto" | Base64Encode}}`,
				Params: vcard.Params{
					"ENCODING":           []string{"b"},
					vcard.ParamType:      []string{"JPEG"},
					vcard.ParamPreferred: []string{"1"},
				},
			},
			{
				Template: `{{.entry.Attr "photo" | Base64Encode}}`,
				Params: vcard.Params{
					"ENCODING":           []string{"b"},
					vcard.ParamPreferred: []string{"2"},
				},
			},
			{
				Template: `{{.entry.Attr "thumbnailPhoto" | Base64Encode}}`,
				Params: vcard.Params{
					"ENCODING": []string{"b"},
				},
			},
		},
	},
}

// Config the basic structure configuration.
type Config struct {
	Fields map[string][]Field
}

// Decode the YAML settings from reader.
func (c *Config) Decode(r io.Reader) error {
	return yaml.NewDecoder(r).Decode(c)
}

// Encode the YAML settings to writer.
func (c Config) Encode(w io.Writer) error {
	return yaml.NewEncoder(w).Encode(c)
}

// Field a field contains a value and some parameters.
type Field struct {
	Group string
	// Params is a set of field parameters.
	Params vcard.Params
	// Template golang text template string.
	Template string
}

// Default returns copy of default config.
func Default() Config {
	return _default
}
