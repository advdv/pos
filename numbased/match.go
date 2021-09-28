package pos

// FindMatches will compare the L and R buckets and efficiently match as described with the matching funcion
// 'M' in the published paper.
func FindMatches(params *Params, bucket_L, bucket_R []uint64) (idx_count int) {
	if len(bucket_L) != len(bucket_R) || len(bucket_L) < 1 {
		panic("pos: must have equal y values match buckets, and need at least one entry")
	}

	parity := (bucket_L[0] / param_bc) % 2
	rmap := map[uint64]struct{ count, pos int }{}
	remove := (bucket_R[0] / param_bc) * param_bc

	for pos_R := 0; pos_R < len(bucket_R); pos_R++ {
		r_y := bucket_R[pos_R] - remove

		entry := rmap[r_y]
		if entry.count < 1 {
			entry.pos = pos_R
		}
		entry.count++
		rmap[r_y] = entry
	}

	remove_y := remove - param_bc
	for pos_L := 0; pos_L < len(bucket_L); pos_L++ {
		r := bucket_L[pos_L] - remove_y

		for i := 0; i < int(param_m); i++ {
			if r > uint64(len(params.l_targets[parity])-1) || i > len(params.l_targets[parity][r])-1 {
				continue // out-of-bound of l_targets: r_target will not exist
			}

			r_target := params.l_targets[parity][r][i]
			for j := 0; j < rmap[r_target].count; j++ {
				// @TODO handle registering the found matches if it's more then just the check
				idx_count++
			}
		}
	}

	return
}
