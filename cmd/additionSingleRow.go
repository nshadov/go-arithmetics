// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

// additionSingleRowCmd represents the additionSingleRow command
var additionSingleRowCmd = &cobra.Command{
	Use:   "additionSingleRow",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, i := range getRandomSingleRow(6) {
			fmt.Printf("%5d\n", i)
		}

		fmt.Println("-----")
	},
}

func getRandomSingleRow(n int) []int {
	var out = make([]int, n)

	for i := 0; i < n; i++ {
		out[i] = rand.Intn(9) + 1
	}

	return out
}

func init() {
	rootCmd.AddCommand(additionSingleRowCmd)

	rand.Seed(time.Now().UTC().UnixNano())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// additionSingleRowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// additionSingleRowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
