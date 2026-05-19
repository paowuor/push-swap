package main

import (
	"fmt"
	"os"

	"paowuor-push-swap/internal/models"
	"paowuor-push-swap/internal/parser"
	"paowuor-push-swap/internal/stack"
	"paowuor-push-swap/internal/utils"
)

func main() {
	numbers, err := parser.ParseArguments(os.Args[1:])
	if err != nil {
		utils.PrintError()
		return
	}

	if len(numbers) == 0 || utils.IsSorted(numbers) {
		return
	}

	a := models.Stack{
		Name:   "a",
		Values: numbers,
	}

	b := models.Stack{
		Name: "b",
	}

	var ops []string

	stack.Sa(&a, &ops)
	stack.Pb(&a, &b, &ops)

	for _, op := range ops {
		fmt.Println(op)
	}
}
