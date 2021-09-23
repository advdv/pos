package pos

import "math/big"

// Concat numbers with a domain, the returned number is of the implied domain of all three numbers combined
func Concat(nums ...Num) Num {
	res, z := new(big.Int), uint(0)
	for _, n := range nums {
		z += n.z
		res.Lsh(res, uint(n.z)).Add(n.Int, res)
	}
	return NewNum(res, z)
}
