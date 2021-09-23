package pos

import (
	"math/big"

	"github.com/zeebo/blake3"
)

// Blake3 is a blake3 hash of 'x'
func Blake3(x []byte) (d [32]byte) { return blake3.Sum256(x) }

// A perform a high-level hash function on the combined input of 'l','r' and 'y'. The returned number
// has implied domain of 256 bits.
func A(l, r, y Num) Num {
	sum := Blake3(Concat(y, l, r).Bytes())
	return NewNum(new(big.Int).SetBytes(sum[:]), 256)
}
