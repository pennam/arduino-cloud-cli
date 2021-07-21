package cli

import (
	"fmt"
	"os"

	"github.com/bcmi-labs/iot-cloud-cli/cli/config"
	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(config.NewCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
