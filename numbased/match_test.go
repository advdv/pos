package pos

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMatchReferenceProofDepth1(t *testing.T) {
	params := NewParams(25, [32]byte{0x01, 0x02, 0x2f, 0xb4, 0x2c, 0x08, 0xc1, 0x2d, 0xe3, 0xa6, 0xaf, 0x05, 0x38, 0x80, 0x19, 0x98, 0x06, 0x53, 0x2e, 0x79, 0x51, 0x5f, 0x94, 0xe8, 0x34, 0x61, 0x61, 0x21, 0x01, 0xf9, 0x41, 0x2f})
	proof, _ := ProofFromHex("228f532336a70179e3a96fef5d43cfc7753a527e876cfe328d7a169b4632bf95c62863df453c2d36e6f49a6967e7d58a57249a02c36638676117a73ca0db52c12a118e359346115a75ca5c454a67f8a3de32832801d33dab42246890142e247237f77dfae81c108cd1e01d9e195a9d4cee6491abf509acb301cc00b9bd2dab5a18aa6c07ee3583afd0b24937077557eb52797161b25ba308a440fbd4d35365d08d56d58d74028355ba33a44bef583f1af1801f995d32f4b228002d93c79a7555c87cdb00d7d11670", 25)

	for i := 0; i < 64; i += 2 {
		l, r := Fx(params, proof[i]), Fx(params, proof[i+1])
		require.Equal(t, 1, FindMatches(params, []uint64{l.Uint64()}, []uint64{r.Uint64()}))
	}
}

func TestFailProof(t *testing.T) {
	params := NewParams(25, [32]byte{0x01, 0x02, 0x2f, 0xb4, 0x2c, 0x08, 0xc1, 0x2d, 0xe3, 0xa6, 0xaf, 0x05, 0x38, 0x80, 0x19, 0x98, 0x06, 0x53, 0x2e, 0x79, 0x51, 0x5f, 0x94, 0xe8, 0x34, 0x61, 0x61, 0x21, 0x01, 0xf9, 0x41, 0x2f})
	proof, _ := ProofFromHex("228f532336a70179e3a96fef5d43cfc7753a527e876cfe328d7a169b4632bf95c62863df453c2d36e6f49a6967e7d58a57249a02c36638676117a73ca0db52c12a118e359346115a75ca5c454a67f8a3de32832801d33dab42246890142e247237f77dfae81c108cd1e01d9e195a9d4cee6491abf509acb301cc00b9bd2dab5a18aa6c07ee3583afd0b24937077557eb52797161b25ba308a440fbd4d35365d08d56d58d74028355ba33a44bef583f1af1801f995d32f4b228002d93c79a7555c87cdb00d7d11670", 25)
	proof[1] = Num64(1100, proof[1].Domain()) // corrupt one part of the proof

	l, r := Fx(params, proof[0]), Fx(params, proof[1])
	require.Equal(t, 0, FindMatches(params, []uint64{l.Uint64()}, []uint64{r.Uint64()}))
}

func TestPanicFindMatches(t *testing.T) {
	require.Panics(t, func() { FindMatches(nil, []uint64{1}, []uint64{}) })
	require.Panics(t, func() { FindMatches(nil, []uint64{}, []uint64{}) })
}
