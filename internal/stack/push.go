package stack

import "push-swap/internal/models"

func Pa(a, b *models.Stack, ops *[]string) {
	if len(b.Values) == 0 {
		return
	}

	value := b.Values[0]
	b.Values = b.Values[1:]

	a.Values = append([]int{value}, a.Values...)

	if ops != nil {
		*ops = append(*ops, "pa")
	}
}

func Pb(a, b *models.Stack, ops *[]string) {
	if len(a.Values) == 0 {
		return
	}

	value := a.Values[0]
	a.Values = a.Values[1:]

	b.Values = append([]int{value}, b.Values...)

	if ops != nil {
		*ops = append(*ops, "pb")
	}
}