package pos

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSmallXCollate(t *testing.T) {
	params := NewParams(25, [32]byte{})
	require.Equal(t, "1000000000000000000000001000000000000000000000000110000000000000000000000100",
		fmt.Sprintf("%b", Collate(params, 1, 2, 3, 4)))
}
