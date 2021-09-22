package pos

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSlice(t *testing.T) {
	require.Equal(t, big.NewInt(2), Slice(big.NewInt(0b100100), 2, 5, 6))
	require.Equal(t, big.NewInt(36), Slice(big.NewInt(0b100100), 0, 0, 6))
	require.Panics(t, func() {
		Slice(big.NewInt(0b100100), 1, 0, 6)
	})
	require.Panics(t, func() {
		Slice(big.NewInt(0b100100), 1, 7, 6)
	})

	t.Run("slice with big number", func(t *testing.T) {
		v := bytes32()
		require.Equal(t,
			hex.EncodeToString(v[:16]),
			hex.EncodeToString(Slice256(v, 2, 128).Bytes()),
		)
	})
}

func TestSlice512(t *testing.T) {
	v := bytes64()
	require.Equal(t,
		hex.EncodeToString(v[:32]),
		hex.EncodeToString(Slice512(v, 0, 256).Bytes()))
}

func TestSlice256(t *testing.T) {
	v := bytes32()
	require.Equal(t,
		hex.EncodeToString(v[:16]),
		hex.EncodeToString(Slice256(v, 2, 128).Bytes()),
	)
}

func TestSlice64(t *testing.T) {
	require.Equal(t, uint64(0x0001020304050607), Slice64(0x0001020304050607, 0, 64))
	require.Equal(t, uint64(0x00010203), Slice64(0x0001020304050607, 0, 32))
}

func bytes32() (v [32]byte) {
	for i := 1; i <= len(v); i++ {
		v[i-1] = byte(i)
	}
	return
}

func bytes64() (v [64]byte) {
	for i := 1; i <= len(v); i++ {
		v[i-1] = byte(i)
	}
	return
}
