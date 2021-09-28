package pos

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConcat(t *testing.T) {
	require.Equal(t, "025b0000000000000000000000001", fmt.Sprint(
		Concat(Num64(1, 25))))
	require.Equal(t, "029b00000000000000000000000010010", fmt.Sprint(
		Concat(Num64(1, 25), Num64(2, 4))))
	require.Equal(t, "034b0000000000000000000000001001000001", fmt.Sprint(
		Concat(Num64(1, 25), Num64(2, 4), Num64(1, 5))))
}
