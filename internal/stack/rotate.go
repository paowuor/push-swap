package stack

import "push-swap/internal/models"

func Ra(a *models.Stack, ops *[]string) {
	if len(a.Values) < 2 {
		return
	}

	first := a.Values[0]
	a.Values = append(a.Values[1:], first)

	if ops != nil {
		*ops = append(*ops, "ra")
	}
}

func Rb(b *models.Stack, ops *[]string) {
	if len(b.Values) < 2 {
		return
	}

	first := b.Values[0]
	b.Values = append(b.Values[1:], first)

	if ops != nil {
		*ops = append(*ops, "rb")
	}
}

func Rr(a, b *models.Stack, ops *[]string) {
	Ra(a, nil)
	Rb(b, nil)

	if ops != nil {
		*ops = append(*ops, "rr")
	}
}