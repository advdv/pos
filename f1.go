package pos

import (
	"math/big"
)

// readChaCha peforms the chacha byte reading
func readChaCha(params *Params, x uint64) uint64 {
	k, seed := params.K(), params.PlotSeed()
	q, r := divmod(x*k, 512)
	ciphertext0, end := ChaCha8(q, seed), r+k
	if end < 512 { // the bytes we need to read can be found in the first round of chacha8
		return Slice512(ciphertext0, uint(r), uint(end)).Uint64()
	}

	// else, append extra bytes of ciphertext, and slice that instead
	ciphertext1 := ChaCha8(q+1, seed)
	comb := new(big.Int).SetBytes(append(ciphertext0[:], ciphertext1[:]...))
	return Slice(comb, uint(r), uint(end), 1024).Uint64()
}

// F1 performs a ChaCha8 cipher of the provided x value
func F1(params *Params, x uint64) uint64 {
	ciphered := readChaCha(params, x)                                                    // ChaCha8(...)
	sliced := Slice(new(big.Int).SetUint64(x), 0, uint(params.Ext()), 25)                // x[:param_ext]
	return Concat(uint(params.Ext()), new(big.Int).SetUint64(ciphered), sliced).Uint64() // ||
}
