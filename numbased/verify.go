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
