package cmd

import (
	"os"
	"path/filepath"

	"github.com/alomia/fastman/pkg/fileutils"
	"github.com/alomia/fastman/pkg/projectmanager"
	"github.com/alomia/fastman/pkg/sampledata"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [directory-name]",
	Short: "initialize a new FastAPI project",
	Long: `Create the basic directory, package, and file structure to start working with a FastAPI project.
The command initializes a new FastAPI project by setting up the necessary directories and files, allowing you to quickly get started with development.

Examples:
fastman init
fastman init project-name`,
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := os.Getwd()
		if err != nil {
			return err
		}

		if len(args) > 0 {
			directory := args[0]
			path = filepath.Join(path, directory)
			fileutils.CreateDirectory(path)
		}

		configFileName := "config.yaml"
		content, err := sampledata.GetSampleContent(configFileName)
		if err != nil {
			return err
		}

		configFilePath := filepath.Join(path, configFileName)
		err = fileutils.CreateFile(configFilePath, []byte(content))
		if err != nil {
			return err
		}

		viper.AddConfigPath(path)
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")

		err = viper.ReadInConfig()
		if err != nil {
			return err
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

		fastman := projectmanager.NewProjectStructure(
			path,
			packages,
			files,
		)

		fastman.CreateProjectStructure()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
