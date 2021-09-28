package pos

import (
	"math/big"
)

// Trunc returns the firxt (most significant) bits of x.
func Trunc(x Num, b uint) Num {
	return NewNum(new(big.Int).Rsh(x.Int, x.z-b), b)
}
