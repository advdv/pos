package pos

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumberPrinting(t *testing.T) {
	require.Equal(t, "025b0000000000000000000000001", fmt.Sprint(Num64(1, 25)))
	require.Equal(t, "001b1", fmt.Sprint(Num64(1, 1)))

	require.Panics(t, func() { Num64(100, 0) }) // zero domai not supported
	require.Panics(t, func() { Num64(2, 1) })   // implied domain too small

	n1 := Num64(1, 1)
	n1.z = 0
	require.Equal(t, "errb1", fmt.Sprint(n1))
}

func TestUint64(t *testing.T) {
	require.Equal(t, uint64(12), Num64(12, 64).Uint64())
	require.Panics(t, func() {
		n1 := Num64(1, 1)
		n1.z = 65
		n1.Uint64()
	}) // domain too small
}

func TestToBytes(t *testing.T) {
	nb, _ := hex.DecodeString("3e34e1208a3d4c8cda9c")
	n := new(big.Int)
	n.SetBytes(nb)
	num := NewNum(n, 81)

	// input as measured in the reference implementation
	require.Equal(t, "081b000111110001101001110000100100000100010100011110101001100100011001101101010011100", fmt.Sprint(num))

	// refb is the byte representation of the input into blake3 as observed by the reference implementation
	refb, _ := hex.DecodeString("1f1a7090451ea6466d4e006d01000000c067c90201800b6d48388d0f00000000b8e4166d01000000f9698d0f00000000a0e4166d0100000020e5166d01000000")

	var inputBytes [64]byte
	num.ToBlakeBytes(inputBytes[:])
	require.Equal(t, refb[:11], inputBytes[:11])
}
