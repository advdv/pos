package pos

// ChaCha8 is an 8-round ChaCha encryption as specified in the design document. The paper doesn't explicitely
// mention the iv but the reference implementation has it filled with zeros:
// https://github.com/Chia-Network/chiapos/blob/ce6ccc28ea300bfbad7e054a3f30433eb6090cf5/src/calculate_bucket.hpp#L88
func ChaCha8(c uint64, K [32]byte) [64]byte {
	return ChaCha(8, c, K, 0)
}

// ChaCha performs a N-round ChaCha encryption of a data block with zeros at block offset 'c' with key 'K'.
// The low and high bits of block counter 'c' are placed into state words 12 and 13, respectively. This
// implementation is based on: https://github.com/romain-jacotin/ChaCha/blob/master/ChaCha.go
func ChaCha(n int, c uint64, K [32]byte, iv uint64) (d [64]byte) {
	var grid [16]uint32
	var i, j int

	// the first four input words are constants
	grid[0], grid[1], grid[2], grid[3] = 0x61707865, 0x3320646e, 0x79622d32, 0x6b206574

	// input words 4 through 11 are taken from the 256-bit key by reading the bytes in little-endian order
	for j = 0; j < 8; j++ {
		grid[j+4] = 0
		for i = 0; i < 4; i++ {
			grid[j+4] += uint32(K[j*4+i]) << (8 * i)
		}
	}

	// The low and high bits of c are placed into state words 12 and 13, respectively
	grid[12] = uint32(c)
	grid[13] = uint32((c >> 32))

	// Lastly, words 14 and 15 are taken from an 8-byte nonce, again by reading the bytes in little-endian
	grid[14] = uint32(iv)
	grid[15] = uint32((iv >> 32))
	return ccrounds(n, grid)
}

// ccrounds run 'n' rounds of ChaCha on the grid
func ccrounds(n int, grid [16]uint32) (d [64]byte) {
	var j uint32
	var x [16]uint32
	for i := 0; i < 16; i++ {
		x[i] = grid[i]
	}

	// ChaCha8 consists of 8 rounds, alternating between "column" rounds and "diagonal" rounds.
	// Each round applies the "quarterround" function four times, to a different set of words each time.
	for i := 0; i < n/2; i++ {

		// QUARTER-ROUND on column 1:
		x[0], x[4], x[8], x[12] = ccquarterround(x[0], x[4], x[8], x[12])

		// QUARTER-ROUND on column 2:
		x[1], x[5], x[9], x[13] = ccquarterround(x[1], x[5], x[9], x[13])

		// QUARTER-ROUND on column 3:
		x[2], x[6], x[10], x[14] = ccquarterround(x[2], x[6], x[10], x[14])

		// QUARTER-ROUND on column 4:
		x[3], x[7], x[11], x[15] = ccquarterround(x[3], x[7], x[11], x[15])

		// QUARTER-ROUND on diagonal 1:
		x[0], x[5], x[10], x[15] = ccquarterround(x[0], x[5], x[10], x[15])

		// QUARTER-ROUND on diagonal 2:
		x[1], x[6], x[11], x[12] = ccquarterround(x[1], x[6], x[11], x[12])

		// QUARTER-ROUND on diagonal 3:
		x[2], x[7], x[8], x[13] = ccquarterround(x[2], x[7], x[8], x[13])

		// QUARTER-ROUND on diagonal 4:
		x[3], x[4], x[9], x[14] = ccquarterround(x[3], x[4], x[9], x[14])
	}

	// After 20 rounds of the above processing, the original 16 input words are added to the 16 words to form the 16 output words.
	for i := 0; i < 16; i++ {
		x[i] += grid[i]
	}

	// The 64 output bytes are generated from the 16 output words by serialising them in little-endian order and concatenating the results.
	for i := 0; i < 64; i += 4 {
		j = x[i>>2]
		d[i] = byte(j)
		d[i+1] = byte(j >> 8)
		d[i+2] = byte(j >> 16)
		d[i+3] = byte(j >> 24)
	}

	return
}

// quarter-round function performs 4 additions, 4 XORs and 4 bitwise left rotations between 4 choosen uint32 value
func ccquarterround(a, b, c, d uint32) (ra, rb, rc, rd uint32) {
	a += b
	d ^= a
	d = d<<16 | d>>16

	c += d
	b ^= c
	b = b<<12 | b>>20

	a += b
	d ^= a
	d = d<<8 | d>>24

	c += d
	b ^= c
	b = b<<7 | b>>25
	return a, b, c, d
}
