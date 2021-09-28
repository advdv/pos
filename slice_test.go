package pos

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSlice(t *testing.T) {
	require.Equal(t, big.NewInt(2), Slice(Num64(0b100100, 6), 2, 5).Int)
	require.Equal(t, "003b010", fmt.Sprint(Slice(Num64(0b100100, 6), 2, 5)))
	require.Equal(t, big.NewInt(36), Slice(Num64(0b100100, 6), 0, 0).Int)

	require.Panics(t, func() { Slice(Num64(1, 1), 1, 0) }) // 1>0
	require.Panics(t, func() { Slice(Num64(1, 1), 1, 2) }) // 2 larger then implied domain of 1
}
