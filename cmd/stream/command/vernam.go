package command

import (
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ptrvsrg/crypto/internal/stream/vernam"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var vernamCmd = &cobra.Command{
	Use:   "vernam",
	Short: "Vernam cipher to encrypt/decrypt a string",
	Long:  "Vernam cipher to encrypt/decrypt a string",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var decryptCmd = &cobra.Command{
	Use:     "decrypt <HEX key> <HEX encrypted text>",
	Aliases: []string{"d"},
	Short:   "Decrypt a string using stream cipher",
	Long:    "Decrypt a string using stream cipher",
	RunE: func(cmd *cobra.Command, args []string) error {
		decryptedText, err := runDecrypt(args)
		if err != nil {
			return err
		}

		fmt.Printf("decrypted text: %s\n", decryptedText)
		return nil
	},
}

var encryptCmd = &cobra.Command{
	Use:     "encrypt <string>",
	Aliases: []string{"e"},
	Short:   "Encrypt a string using stream cipher",
	Long:    "Encrypt a string using stream cipher",
	RunE: func(cmd *cobra.Command, args []string) error {
		key, encryptedText, err := runEncrypt(args)
		if err != nil {
			return err
		}

		fmt.Printf("key (hex): %s\n", key)
		fmt.Printf("encrypted text (hex): %s\n", encryptedText)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(vernamCmd)
	vernamCmd.AddCommand(encryptCmd)
	vernamCmd.AddCommand(decryptCmd)
}

func runEncrypt(args []string) (string, string, error) {
	if len(args) < 1 {
		return "", "", fmt.Errorf("the argument is required")
	}

	text := strings.Join(args, " ")

	key, err := vernam.GenerateKey(len(text))
	if err != nil {
		return "", "", errors.Wrap(err, "failed to generate key")
	}

	encryptedText := vernam.Cipher([]byte(text), key)

	return hex.EncodeToString(key), hex.EncodeToString(encryptedText), nil
}

func runDecrypt(args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("the arguments is required")
	}
	if len(args) > 2 {
		return "", fmt.Errorf("too many arguments")
	}

	keyHex := args[0]
	textHex := strings.Join(args[1:], " ")

	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return "", errors.Wrap(err, "failed to HEX decode key")
	}

	text, err := hex.DecodeString(textHex)
	if err != nil {
		return "", errors.Wrap(err, "failed to HEX decode text")
	}

	decryptedText := vernam.Cipher(text, key)

	return string(decryptedText), nil
}
