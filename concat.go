package pos

import "math/big"

// Concat implements zero-padding concatenation for domain z. NOTE: i'm assuming that the zero-padding only
// applies to all elements except the first, where extra zeros are always ignored
func Concat(z uint, xs ...*big.Int) (res *big.Int) {
	res = new(big.Int)
	for _, x := range xs {
		res.Lsh(res, uint(z)).Add(x, res)
	}
	return res
}
