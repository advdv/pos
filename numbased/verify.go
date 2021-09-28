package pos

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/32bitkid/bitreader"
)

// Proof holds the 64 x-values
type Proof [64]Num

// ProofFromBytes read the 64 x-values
func ProofFromBytes(r io.Reader, k uint) (p Proof, err error) {
	br := bitreader.NewReader(r)
	for i := 0; i < len(p); i++ {
		pn, err := br.Read64(k)
		if err != nil {
			return p, fmt.Errorf("failed to read x-values from proof bytes: %w", err)
		}
		p[i] = Num64(pn, k)
	}

	return
}

// ProofFromHex returns the proof x-values from hex encoded string
func ProofFromHex(s string, k uint) (p Proof, err error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return p, fmt.Errorf("failed to hex-decode proof: %w", err)
	}

	return ProofFromBytes(bytes.NewReader(b), k)
}

// Verify proof 'p'
func Verify(params *Params, p Proof) (ok bool) {
	ys := make([]Num, 0, 64)
	for i := 0; i < 64; i++ {
		ys = append(ys, Fx(params, p[i])) // f1 values
	}

	for depth := 2; depth < 8; depth++ {
		numFxArgs := 1 << (depth - 1)
		numMatch := (1 << (8 - depth))
		newys := make([]Num, 0, numMatch)

		for i := 0; i < numMatch; i += 2 {
			start, end := (i/2)*numFxArgs, ((i/2)*numFxArgs)+numFxArgs
			lpe, rpe := ys[i], ys[i+1]
			if FindMatches(params, []uint64{lpe.Uint64()}, []uint64{rpe.Uint64()}) != 1 {
				return false // verification fails
			}

			newys = append(newys, Fx(params, p[start:end]...))
		}

		ys = newys
	}

	return true // all x-values match, verificate succeeded
}
