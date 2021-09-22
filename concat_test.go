package pos

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConcat(t *testing.T) {
	require.Equal(t, "1000001",
		fmt.Sprintf("%b", Concat(6, big.NewInt(0b1), big.NewInt(0b1)).Uint64()))
	require.Equal(t, "1000001000010",
		fmt.Sprintf("%b", Concat(6, big.NewInt(0b1), big.NewInt(0b1), big.NewInt(0b10)).Uint64()))
}
