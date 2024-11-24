package command

import (
	"encoding/hex"
	"fmt"
	"github.com/ptrvsrg/crypto/internal/hash/sha1"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var sha1Cmd = &cobra.Command{
	Use:   "sha1",
	Short: "SHA1 hash a string",
	Long:  "SHA1 hash a string",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		text := strings.Join(args, " ")

		hash := sha1.Sum([]byte(text))
		fmt.Printf("hash: %s", hex.EncodeToString(hash))
	},
}

func init() {
	rootCmd.AddCommand(sha1Cmd)
}
