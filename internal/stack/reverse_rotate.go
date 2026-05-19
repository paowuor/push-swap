package stack

import "paowuor-push-swap/internal/models"

func Rra(a *models.Stack, ops *[]string) {
	if len(a.Values) < 2 {
		return
	}

	last := a.Values[len(a.Values)-1]

	a.Values = append([]int{last}, a.Values[:len(a.Values)-1]...)

	if ops != nil {
		*ops = append(*ops, "rra")
	}
}

func Rrb(b *models.Stack, ops *[]string) {
	if len(b.Values) < 2 {
		return
	}

	last := b.Values[len(b.Values)-1]

	b.Values = append([]int{last}, b.Values[:len(b.Values)-1]...)

	if ops != nil {
		*ops = append(*ops, "rrb")
	}
}

func Rrr(a, b *models.Stack, ops *[]string) {
	Rra(a, nil)
	Rrb(a, nil)

	if ops != nil {
		*ops = append(*ops, "rrr")
	}
}