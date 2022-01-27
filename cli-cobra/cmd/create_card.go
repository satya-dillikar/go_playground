/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// cardCmd represents the card command
var createCardCmd = &cobra.Command{
	Use:   "card <name>",
	Short: "Create cards.",
	Long: ` This command creates cards. Example:
	greetctl create card eva -n="Eva Green" -o=thanksgiving -l=fr
	greetctl create card bob --name="Bob Marley" --occasion=birthday`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create_card called")
		fmt.Printf("%v\n", args)
		fmt.Println("Here are the arguments of card command : " + strings.Join(args, ","))
		name, _ := cmd.Flags().GetString("name")
		language, _ := cmd.Flags().GetString("language")
		occasion, _ := cmd.Flags().GetString("occasion")
		fmt.Println("value of the flag name :" + name)
		fmt.Println("value of the flag language :" + language)
		fmt.Println("value of the flag occasion :" + occasion)
	},
}

func init() {
	createCmd.AddCommand(createCardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	createCardCmd.PersistentFlags().StringP("occasion", "o", "", "Possible values: newyear, thanksgiving, birthday")
	createCardCmd.PersistentFlags().StringP("language", "l", "en", "Possible values: en, fr")
	createCardCmd.PersistentFlags().StringP("name", "n", "", "Name of the user to whom you want to greet")
	createCardCmd.MarkPersistentFlagRequired("name")
	createCardCmd.MarkPersistentFlagRequired("occasion")

}
