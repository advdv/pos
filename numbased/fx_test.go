package pos

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadChaCha(t *testing.T) {
	params := NewParams(25, [32]byte{0x10})
	require.Equal(t, uint64(32204853), readChaCha(params, Num64(30479984, 64)).Uint64())
	require.Equal(t, uint64(9513690), readChaCha(params, Num64(20, 20)).Uint64()) // requires extra chacha round
}

func TestF1(t *testing.T) {
	params := NewParams(25, [32]byte{0x10})
	require.Equal(t, "031b1111010110110100000110101111010", fmt.Sprint(Fx(params, Num64(30479984, 25))))
}

func TestReferenceF1DoubleChaCha(t *testing.T) {
	params := NewParams(25, [32]byte{0x01, 0x02, 0x2f, 0xb4, 0x2c, 0x08, 0xc1, 0x2d, 0xe3, 0xa6, 0xaf, 0x05, 0x38, 0x80, 0x19, 0x98, 0x06, 0x53, 0x2e, 0x79, 0x51, 0x5f, 0x94, 0xe8, 0x34, 0x61, 0x61, 0x21, 0x01, 0xf9, 0x41, 0x2f})
	x1 := uint64(0b1101010010011110010111000)
	f1x1 := Fx(params, Num64(x1, 25))
	require.Equal(t, uint64(0b1100000110111000111010000110101), f1x1.Uint64())
}

func TestReferenceF1SingleChaCha(t *testing.T) {
	params := NewParams(25, [32]byte{0x01, 0x02, 0x2f, 0xb4, 0x2c, 0x08, 0xc1, 0x2d, 0xe3, 0xa6, 0xaf, 0x05, 0x38, 0x80, 0x19, 0x98, 0x06, 0x53, 0x2e, 0x79, 0x51, 0x5f, 0x94, 0xe8, 0x34, 0x61, 0x61, 0x21, 0x01, 0xf9, 0x41, 0x2f})
	proof, err := ProofFromHex("228f532336a70179e3a96fef5d43cfc7753a527e876cfe328d7a169b4632bf95c62863df453c2d36e6f49a6967e7d58a57249a02c36638676117a73ca0db52c12a118e359346115a75ca5c454a67f8a3de32832801d33dab42246890142e247237f77dfae81c108cd1e01d9e195a9d4cee6491abf509acb301cc00b9bd2dab5a18aa6c07ee3583afd0b24937077557eb52797161b25ba308a440fbd4d35365d08d56d58d74028355ba33a44bef583f1af1801f995d32f4b228002d93c79a7555c87cdb00d7d11670", 25)
	require.NoError(t, err)
	f1x1, f1x2 := Fx(params, proof[0]), Fx(params, proof[1])
	require.Equal(t, uint64(0b0001111100011010011100001001000), f1x1.Uint64()) //f1(x1)
	require.Equal(t, uint64(0b0001111100011010111001101010001), f1x2.Uint64()) //f1(x2)

	f1x3, f1x4 := Fx(params, proof[2]), Fx(params, proof[3])
	require.Equal(t, uint64(0b0111001100110101100111011000001), f1x3.Uint64()) //f1(x3)
	require.Equal(t, uint64(0b0111001100110110001001011010010), f1x4.Uint64()) //f1(x4)
}

func TestF2(t *testing.T) {
	params := NewParams(25, [32]byte{0x01, 0x02, 0x2f, 0xb4, 0x2c, 0x08, 0xc1, 0x2d, 0xe3, 0xa6, 0xaf, 0x05, 0x38, 0x80, 0x19, 0x98, 0x06, 0x53, 0x2e, 0x79, 0x51, 0x5f, 0x94, 0xe8, 0x34, 0x61, 0x61, 0x21, 0x01, 0xf9, 0x41, 0x2f})
	proof, _ := ProofFromHex("228f532336a70179e3a96fef5d43cfc7753a527e876cfe328d7a169b4632bf95c62863df453c2d36e6f49a6967e7d58a57249a02c36638676117a73ca0db52c12a118e359346115a75ca5c454a67f8a3de32832801d33dab42246890142e247237f77dfae81c108cd1e01d9e195a9d4cee6491abf509acb301cc00b9bd2dab5a18aa6c07ee3583afd0b24937077557eb52797161b25ba308a440fbd4d35365d08d56d58d74028355ba33a44bef583f1af1801f995d32f4b228002d93c79a7555c87cdb00d7d11670", 25)

	f2y1 := Fx(params, proof[0], proof[1]) // f2(x1, x2)
	require.Equal(t, "031b0000100101011011000000001110011", fmt.Sprint(f2y1))
}

func TestF3(t *testing.T) {
	params := NewParams(25, [32]byte{0x01, 0x02, 0x2f, 0xb4, 0x2c, 0x08, 0xc1, 0x2d, 0xe3, 0xa6, 0xaf, 0x05, 0x38, 0x80, 0x19, 0x98, 0x06, 0x53, 0x2e, 0x79, 0x51, 0x5f, 0x94, 0xe8, 0x34, 0x61, 0x61, 0x21, 0x01, 0xf9, 0x41, 0x2f})
	proof, _ := ProofFromHex("228f532336a70179e3a96fef5d43cfc7753a527e876cfe328d7a169b4632bf95c62863df453c2d36e6f49a6967e7d58a57249a02c36638676117a73ca0db52c12a118e359346115a75ca5c454a67f8a3de32832801d33dab42246890142e247237f77dfae81c108cd1e01d9e195a9d4cee6491abf509acb301cc00b9bd2dab5a18aa6c07ee3583afd0b24937077557eb52797161b25ba308a440fbd4d35365d08d56d58d74028355ba33a44bef583f1af1801f995d32f4b228002d93c79a7555c87cdb00d7d11670", 25)

	f3y1 := Fx(params, proof[0], proof[1], proof[2], proof[3]) // f2(x1, x2, x3, x4)
	require.Equal(t, "031b0001101101001111110011001011001", fmt.Sprint(f3y1))
}

func TestF4(t *testing.T) {
	params := NewParams(25, [32]byte{0x01, 0x02, 0x2f, 0xb4, 0x2c, 0x08, 0xc1, 0x2d, 0xe3, 0xa6, 0xaf, 0x05, 0x38, 0x80, 0x19, 0x98, 0x06, 0x53, 0x2e, 0x79, 0x51, 0x5f, 0x94, 0xe8, 0x34, 0x61, 0x61, 0x21, 0x01, 0xf9, 0x41, 0x2f})
	proof, _ := ProofFromHex("228f532336a70179e3a96fef5d43cfc7753a527e876cfe328d7a169b4632bf95c62863df453c2d36e6f49a6967e7d58a57249a02c36638676117a73ca0db52c12a118e359346115a75ca5c454a67f8a3de32832801d33dab42246890142e247237f77dfae81c108cd1e01d9e195a9d4cee6491abf509acb301cc00b9bd2dab5a18aa6c07ee3583afd0b24937077557eb52797161b25ba308a440fbd4d35365d08d56d58d74028355ba33a44bef583f1af1801f995d32f4b228002d93c79a7555c87cdb00d7d11670", 25)

	f3y1 := Fx(params, proof[0], proof[1], proof[2], proof[3], proof[4], proof[5], proof[6], proof[7]) // f2(x1, ... x8)
	require.Equal(t, "031b0110011101100011010001110100100", fmt.Sprint(f3y1))
}

func TestDivMod(t *testing.T) {
	for i, c := range []struct{ num, den, quo, rem uint64 }{
		{8, 3, 2, 2},
		{3, 8, 0, 3},
		{5, 5, 1, 0},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			quo, rem := divmod(c.num, c.den)
			require.Equal(t, c.quo, quo)
			require.Equal(t, c.rem, rem)
		})
	}
}
