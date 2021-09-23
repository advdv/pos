package pos

import "strconv"

// Collate the xs values
func Collate(params *Params, xs ...Num) Num {
	switch len(xs) {
	case 1, 2, 4:
		return Concat(xs...)
	case 8:
		panic("pos: collate not implemented: 8")
	case 16:
		panic("pos: collate not implemented: 16")
	case 32:
		panic("pos: collate not implemented: 32")
	default:
		panic("pos: collate called with unexpected nr of x-values: " + strconv.Itoa(len(xs)))
	}
}
