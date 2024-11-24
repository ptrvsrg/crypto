package command

import (
	"encoding/hex"
	"fmt"
	"github.com/ptrvsrg/crypto/internal/hash/sha1"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crypto-hash",
	Short: "Hash a string",
	Long:  "Hash a string",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		text := strings.Join(args, " ")

		hash := sha1.Sum([]byte(text))
		fmt.Printf("Hash: %s\n", hex.EncodeToString(hash))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
