package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(VersionCmd)
}

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the current version of challengefile",
	RunE:  versionHandler,
}

const (
	version = "v0.0.1"
)

func versionHandler(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("version command does not accept any arguments")
	}

	fmt.Println(version)
	return nil
}
