package pos

import (
	"math/big"
)

// Slice implements x[a:b]. If x has implied domain [2^z], then this is the a'th to (b-1)th bits of x
// taken as a number, where the first (most significant) bit is considered the 0th. NOTE: so any leading
// non-significant zeros are skipped while slicing.
func Slice(x *big.Int, a, b, z uint) *big.Int {
	if b > z {
		panic("slice: end index 'b' larger than domain")
	}
	if a > b {
		panic("slice: start index 'a' larger than end index 'b'")
	}
	if b < 1 {
		b = z
	}

	l := new(big.Int).Rsh(x, z-b)             // x >> (z-b)
	r := new(big.Int).Lsh(big.NewInt(1), b-a) // 1 << (b-a)
	return new(big.Int).Mod(l, r)             // l % r
}

// Slice512 is a readable way to slice the 512 chacha output.
func Slice512(x [64]byte, a, b uint) *big.Int {
	return Slice(new(big.Int).SetBytes(x[:]), a, b, 512)
}

// Slice256 slices a 256 bit output, usefull for slice the Blake3 output
func Slice256(x [32]byte, a, b uint) *big.Int {
	return Slice(new(big.Int).SetBytes(x[:]), a, b, 256)
}

// Slice64 slices a uint64
func Slice64(x uint64, a, b uint) uint64 {
	return Slice(new(big.Int).SetUint64(x), a, b, 64).Uint64()
}
