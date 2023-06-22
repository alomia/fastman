package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alomia/fastman/pkg/fileutils"
	"github.com/alomia/fastman/pkg/projectmanager"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "project [directory-name]",
	Short: "Create a new project",
	Long: `Create a new project using Fastman. Currently, only FastAPI projects are supported.
This command allows you to create a new project with a specific structure.

Examples:
	fastman create project myproject --fastapi`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fastapi, _ := cmd.Flags().GetBool("fastapi")

		currentDir, err := os.Getwd()
		if err != nil {
			return err
		}

		if fastapi {
			if len(args) > 0 {
				directory := args[0]
				targetDir := filepath.Join(currentDir, directory)

				err := fileutils.CreateDirectory(targetDir)
				if err != nil {
					return err
				}

				currentDir = targetDir
			}

			packages := []projectmanager.Package{
				{
					Name:  "models",
					Files: []string{"products.py"},
				},
				{
					Name:  "routes",
					Files: []string{"products.py"},
				},
			}

			files := []string{
				"main.py",
				"requirements.txt",
			}

			projectStruct := projectmanager.NewProjectStructure(
				currentDir,
				packages,
				files,
			)

			err := projectStruct.CreateProjectStructure()
			if err != nil {
				return err
			}

			fmt.Printf("FastAPI project structure created in: %s\n", currentDir)
			return nil
		}

		return cmd.Help()
	},
}

func init() {
	projectCmd.Flags().Bool("fastapi", false, "create a simple FastAPI project structure")
	createCmd.AddCommand(projectCmd)
}
