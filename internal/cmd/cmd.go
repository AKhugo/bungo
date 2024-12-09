package cmd

import (
	"context"
	"fmt"
	"os/exec"
	"github.com/koffihuguesagossadou/bungo/pkg/common"
	"github.com/spf13/cobra"
)

var version = common.VERSION

func Do(args []string) int {
	rootCmd := &cobra.Command{ 
		Use: "bungo", 
		SilenceUsage: true,
		Short: "bungo is a tool to encode and decode files in base64",
		Long:  `bungo is a tool to encode and decode files in base64`,
	}
	
	rootCmd.AddCommand(versionCmd)

	// encode command with default flags
	encodeCmd.Flags().StringP("i", "i", "", "input file")
	encodeCmd.Flags().StringP("o", "o", "", "output file")
	rootCmd.AddCommand(encodeCmd)

	ctx := context.Background()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode()
		} else {
			return 1
		}
	}
	return 0
}

// version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of bungo",
	Long:  `All software has versions. This is bungo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bungo version", version)
	},
}

