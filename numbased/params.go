package pos

// Params configure the proof of space
type Params struct {
	k     uint
	ext   uint64
	m     uint64
	b, c  uint64
	bc    uint64
	pseed [32]byte
}

// NewParams inits a new set of parameters
func NewParams(k uint, pseed ...[32]byte) (p *Params) {
	p = &Params{
		k:   k,
		ext: 6,
		b:   119,
		c:   127,
	}
	if len(pseed) > 0 {
		p.pseed = pseed[0]
	}

	p.bc = p.b * p.c
	p.m = 1 << p.ext
	return
}
