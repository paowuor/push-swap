package algorithms

import (
	"push-swap/internal/models"
	"push-swap/internal/stack"
)

func SortTwo(a *models.Stack, ops *[]string) {
	if len(a.Values) < 2 {
		return
	}

	if a.Values[0] > a.Values[1] {
		stack.Sa(a, ops)
	}
}
