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
func NewParams(k uint, pseed [32]byte) (p *Params) {
	p = &Params{
		k:     k,
		ext:   6,
		b:     119,
		c:     127,
		pseed: pseed,
	}
	p.bc = p.b * p.c
	p.m = 1 << p.ext
	return
}

// K returns the configure K
func (params Params) K() uint64 { return uint64(params.k) }

// Ext returns the Ext param
func (params Params) Ext() uint64 { return uint64(params.ext) }

// PlotSeed returns the configured plot seed
func (Params Params) PlotSeed() [32]byte { return Params.pseed }

// C returns the C param
func (params Params) C() uint64 { return params.c }

// M returns the M param
func (params Params) M() uint64 { return params.m }

// B returns the B param
func (params Params) B() uint64 { return params.b }

// BucketID function as mentioned in the paper. Confirmed implementation by
// https://github.com/kargakis/chiapos/blob/7cccea7476a5fb342be47f420bce70a82b96b00a/pkg/parameters/parameters.go#L27
func (params Params) BucketID(x uint64) uint64 {
	return x / params.bc // the paper uses the semi-square brackets to indicate the 'floor'
}

// IDbc function as mentioned in the paper. More explicite but otherwise the same as reference:
// https://github.com/kargakis/chiapos/blob/7cccea7476a5fb342be47f420bce70a82b96b00a/pkg/parameters/parameters.go#L27
func (params Params) IDbc(x uint64) (bid, cid uint64) {
	return divmod(x%params.bc, params.c)
}

// divmod as taken from: https://stackoverflow.com/questions/43945675/division-with-returning-quotient-and-remainder
// and tested again python implementation https://www.programiz.com/python-programming/methods/built-in/divmod
func divmod(x, m uint64) (quo, rem uint64) {
	quo = x / m
	rem = x % m
	return
}
