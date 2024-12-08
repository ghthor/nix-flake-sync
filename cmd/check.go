// Copyright (c) Will Drengwitz <ghthor@gmail.com>
// SPDX-License-Identifier: MIT

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "TODO",
	Long: `
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("check called")
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
