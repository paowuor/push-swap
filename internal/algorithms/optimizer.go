package algorithms

func OptimizeOperations(ops []string) []string {
	var optimized []string

	i := 0

	for i < len(ops) {

		if i+1 < len(ops) {
			// ra + rb => rr
			if ops[i] == "ra" && ops[i+1] == "rb" {
				optimized = append(optimized, "rr")
				i += 2
				continue
			}

			// rb + ra => rr
			if ops[i] == "rb" && ops[i+1] == "ra" {
				optimized = append(optimized, "rr")
				i += 2
				continue
			}

			// rra + rrb => rrr
			if ops[i] == "rra" && ops[i+1] == "rrb" {
				optimized = append(optimized, "rrr")
				i += 2
				continue
			}
			
			// rrb + rra => rrr
			if ops[i] == "rrb" && ops[i+1] == "rra" {
				optimized = append(optimized, "rrr")
				i += 2
				continue
			}
		}

		optimized = append(optimized, ops[i])
		i++
	}

	return optimized
}