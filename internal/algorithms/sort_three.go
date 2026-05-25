package algorithms

import (
	"push-swap/internal/models"
	"push-swap/internal/stack"
)

func SortThree(a *models.Stack, ops *[]string) {
	first := a.Values[0]
	second := a.Values[1]
	third := a.Values[2]

	if first > second && second > third {
		stack.Sa(a, ops)
		stack.Rra(a, ops)
	} else if first > second && second < third && first > third {
		stack.Ra(a, ops)
	} else if first < second && second > third && first > third {
		stack.Rra(a, ops)
	} else if first < second && second > third && first < third {
		stack.Sa(a, ops)
		stack.Ra(a, ops)
	} else if first > second && second < third && first < third {
		stack.Sa(a, ops)
	}
}
