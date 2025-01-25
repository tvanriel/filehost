/*
Copyright © 2025 Ted van Riel <80752652+tvanriel@users.noreply.github.com>
Licensed under the EUPL
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tvanriel/filehost/pkg/app"
)

// rootCmd represents the base command when called without any subcommands.
//
//nolint:gochecknoglobals // Cobra convention.
var rootCmd = &cobra.Command{
	Use:   "filehost",
	Short: "Filehost allows your users to upload files to S3 and hosts them :D",
	Run: func(cmd *cobra.Command, args []string) {
		app.Run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
