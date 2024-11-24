package gost341264_test

import (
	"bytes"
	"github.com/ptrvsrg/crypto/internal/block/gost341264"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_EncryptStream_DecryptStream(t *testing.T) {
	key := gost341264.GenerateKey()
	pt := []byte{0x41, 0x41, 0x41, 0x41, 0x41, 0x41}
	ct := []byte{0x9b, 0x5d, 0x44, 0xd7, 0xea, 0x20, 0x04, 0x22}

	encryptBuffer := bytes.NewBuffer(make([]byte, 8))
	err := gost341264.EncryptStream(bytes.NewReader(pt), encryptBuffer, key)
	require.NoError(t, err)

	encryptBytes := encryptBuffer.Bytes()
	require.ElementsMatch(t, ct, encryptBytes[len(encryptBytes)-len(ct):])

	decryptBuffer := bytes.NewBuffer(make([]byte, 8))
	err = gost341264.DecryptStream(bytes.NewReader(ct), decryptBuffer, key)
	require.NoError(t, err)

	decryptBytes := decryptBuffer.Bytes()
	require.ElementsMatch(t, pt, decryptBytes[len(decryptBytes)-len(pt):])
}
