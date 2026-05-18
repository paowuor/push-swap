package parser

func HasDuplicates(numbers []int) bool {
	seen := make(map[int]bool)

	for _, n := range numbers {
		if seen[n] {
			return true
		}
		seen[n] = true
	}

	return false
}
