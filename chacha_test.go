package pos

import (
	"encoding/binary"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChaCha(t *testing.T) {
	t.Run("vector 1", func(t *testing.T) {
		var k [32]byte
		d1 := ChaCha(20, 0, k, 0)
		require.Equal(t,
			"76b8e0ada0f13d90405d6ae55386bd28bdd219b8a08ded1aa836efcc8b770dc7da41597c5157488d7724e03fb8d84a376a43b8f41518a11cc387b669b2ee6586",
			hex.EncodeToString(d1[:]),
		)
	})

	t.Run("vector 2", func(t *testing.T) {
		var k [32]byte
		k[31] = 0x01
		d1 := ChaCha(20, 0, k, 0)
		require.Equal(t,
			"4540f05a9f1fb296d7736e7b208e3c96eb4fe1834688d2604f450952ed432d41bbe2a0b6ea7566d2a5d1e7e20d42af2c53d792b1c43fea817e9ad275ae546963",
			hex.EncodeToString(d1[:]),
		)
	})

	t.Run("vector 4", func(t *testing.T) {
		var k [32]byte
		iv := binary.LittleEndian.Uint64([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01})
		d1 := ChaCha(20, 0, k, iv)
		require.Equal(t,
			"de9cba7bf3d69ef5e786dc63973f653a0b49e015adbff7134fcb7df137821031e85a050278a7084527214f73efc7fa5b5277062eb7a0433e445f41e31afab757",
			hex.EncodeToString(d1[:]),
		)
	})

	t.Run("vector 4", func(t *testing.T) {
		var k [32]byte
		iv := binary.LittleEndian.Uint64([]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
		d1 := ChaCha(20, 0, k, iv)
		require.Equal(t,
			"ef3fdfd6c61578fbf5cf35bd3dd33b8009631634d21e42ac33960bd138e50d32111e4caf237ee53ca8ad6426194a88545ddc497a0b466e7d6bbdb0041b2f586b",
			hex.EncodeToString(d1[:]),
		)
	})
}
