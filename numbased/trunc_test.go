package pos

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTrunc(t *testing.T) {
	require.Equal(t, "001b1", fmt.Sprint(
		Trunc(Num64(0b10, 2), 1)))
	require.Equal(t, "001b0", fmt.Sprint(
		Trunc(Num64(0b01, 2), 1)))
	require.Equal(t, "002b01", fmt.Sprint(
		Trunc(Num64(0b01, 2), 2)))
	require.Equal(t, "002b01", fmt.Sprint(
		Trunc(Num64(0b01101, 5), 2)))
	require.Equal(t, "003b011", fmt.Sprint(
		Trunc(Num64(0b01101, 5), 3)))
	require.Equal(t, "003b110", fmt.Sprint(
		Trunc(Num64(0b01101, 4), 3)))
}
