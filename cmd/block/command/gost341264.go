package command

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/ptrvsrg/crypto/internal/block/gost341264"
	"github.com/spf13/cobra"
	"os"
)

var gost341264Cmd = &cobra.Command{
	Use:   "gost341264",
	Short: "GOST 34.12-2015 64-bit to encrypt/decrypt a file",
	Long:  "GOST 34.12-2015 64-bit to encrypt/decrypt a file",
	PreRun: func(cmd *cobra.Command, args []string) {
		dstPath, _ = cmd.Flags().GetString("output")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var encryptCmd = &cobra.Command{
	Use:     "encrypt <path>",
	Aliases: []string{"e"},
	Short:   "Encrypt a file using block cipher",
	Long:    "Encrypt a file using block cipher",
	PreRun: func(cmd *cobra.Command, args []string) {
		dstPath, _ = cmd.Flags().GetString("output")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return runEncrypt(args)
	},
}

var decryptCmd = &cobra.Command{
	Use:     "decrypt <path>",
	Aliases: []string{"d"},
	Short:   "Decrypt a file using stream cipher",
	Long:    "Decrypt a string using stream cipher",
	PreRun: func(cmd *cobra.Command, args []string) {
		dstPath, _ = cmd.Flags().GetString("output")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return runDecrypt(args)
	},
}

func init() {
	rootCmd.AddCommand(gost341264Cmd)
	gost341264Cmd.AddCommand(decryptCmd)
	gost341264Cmd.AddCommand(encryptCmd)
}

func runEncrypt(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("the argument is required")
	}
	if len(args) > 1 {
		return fmt.Errorf("too many arguments")
	}

	inPath := args[0]

	in, err := os.Open(inPath)
	if err != nil {
		return errors.Wrap(err, "failed to open input file")
	}
	defer in.Close()

	out := os.Stdout
	if dstPath != "" {
		out, err = os.Create(dstPath)
		if err != nil {
			return errors.Wrap(err, "failed to create output file")
		}
	}
	defer func() {
		if dstPath != "" {
			out.Close()
		}
	}()

	key := gost341264.GenerateKey()

	return gost341264.EncryptStream(in, out, key)
}

func runDecrypt(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("the argument is required")
	}
	if len(args) > 1 {
		return fmt.Errorf("too many arguments")
	}

	inPath := args[0]

	in, err := os.Open(inPath)
	if err != nil {
		return errors.Wrap(err, "failed to open input file")
	}
	defer in.Close()

	out := os.Stdout
	if dstPath != "" {
		out, err = os.Create(dstPath)
		if err != nil {
			return errors.Wrap(err, "failed to create output file")
		}
	}
	defer func() {
		if dstPath != "" {
			out.Close()
		}
	}()

	key := gost341264.GenerateKey()

	return gost341264.DecryptStream(in, out, key)
}
