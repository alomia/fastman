package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alomia/fastman/pkg/fileutils"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [directory-name]",
	Short: "initialize fastman",
	Long: `Initialize Fastman by creating a configuration file called "fastman.yaml" in the specified directory.
The configuration file contains the necessary information for Fastman to function properly.
	
Examples:
	fastman init
	fastman init [directory-name]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		currentDir, err := os.Getwd()
		if err != nil {
			return err
		}

		if len(args) > 0 {
			directory := args[0]

			targetDir := filepath.Join(currentDir, directory)

			exists, err := fileutils.FileOrDirExists(targetDir)
			if err != nil {
				return err
			}

			if !exists {
				fmt.Printf("Directory '%s' does not exist.\n", directory)
				return nil
			}

			currentDir = targetDir
		}

		err = fileutils.CreateConfigFile(currentDir)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
