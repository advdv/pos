package pos

import (
	"fmt"
	"math/big"
	"strconv"
)

// Num describes arbitrary large number with an implied domain
type Num struct {
	*big.Int
	z uint
}

// NewNum inits a num with implied domain of 'z' bits
func NewNum(n *big.Int, z uint) (num Num) {
	if z == 0 {
		panic("pos: number with implied domain '0' not allowed")
	}
	num = Num{n, z}
	num.checkOverflow()
	return
}

// Num64 creates a number with a implied domain of 'z' bits from 'n'
func Num64(n uint64, z uint) (num Num) {
	return NewNum(new(big.Int).SetUint64(n), z)
}

// checkOverflow will assert the num itself and panic if the implied domain is too small to represent
// the number this num holds.
func (num Num) checkOverflow() {
	if num.BitLen() <= int(num.z) {
		return
	}
	panic("pos: number too large for implied domain")
}

// Format is called by stringer
func (num Num) Format(f fmt.State, verb rune) {
	if num.BitLen() <= int(num.z) {
		fmt.Fprintf(f, "%03db%0"+strconv.Itoa(int(num.z))+"b", num.z, num.Int)
		return
	}
	fmt.Fprintf(f, "errb%0"+strconv.Itoa(int(num.z))+"b", num.Int)
}
