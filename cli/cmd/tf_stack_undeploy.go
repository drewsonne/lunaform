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
	"github.com/getlunaform/lunaform/client/stacks"
	"github.com/spf13/cobra"
)

var (
	tfStackDestroyIdFlag  string
	tfStackDestroyIdsFlag []string
)

// tfStackUndeployCmd represents the tfStackUndeploy command
var tfStackDestroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "destroy the deployed stack",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		var err error

		if tfStackDestroyIdFlag == "" && tfStackDestroyIdsFlag == nil {
			handlerError(fmt.Errorf("`--id` or `--ids` must be provided"))
		} else if tfStackDestroyIdFlag != "" {
			tfStackDestroyIdsFlag = append(tfStackDestroyIdsFlag, tfStackDestroyIdFlag)
		}

		for _, id := range tfStackDestroyIdsFlag {
			if _, err = lunaformClient.Stacks.UndeployStack(
				stacks.NewUndeployStackParams().WithID(id),
				authHandler,
			); err != nil {
				break
			}
		}

		if err != nil {
			handleOutput(cmd, nil, useHal, err)
		}
	},
}

func init() {
	flags := tfStackDestroyCmd.Flags()
	flags.StringVar(&tfStackDestroyIdFlag, "id", "",
		"ID of the terraform module in lunaform")
	flags.StringSliceVar(&tfStackDestroyIdsFlag, "ids", []string{},
		"List of comma separated IDs for lunaform stacks")

}
