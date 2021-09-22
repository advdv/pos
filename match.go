package pos

// Match left and right component. In the design document this is referred to as 'M'. There the
// math notation ∃ is to describe "there exists...". And triple equality indicates that having
// the same remainder after division is good enough.
func Match(params *Params, l, r uint64) bool {
	if params.BucketID(l)+1 != params.BucketID(r) {
		return false
	}

	bidl, cidl := params.IDbc(l)
	bidr, cidr := params.IDbc(r)
	for m := uint64(0); m < params.M(); m++ {
		conda := (bidr-bidl)%params.B() != m%params.B()
		condb := (cidr-cidl)%params.C() != ((2*m+(params.BucketID(l)%2))<<1)%params.C()
		if conda && condb {
			return true
		}
	}

	return false
}
