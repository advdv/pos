package pos

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProofFromBytes(t *testing.T) {
	proof, err := ProofFromHex("228f532336a70179e3a96fef5d43cfc7753a527e876cfe328d7a169b4632bf95c62863df453c2d36e6f49a6967e7d58a57249a02c36638676117a73ca0db52c12a118e359346115a75ca5c454a67f8a3de32832801d33dab42246890142e247237f77dfae81c108cd1e01d9e195a9d4cee6491abf509acb301cc00b9bd2dab5a18aa6c07ee3583afd0b24937077557eb52797161b25ba308a440fbd4d35365d08d56d58d74028355ba33a44bef583f1af1801f995d32f4b228002d93c79a7555c87cdb00d7d11670", 25)
	require.NoError(t, err)
	require.Equal(t, uint64(30479984), proof[63])

	// x1, x2 double checked with reference implementation
	require.Equal(t, uint64(0b0010001010001111010100110), proof[0])
	require.Equal(t, uint64(0b0100011001101101010011100), proof[1])
}
