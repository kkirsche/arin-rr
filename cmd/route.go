// Copyright © 2016 Kevin Kirsche <kev.kirsche@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"html/template"
	"os"

	"github.com/kkirsche/arin-rr/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// routeCmd represents the arin command
var routeCmd = &cobra.Command{
	Use:   "route",
	Short: "For modifying entries in ARIN's Internet Route Registry",
	Long:  `Use to modify ARIN's route registry.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := irr.NewLogger(viper.GetBool("verbose"))
		logger.Verboseln("Creating email structure.")
		email := irr.NewEmail(
			viper.GetString("email.from"),
			viper.GetString("email.to"),
			viper.GetString("email.subject"),
			viper.GetString("email.smtp"),
		)

		logger.Verboseln("Creating route registry entry structure.")
		entry := irr.NewRouteRegistryEntry(
			viper.GetString("arin.route"),
			viper.GetString("arin.description"),
			viper.GetInt("arin.asn"),
			viper.GetString("arin.notify-email"),
			viper.GetString("arin.maintained-by"),
			viper.GetString("arin.changed-email"),
			viper.GetString("arin.source"),
		)

		logger.Verboseln("Flattening structure.")
		forArin := irr.NewARINRouteEntry(email, entry)
		flatArinRouteEntry := forArin.Flatten()

		logger.Verboseln("Validating ASN entry.")
		if entry.ASN == 0 {
			logger.Println("ASN cannot be 0. Please set the ASN with --asn or -a and try again.")
			return
		}

		logger.Verboseln("Parsing output template.")
		t := template.Must(template.New("ArinRouteEntry").Parse(irr.ArinRouteEntryTemplate))
		logger.Verboseln("Outputting template to stdout.")
		err := t.Execute(os.Stdout, flatArinRouteEntry)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if viper.GetBool("general.submit") {
			logger.Verboseln("Send email enabled. Connecting to SMTP server")
			err := irr.SendArinEmail(flatArinRouteEntry, t)
			if err != nil {
				logger.Println(err.Error())
				return
			}
		} else {
			logger.Verboseln("Send email disabled.")
		}
	},
}

func init() {
	RootCmd.AddCommand(routeCmd)

	// Email Related
	RootCmd.PersistentFlags().StringP("from", "f", "", "The email address to use as the 'from' email address.")
	viper.BindPFlag("email.from", RootCmd.PersistentFlags().Lookup("from"))

	RootCmd.PersistentFlags().StringP("to", "t", irr.DefaultToEmail, "The email address to use as the 'to' email address.")
	viper.BindPFlag("email.to", RootCmd.PersistentFlags().Lookup("to"))

	RootCmd.PersistentFlags().StringP("subject", "u", irr.DefaultSubject, "The subject line to use in the email.")
	viper.BindPFlag("email.subject", RootCmd.PersistentFlags().Lookup("subject"))

	RootCmd.PersistentFlags().StringP("smtp", "p", "", "The smtp server and port (server:port) to use when sending the email.")
	viper.BindPFlag("email.smtp", RootCmd.PersistentFlags().Lookup("smtp"))

	// The Route Registry Entry
	RootCmd.PersistentFlags().StringP("route", "r", "", "The route which should be added to ARIN's route registry.")
	viper.BindPFlag("arin.route", RootCmd.PersistentFlags().Lookup("route"))

	RootCmd.PersistentFlags().StringP("desc", "d", "", "The route description.")
	viper.BindPFlag("arin.description", RootCmd.PersistentFlags().Lookup("desc"))

	RootCmd.PersistentFlags().IntP("asn", "a", 0, "The numeric ASN number whic the route will be advertised from.")
	viper.BindPFlag("arin.asn", RootCmd.PersistentFlags().Lookup("asn"))

	RootCmd.PersistentFlags().StringP("notify", "n", "", "Email address to be notified of ARIN's response.")
	viper.BindPFlag("arin.notify-email", RootCmd.PersistentFlags().Lookup("notify"))

	RootCmd.PersistentFlags().StringP("maintained", "m", "", "Who should be noted as the maintainer")
	viper.BindPFlag("arin.maintained-by", RootCmd.PersistentFlags().Lookup("maintained"))

	RootCmd.PersistentFlags().StringP("changed", "g", "", "Email address of who changed the entry")
	viper.BindPFlag("arin.changed-email", RootCmd.PersistentFlags().Lookup("changed"))

	RootCmd.PersistentFlags().StringP("source", "o", irr.DefaultSource, "The source to use for the route registry entry")
	viper.BindPFlag("arin.source", RootCmd.PersistentFlags().Lookup("source"))

	// General
	RootCmd.PersistentFlags().BoolP("submit", "s", false, "Submit route announcement via email to ARIN")
	viper.BindPFlag("general.submit", RootCmd.PersistentFlags().Lookup("submit"))
}
