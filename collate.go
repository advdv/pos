package pos

import (
	"math/big"
	"strconv"
)

// Collate is a collation function. Returns 2^bk for b=colla_size(t). Xs are provided in lists of
// 1, 2, 4, 8, 16 or 32 length
func Collate(params *Params, xs ...uint64) *big.Int {
	switch len(xs) {
	case 1, 2, 4:
		return Concat64(uint(params.K()), xs...)
	case 8:
		panic("pos: not implemented: 8")
	case 16:
		panic("pos: not implemented: 16")
	case 32:
		panic("pos: not implemented: 32")
	default:
		panic("pos: collate called with unexpected nr of x-values: " + strconv.Itoa(len(xs)))
	}
}
