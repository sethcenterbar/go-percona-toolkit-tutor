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
	"strconv"

	"github.com/mgutz/ansi"
	"github.com/sethcenterbar/percona-toolkit-tutor/structs"
	"github.com/sethcenterbar/percona-toolkit-tutor/utilities"
	"github.com/spf13/cobra"
)

func listVideos(t structs.Tool) string {
	output := ansi.Color("Some relevant videos for learning "+t.Name+":", "cyan+b") + "\n"
	for i, video := range t.Videos {
		output += "\n" + ansi.Color("Video "+strconv.Itoa(i+1), "yellow+b") + "\n"
		output += ansi.Color("Description: ", "magenta+b") + video.Description + "\n"
		output += ansi.Color("       Link: ", "magenta+b") + video.Link + "\n"
	}
	return output + "\n"
}

// videosCmd represents the videos command
var videosCmd = &cobra.Command{
	Use:   "videos",
	Short: "List videos associated with a tool",
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
		println(listVideos(tl))
	},
}

func init() {
	rootCmd.AddCommand(videosCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// videosCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// videosCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
