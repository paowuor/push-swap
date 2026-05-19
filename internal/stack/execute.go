package stack

import (
	"errors"

	"paowuor-push-swap/internal/models"
)

func Execute(op string, a, b *models.Stack) error {
	switch op {

	case "sa":
		Sa(a, nil)

	case "sb":
		Sb(b, nil)

	case "ss":
		Ss(a, b, nil)

	case "pa":
		Pa(a, b, nil)

	case "pb":
		Pb(a, b, nil)

	case "ra":
		Ra(a, nil)

	case "rb":
		Rb(b, nil)

	case "rr":
		Rr(a, b, nil)

	case "rra":
		Rra(a, nil)

	case "rrb":
		Rrb(b, nil)

	case "rrr":
		Rrr(a, b, nil)

	default:
		return errors.New("invalid instruction")
	}

	return nil
}