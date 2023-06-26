package cmd

import (
	"fmt"
	"strings"

	"github.com/alomia/fastman/pkg/executils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [script]",
	Short: "Run a script",
	Long: `Run a script specified by the provided script name.
The script must be defined in the configuration file.

Examples:
  fastman run server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		runScripts := viper.GetStringMapString("scripts.run")

		if len(args) == 0 {
			return cmd.Help()
		}

		script, ok := runScripts[args[0]]
		if !ok {
			return fmt.Errorf("%s script does not exist", args[0])
		}

		command := strings.Fields(script)
		program := command[0]
		arguments := command[1:]

		err := executils.ExecuteCommand(program, arguments)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
