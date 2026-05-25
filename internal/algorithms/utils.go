package algorithms

func FindMinIndex(numbers []int) int {
	minIndex := 0

	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[minIndex] {
			minIndex = i
		}
	}

	return minIndex
}