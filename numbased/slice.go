package pos

import "math/big"

// slice the big int part
func slice(x *big.Int, a, b, z uint) *big.Int {
	l := new(big.Int).Rsh(x, z-b)             // x >> (z-b)
	r := new(big.Int).Lsh(big.NewInt(1), b-a) // 1 << (b-a)
	return new(big.Int).Mod(l, r)             // l % r
}

// Slice implements x[a:b]. If x has implied domain [2^z], then this is the a'th to (b-1)th bits of x
// taken as a number, where the first (most significant) bit is considered the 0th. The implied domain
// of what is returned is always a-b
func Slice(x Num, a, b uint) Num {
	if b > x.z {
		panic("pos: slice end index 'b' larger than implied domain of 'x'")
	}
	if a > b {
		panic("pos: slice start index 'a' larger than end index 'b'")
	}
	if b < 1 {
		b = x.z
	}
	return NewNum(slice(x.Int, a, b, x.z), b-a)
}
