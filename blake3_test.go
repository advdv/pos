package pos

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBlake3(t *testing.T) {
	s1 := Blake3([]byte{0x01})
	require.Equal(t, byte(0x48), s1[0])
	require.Equal(t, byte(0xfc), s1[1])
}
