package algorithms

import (
	"push-swap/internal/models"
	"push-swap/internal/stack"
)

func getMaxBits(numbers []int) int {
	max := numbers[0]

	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	
	bits := 0

	for (max >> bits) != 0 {
		bits++
	}

	return bits
}

func RadixSort(a, b *models.Stack, ops *[]string) {
	normalized := Normalize(a.Values)
	
	maxBits := getMaxBits(normalized)

	size := len(normalized)

	for i := 0; i < maxBits; i++ {

		for j := 0; j < size; j++ {

			num := normalized[0]

			if ((num >> i) & 1) == 1 {
				stack.Ra(a, ops)

				normalized = append(normalized[1:], normalized[0])
			} else {
				stack.Pb(a, b, ops)

				normalized = normalized[1:]
			}
		}

		for len(b.Values) > 0 {
			stack.Pa(a, b, ops)
		}
		
		normalized = Normalize(a.Values)
	}
}
