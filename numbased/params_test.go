package pos

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParams(t *testing.T) {
	params := NewParams(25)
	require.Equal(t, uint64(64), params.m)

	params = NewParams(25, [32]byte{0x01})
	require.Equal(t, [32]byte{0x01}, params.pseed)
}
