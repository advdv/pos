package pos

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParams(t *testing.T) {
	params := NewParams(25, [32]byte{0x01})
	require.Equal(t, [32]byte{0x01}, params.pseed)

	// the reference implementation has this value for the last entry, but the full table isn't compared
	require.Equal(t, uint64(8000), params.l_targets[1][15112][63])
}
