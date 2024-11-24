package sha1

import (
	"encoding/hex"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Sum(t *testing.T) {
	hash := Sum([]byte("Hello, World!"))
	require.Equal(t, "0a0a9f2a6772942557ab5355d76af442f8f65e01", hex.EncodeToString(hash))
}
