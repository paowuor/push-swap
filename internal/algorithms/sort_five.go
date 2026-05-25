package algorithms

import (
	"push-swap/internal/models"
	"push-swap/internal/stack"
)

func moveMinToTop(a *models.Stack, ops *[]string) {
	minIndex := FindMinIndex(a.Values)

	if minIndex == 1 {
		stack.Ra(a, ops)
	} else if minIndex == 2 {
		stack.Ra(a, ops)
		stack.Ra(a, ops)
	} else if minIndex == 3 {
		stack.Rra(a, ops)
		stack.Rra(a, ops)
	} else if minIndex == 4 {
		stack.Rra(a, ops)
	}
}

func SortFive(a, b *models.Stack, ops *[]string) {
	for len(a.Values) > 3 {
		moveMinToTop(a, ops)
		stack.Pb(a, b, ops)
	}
	
	SortThree(a, ops)

	for len(b.Values) > 0 {
		stack.Pa(a, b, ops)
	}
}
