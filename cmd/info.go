// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"errors"
	"fmt"

	"github.com/mgutz/ansi"
	"github.com/sethcenterbar/percona-toolkit-tutor/structs"
	"github.com/sethcenterbar/percona-toolkit-tutor/utilities"
	"github.com/spf13/cobra"
)

func listToolInfo(t structs.Tool) string {
	output := ansi.Color("Some information about "+t.Name+":", "cyan+b") + "\n\n"
	output += ansi.Color("        Name: ", "magenta+b") + t.Name + "\n"
	output += ansi.Color("     Summary: ", "magenta+b") + t.Summary + "\n"
	output += ansi.Color("   Situation: ", "magenta+b") + t.Situation + "\n\n"
	return output
}

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "List general info about a tool",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires postional argument for tool name")
		}
		if utilities.ValidateTool(tb, args[0]) {
			return nil
		}
		return fmt.Errorf("invalid toolname specified: %s", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		tl, _ := utilities.GrabTool(tb, args[0])
		println(listToolInfo(tl))
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
