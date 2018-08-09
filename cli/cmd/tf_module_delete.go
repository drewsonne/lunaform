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
	"github.com/getlunaform/lunaform/client/modules"
	"github.com/getlunaform/lunaform/models"
	"github.com/spf13/cobra"
)

var tfModuleDeleteIdFlag string

// tfModuleDeleteCmd represents the tfModuleDelete command
var tfModuleDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := lunaformClient.Modules.DeleteModule(
			modules.NewDeleteModuleParams().WithID(tfModuleDeleteIdFlag),
			authHandler,
		)

		if err1, ok := err.(*modules.DeleteModuleUnprocessableEntity); ok {
			handleOutput(cmd, err1.Payload, useHal, nil)
		} else if err != nil {
			handleOutput(cmd, nil, useHal, err)
		} else {
			handleOutput(cmd, models.StringHalResponse("Successfully deleted"), useHal, err)
		}
	},
}

func init() {
	tfModuleDeleteCmd.Flags().StringVar(&tfModuleDeleteIdFlag, "id", "",
		"ID of the terraform module in lunaform")
	tfModuleDeleteCmd.MarkFlagRequired("id")
}
