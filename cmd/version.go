package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version is semantic version of this app.
const Version string = "0.1.0"

// Revision is the revision of this app's binary.
var Revision = "unset"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `All software has versions. This is colout's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("colout: A CLI tool that makes it easy to use colored output.\n Version : %s\n Revision: %s\n", Version, Revision)
	},
}
