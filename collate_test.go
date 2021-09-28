package pos

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCollate(t *testing.T) {
	params := NewParams(4)
	require.Panics(t, func() { C(params) }) // unsupported number of xs: 0

	xs := []Num{Num64(1, params.k), Num64(2, params.k), Num64(3, params.k), Num64(4, params.k)}
	collated := C(params, xs...)

	require.Equal(t, "016b0001001000110100", fmt.Sprint(collated))
	require.Equal(t, uint(int(params.k)*len(xs)), collated.Domain())
}
