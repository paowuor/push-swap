package main

import (
	"os"

	"paowuor-push-swap/internal/parser"
	"paowuor-push-swap/internal/utils"
)

func main() {
	numbers, err := parser.ParseArguments(os.Args[1:])
	if err != nil {
		utils.PrintError()
		return
	}

	if len(numbers) == 0 {
		return
	}

	_ = numbers
}
