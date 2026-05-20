package checker

import "bufio"

func ReadInstructions(scanner *bufio.Scanner) []string {
	var instructions []string

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		instructions = append(instructions, line)
	}

	return instructions
}
