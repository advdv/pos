package pos

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBucketID(t *testing.T) {
	p := NewParams(32, [32]byte{})

	require.Equal(t, uint64(1), p.BucketID(15113))
	require.Equal(t, uint64(2), p.BucketID((2*15113)+1)) //floor
}

func TestIDbc(t *testing.T) {
	p := NewParams(32, [32]byte{})
	bid, cid := p.IDbc(1)
	require.Equal(t, uint64(0), bid)
	require.Equal(t, uint64(1), cid)

	bid, cid = p.IDbc(1 << 32)
	require.Equal(t, uint64(30), bid)
	require.Equal(t, uint64(16), cid)
}

func TestDivMod(t *testing.T) {
	for i, c := range []struct{ num, den, quo, rem uint64 }{
		{8, 3, 2, 2},
		{3, 8, 0, 3},
		{5, 5, 1, 0},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			quo, rem := divmod(c.num, c.den)
			require.Equal(t, c.quo, quo)
			require.Equal(t, c.rem, rem)
		})
	}
}
