package pos

import (
	"math/big"

	"github.com/zeebo/blake3"
)

// A perform a high-level hash function on the combined input of 'l','r' and 'y'. The returned number
// has implied domain of 256 bits.
func A(l, r, y Num) Num {
	in, b := Concat(y, l, r), make([]byte, 64)
	numb := cdiv(int(in.Domain()), 8)
	in.ToBlakeBytes(b)
	sum := blake3.Sum256(b[:numb])
	return NewNum(new(big.Int).SetBytes(sum[:]), 256)
}

// cdiv is used to determine the nr of bytes required to bit the concat bits
func cdiv(a, b int) int { return (a + b - 1) / b }
