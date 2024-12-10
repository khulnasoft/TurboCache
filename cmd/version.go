package cmd

import (
	"fmt"

	"github.com/khulnasoft/turbocache/pkg/turbocache"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of this turbocache build",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(turbocache.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
