package checker

import "bufio"

func ReadInstructions(scanner *bufio.Scanner) []string {
	var instructions []string

	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	return instructions
}
