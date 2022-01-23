/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"
	"priceloader/sources"

	"github.com/spf13/cobra"
	// coingecko "github.com/superoo7/go-gecko/v3"
)

// getPriceCmd represents the getPrice command
var getPriceCmd = &cobra.Command{
	Use:   "getPrice",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// httpClient := &http.Client{
		// 	Timeout: time.Second * 10,
		// }
		// // CG := coingecko.NewClient(httpClient)
		// // log.Print(CG.SimplePrice([]string{"fantom"}, []string{"usd"}))
		// // log.Print(CG.SimpleSinglePrice("fantom", "usd"))
		// r, err := httpClient.Get("https://api.coingecko.com/api/v3/simple/token_price/fantom?contract_addresses=0xc1be9a4d5d45beeacae296a7bd5fadbfc14602c4&vs_currencies=usd")
		// if err != nil {
		// 	panic("Karamba")
		// }
		// b, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	panic("Karamba2")
		// }
		// log.Print(string(b))
		// price := map[string]interface{}{}
		// json.Unmarshal(b, &price)
		// log.Print(price)
		s := sources.NewCoinGeckoDataSource()
		f := s.GetFloat64("gton")
		log.Print(f)

	},
}

func init() {
	rootCmd.AddCommand(getPriceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getPriceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getPriceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
