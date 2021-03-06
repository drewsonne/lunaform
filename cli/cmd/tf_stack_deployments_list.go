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
	"github.com/getlunaform/lunaform/client/stacks"
	"github.com/spf13/cobra"
)

var tfStackDeploymentsListIdFlag string

// tfStackDeploymentsListCmd represents the tfStackDeploymentsList command
var tfStackDeploymentsListCmd = &cobra.Command{
	Use:   "list-deployments",
	Short: "List deployments for a given stack",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if deployments, err := lunaformClient.Stacks.ListDeployments(
			stacks.NewListDeploymentsParams().WithID(tfStackDeploymentsListIdFlag),
			authHandler,
		); err == nil {
			handleOutput(cmd, deployments.Payload, useHal, err)
		} else {
			handleOutput(cmd, nil, useHal, err)
		}
	},
}

func init() {
	tfStackDeploymentsListCmd.Flags().StringVar(&tfStackDeploymentsListIdFlag, "id", "",
		"ID of the stack in lunaform")
	tfStackDeploymentsListCmd.MarkFlagRequired("id")
}
