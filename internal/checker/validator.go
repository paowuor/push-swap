package checker

func IsValidInstruction(op string) bool {
	valid := map[string]bool{
		"sa":  true,
		"sb":  true,
		"ss":  true,
		"pa":  true,
		"pb":  true,
		"ra":  true,
		"rb":  true,
		"rr":  true,
		"rra": true,
		"rrb": true,
		"rrr": true,
	}

	return valid[op]
}
