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
	maxBits := getMaxBits(a.Values)

	size := len(a.Values)

	for i := 0; i < maxBits; i++ {

		for j := 0; j < size; j++ {

			num := a.Values[0]

			if ((num >> i) & 1) == 1 {
				stack.Ra(a, ops)
			} else {
				stack.Pb(a, b, ops)
			}
		}

		for len(b.Values) > 0 {
			stack.Pa(a, b, ops)
		}	
	}
}
