package pos

import (
	"github.com/zeebo/blake3"
)

// Blake3 is a blake3 hash of 'x'
func Blake3(x []byte) (d [32]byte) { return blake3.Sum256(x) }
