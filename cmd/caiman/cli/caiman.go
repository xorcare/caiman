// Copyright © 2019-2020 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cli

import (
	"errors"
	"log"
	"os"

	"github.com/go-ldap/ldif"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/xorcare/caiman/internal/config"
	"github.com/xorcare/caiman/internal/converter"
)

var _config = config.Default()
var _configDump = false
var _configFile = ""

var _caimanCMD = &cobra.Command{
	Use:  "caiman",
	Long: "caiman is a tool to convert people data from LDAP(LDIF) format to vCard4 contact format",
	Example: "caiman < person.ldif > person.vcf" +
		"\ncaiman --config-file ~/.caiman.yaml < person.ldif > person.vcf" +
		"\ncaiman --config-dump > .caiman.yaml" +
		"\ncat person.ldif | caiman > person.vcf" +
		"\ncat person.ldif | caiman | tee person.vcf",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		log.SetOutput(cmd.ErrOrStderr())
		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		if _configFile == "" {
			return
		}
		file, err := os.Open(_configFile)
		if os.IsNotExist(err) {
			return nil
		} else if err != nil {
			return err
		} else {
			defer file.Close()
		}
		return _config.Decode(file)
	},
	RunE:         exec,
	SilenceUsage: true,
}

func init() {
	_caimanCMD.Flags().StringVarP(
		&_configFile, "config-file", "f", _configFile,
		"the settings file from which the settings will be loaded")
	_caimanCMD.Flags().BoolVarP(
		&_configDump, "config-dump", "d", _configDump,
		"print to standard output all configuration values,"+
			" it prints configuration data in YAML format")

}

func exec(cmd *cobra.Command, _ []string) error {
	if _configDump {
		return _config.Encode(cmd.OutOrStdout())
	}

	if terminal.IsTerminal(0) {
		return errors.New("no piped data")
	}

	dif := ldif.LDIF{}
	if err := ldif.Unmarshal(cmd.InOrStdin(), &dif); err != nil {
		return err
	}

	conv := converter.Converter{Config: _config}
	cards, err := conv.LDIF2vCARD4(dif)
	if err != nil {
		return err
	}

	if err := cards.Encode(cmd.OutOrStdout()); err != nil {
		return err
	}

	return nil
}

// Execute starts command line processing.
func Execute() {
	if err := _caimanCMD.Execute(); err != nil {
		os.Exit(1)
	}
}
