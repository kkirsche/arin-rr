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
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/kkirsche/arin-rr/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var count = 1

var ipBlockAdditionRate = map[int]int{
	8:  0,
	9:  128,
	10: 64,
	11: 32,
	12: 16,
	13: 8,
	14: 4,
	15: 2,
	16: 0,
	17: 128,
	18: 64,
	19: 32,
	20: 16,
	21: 8,
	22: 4,
	23: 2,
	24: 0,
	25: 128,
	26: 64,
	27: 32,
	28: 16,
	29: 8,
	30: 4,
	31: 2,
	32: 0,
}

// withinCmd represents the routes-within command
var withinCmd = &cobra.Command{
	Use:   "routes-within",
	Short: "For modifying multiple entries in ARIN's Internet Route Registry",
	Long: `For modifying multiple entries in ARIN's Internet Route Registry For example:

	irr.git arin routes-within -v -a 12345 -r "1.2.0.0/16"

This will print out what would be sent to ARIN with ASN 12345 and will submit
all subnets including the /16 itself for a total of 255 possible submissions.
To submit to ARIN, please use the --submit or -s flag.
`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := irr.NewLogger(verbose)

		logger.Verboseln("Parsing provided CIDR block.")
		ip, ipnet, err := net.ParseCIDR(viper.GetString("arin.route"))
		if err != nil {
			logger.Printfln("Failed to parse CIDR route block with error %s", err.Error())
			return
		}

		ipnetSize, ipnetPossible := ipnet.Mask.Size()
		if ipnetSize < 16 {
			logger.Println("Tool is limited to /16's and smaller for safety reasons.")
			return
		}

		switch ipnetPossible {
		case 32:
			logger.Verbosef("Detected IPv4: %s, Mask: %d of %d, Network: %s", ip.String(), ipnetSize, ipnetPossible, ipnet.Network())
		case 128:
			logger.Verbosef("Detected IPv6: %s, Mask: %d of %d, Network: %s", ip.String(), ipnetSize, ipnetPossible, ipnet.Network())
		default:
			logger.Verbosef("Unknown address size shared. Please try again.")
			return
		}

		var ipIntSections []int
		for ipSize := ipnetSize; ipSize <= 24 && ipSize >= 16; ipSize++ {
			logger.Verbosef("Iterating over IP Space. Beginning with /%d", ipSize)
			ipStringSize := strconv.Itoa(ipSize)
			ipSections := strings.Split(ip.String(), ".")
			var ipInt int
			for _, val := range ipSections {
				ipInt, err = strconv.Atoi(val)
				if err != nil {
					logger.Printfln("Failed to convert IP octet with error: %s", err.Error())
					return
				}
				ipIntSections = append(ipIntSections, ipInt)
				if err != nil {
					logger.Printfln("Failed to extract IP address for updating with error: %s.", err.Error())
					return
				}
			}
			additionRate := ipBlockAdditionRate[ipSize]

			if ipIntSections[2] > 255 {
				ipIntSections[2] = 0
			}

			for ipIntSections[2] <= 255 {
				thirdSection := strconv.Itoa(ipIntSections[2])
				ipSections[2] = thirdSection
				logger.Verboseln("Creating email structure.")
				email := irr.NewEmail(
					viper.GetString("email.from"),
					viper.GetString("email.to"),
					viper.GetString("email.subject"),
					viper.GetString("email.smtp"),
				)

				currentRoute := strings.Join(ipSections, ".") + "/" + ipStringSize
				logger.Printfln("%s", "\n")

				logger.Verboseln("Creating route registry entry structure.")
				entry := irr.NewRouteRegistryEntry(
					currentRoute,
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
				logger.Verboseln("Outputting template to stdout.\n")
				err = t.Execute(os.Stdout, flatArinRouteEntry)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				logger.Verboseln("\n")

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

				if additionRate == 0 {
					ipIntSections[2] += 256
				} else {
					ipIntSections[2] += additionRate
				}
			}
		}
	},
}

func init() {
	routeCmd.AddCommand(withinCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// withinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
