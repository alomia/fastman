package cmd

import (
	"strings"

	"github.com/alomia/fastman/pkg/executils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var installCmd = &cobra.Command{
	Use:   "install [package-name]",
	Short: "Install packages or dependencies",
	Long: `Install packages or dependencies specified either by package name or from a file.
If a package name is provided, Fastman will attempt to install the specified package.
If a file is provided, Fastman will install the dependencies listed in the file.
	
Examples:
	fastman install
	fastman install my-package
	fastman install -f dependencies.txt`,
	RunE: func(cmd *cobra.Command, args []string) error {
		filename, _ := cmd.Flags().GetString("file")
		dependenciesFile := viper.GetString("dependencies_file")
		scriptInstallPackage := viper.GetString("scripts.install_package")
		scriptInstallFile := viper.GetString("scripts.install_from_file")

		var program string
		var arguments []string

		switch {
		case filename != "":
			scriptInstallFile = strings.Replace(scriptInstallFile, "{{package_file}}", filename, 1)
			command := strings.Fields(scriptInstallFile)
			program = command[0]
			arguments = command[1:]

		case len(args) == 0:
			scriptInstallFile = strings.Replace(scriptInstallFile, "{{package_file}}", dependenciesFile, 1)
			command := strings.Fields(scriptInstallFile)
			program = command[0]
			arguments = command[1:]

		case len(args) > 0:
			scriptInstallPackage = strings.Replace(scriptInstallPackage, "{{package_name}}", args[0], 1)
			command := strings.Fields(scriptInstallPackage)
			arguments = append(command[1:], args[1:]...)
			program = command[0]
		}

		if program != "" {
			err := executils.ExecuteCommand(program, arguments)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	installCmd.Flags().StringP("file", "f", "", "Installs the dependencies found in the specified file")
	rootCmd.AddCommand(installCmd)
}
