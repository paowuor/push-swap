package stack

import "paowuor-push-swap/internal/models"

func Sa(a *models.Stack, ops *[]string) {
	if len(a.Values) < 2 {
		return
	}

	a.Values[0], a.Values[1] = a.Values[1], a.Values[0]

	if ops != nil {
		*ops = append(*ops, "sa")
	}
}

func Sb(b *models.Stack, ops *[]string) {
	if len(b.Values) < 2 {
		return
	}

	b.Values[0], b.Values[1] = b.Values[1], b.Values[0]

	if ops != nil {
		*ops = append(*ops, "sb")
	}
}

func Ss(a, b *models.Stack, ops *[]string) {
	Sa(a, nil)
	Sb(b, nil)

	if ops != nil {
		*ops = append(*ops, "ss")
	}
}