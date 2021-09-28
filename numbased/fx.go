package pos

import (
	"math/big"
	"strconv"
)

// Fx composes A with C
func Fx(params *Params, xs ...Num) Num {
	switch len(xs) {
	case 1: // f1(x1)
		return Aprime(params, xs[0])
	case 2: // f2(x1,x2)
		return Trunc(
			A(
				C(params, xs[0]),
				C(params, xs[1]),
				Fx(params, xs[0])),
			uint(params.fsize))
	case 4: // f3(x1, x2, x3, x4)
		return Trunc(
			A(
				C(params, xs[0], xs[1]),
				C(params, xs[2], xs[3]),
				Fx(params, xs[0], xs[1])),
			uint(params.fsize))
	case 8: // f4(x1, ... x8)
		return Trunc(
			A(
				C(params, xs[0], xs[1], xs[2], xs[3]),
				C(params, xs[4], xs[5], xs[6], xs[7]),
				Fx(params, xs[0], xs[1], xs[2], xs[3])),
			uint(params.fsize))
	case 16:
		fallthrough
	case 32:
		fallthrough
	default:
		panic("pos: fx called with unexpected nr of x-values: " + strconv.Itoa(len(xs)))
	}
}

// Aprime performs a ChaCha8 cipher of the provided x value
func Aprime(params *Params, x Num) Num {
	return Concat(readChaCha(params, x), Slice(x, 0, param_ext)) // ChaCha8 || x[:param_ext]
}

// readChaCha peforms the chacha byte reading
func readChaCha(params *Params, x Num) Num {
	q, r := divmod(x.Uint64()*uint64(params.k), 512)
	ciphertext0, end := ChaCha8(q, params.pseed), r+uint64(params.k)
	if end < 512 { // the bytes we need to read can be found in the first round of chacha8
		return Slice(NewNum(new(big.Int).SetBytes(ciphertext0[:]), 512), uint(r), uint(end))
	}

	// else, append extra bytes of ciphertext, and slice that instead
	ciphertext1 := ChaCha8(q+1, params.pseed)
	comb := new(big.Int).SetBytes(append(ciphertext0[:], ciphertext1[:]...))
	return Slice(NewNum(comb, 1024), uint(r), uint(end))
}

// divmod as taken from: https://stackoverflow.com/questions/43945675/division-with-returning-quotient-and-remainder
// and tested again python implementation https://www.programiz.com/python-programming/methods/built-in/divmod
func divmod(x, m uint64) (quo, rem uint64) {
	quo = x / m
	rem = x % m
	return
}
