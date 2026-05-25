package algorithms

import "sort"

func Normalize(numbers []int) []int {
	sorted := make([]int, len(numbers))
	copy(sorted, numbers)

	sort.Ints(sorted)

	indexMap := make(map[int]int)

	for i, num := range sorted {
		indexMap[num] = i
	}

	normalized := make([]int, len(numbers))

	for i, num := range numbers {
		normalized[i] = indexMap[num]
	}

	return normalized
}
