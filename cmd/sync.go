// Copyright (c) Will Drengwitz <ghthor@gmail.com>
// SPDX-License-Identifier: MIT

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "TODO",
	Long: `
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sync called")
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
