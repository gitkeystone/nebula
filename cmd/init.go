/*
Copyright © 2025 Chen Xiaohui <ysucxh@163.com>

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
	"github.com/gitkeystone/nebula/cmd/container"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use: "init",
	// Short: "A brief description of your command",
	Long: `Init container process run user's process in container.
Do not call it outside.`,

	Hidden: true,

	Run: func(cmd *cobra.Command, args []string) {
		log.Info("init come on")

		err := container.RunContainerInitProcess()
		log.Error("init container process", zap.Error(err))
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
