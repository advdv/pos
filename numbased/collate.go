package pos

import "strconv"

// C collates the xs values
func C(params *Params, xs ...Num) Num {
	switch len(xs) {
	case 1, 2, 4:
		return Concat(xs...)
	case 8:
		return Slice(
			A(
				C(params, xs[:4]...),
				C(params, xs[4:]...),
				Fx(params, xs[:4]...),
			),
			uint(params.fsize), uint(params.fsize)+(4*params.k))
	case 16:
		return Slice(
			A(
				C(params, xs[:8]...),
				C(params, xs[8:]...),
				Fx(params, xs[:8]...),
			),
			uint(params.fsize), uint(params.fsize)+(3*params.k))
	case 32:
		return Slice(
			A(
				C(params, xs[:16]...),
				C(params, xs[16:]...),
				Fx(params, xs[:16]...),
			),
			uint(params.fsize), uint(params.fsize)+(2*params.k))
	default:
		panic("pos: collate called with unexpected nr of x-values: " + strconv.Itoa(len(xs)))
	}
}
