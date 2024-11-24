package vernam_test

import (
	"github.com/ptrvsrg/crypto/internal/stream/vernam"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Cipher(t *testing.T) {
	text := "Hello, World!"

	key, err := vernam.GenerateKey(len(text))
	require.NoError(t, err)

	encryptedText := vernam.Cipher([]byte(text), key)
	decryptedText := vernam.Cipher(encryptedText, key)

	require.Equal(t, text, string(decryptedText))
}
