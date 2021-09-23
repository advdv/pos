package pos

import (
	"fmt"
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
