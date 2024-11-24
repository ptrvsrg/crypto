package command

import (
	"os"

	"github.com/spf13/cobra"
)

var dstPath string

var rootCmd = &cobra.Command{
	Use:   "crypto-block",
	Short: "Block cipher to encrypt/decrypt a file",
	Long:  "Block cipher to encrypt/decrypt a file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("output", "o", "", "Output file path")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
