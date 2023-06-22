package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alomia/fastman/pkg/fileutils"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project, files, packages, or directories",
	Long: `Create a new project, files, packages, or directories using Fastman.
This command allows you to generate the necessary structure for your project, including files, Python packages, and directories.

Examples:
	fastman create -p mypackage
	fastman create -f myfile.py
	fastman create -d mydir
	fastman create project [name-directory] --fastapi`,
	RunE: func(cmd *cobra.Command, args []string) error {
		packageName, _ := cmd.Flags().GetString("package")
		fileName, _ := cmd.Flags().GetString("file")
		dirName, _ := cmd.Flags().GetString("directory")

		currentDir, err := os.Getwd()
		if err != nil {
			return err
		}

		if packageName != "" {
			packagePath := filepath.Join(currentDir, packageName)
			err := fileutils.CreatePackage(packagePath)
			if err != nil {
				return err
			}
			fmt.Printf("Package %s created in: %s\n", packageName, currentDir)
			return nil
		}

		if fileName != "" {
			filePath := filepath.Join(currentDir, fileName)
			err := fileutils.CreateFile(filePath)
			if err != nil {
				return err
			}
			fmt.Printf("File %s created in: %s\n", fileName, currentDir)
			return nil
		}

		if dirName != "" {
			dirPath := filepath.Join(currentDir, dirName)
			err := fileutils.CreateDirectory(dirPath)
			if err != nil {
				return err
			}
			fmt.Printf("Directory %s created in: %s\n", dirName, currentDir)
			return nil
		}

		return cmd.Help()
	},
}

func init() {
	createCmd.Flags().StringP("package", "p", "", "Create a new Python package")
	createCmd.Flags().StringP("file", "f", "", "Create a new file")
	createCmd.Flags().StringP("directory", "d", "", "directory name")
	rootCmd.AddCommand(createCmd)
}
