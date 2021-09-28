package pos

const (
	// Extra bits of output from the f functions. Instead of being a function from k -> k bits,
	// it's a function from k -> k + kExtraBits bits. This allows less collisions in matches.
	// Refer to the paper for mathematical motivations.
	param_ext = 6
	param_m   = 1 << param_ext

	// B and C groups which constitute a bucket, or BC group. These groups determine how
	// elements match with each other. Two elements must be in adjacent buckets to match.
	param_b  = 119
	param_c  = 127
	param_bc = param_b * param_c
)

// Params configure the proof of space
type Params struct {
	k     uint
	pseed [32]byte
	fsize uint64

	// l_targets tables is used in the efficient implementation of "findMatches"
	l_targets [2][param_bc][param_m]uint64
}

// NewParams inits a new set of parameters
func NewParams(k uint, pseed ...[32]byte) (p *Params) {
	p = &Params{
		k: k,
	}
	if len(pseed) > 0 {
		p.pseed = pseed[0]
	}

	p.fsize = uint64(p.k) + param_ext

	// fill the "ltargets" table
	for parity := uint64(0); parity < 2; parity++ {
		for i := uint64(0); i < param_bc; i++ {
			indJ := i / param_c
			for m := uint64(0); m < param_m; m++ {
				yr := ((indJ+m)%param_b)*param_c + (((2*m+parity)*(2*m+parity) + i) % param_c)
				p.l_targets[parity][i][m] = yr
			}
		}
	}

	return
}
