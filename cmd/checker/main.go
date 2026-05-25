package main

import (
	"bufio"
	"fmt"
	"os"

	ck "push-swap/internal/checker"
	"push-swap/internal/models"
	"push-swap/internal/parser"
	"push-swap/internal/stack"
	"push-swap/internal/utils"
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

	a := models.Stack{
		Name:   "a",
		Values: numbers,
	}

	b := models.Stack{
		Name: "b",
	}

	scanner := bufio.NewScanner(os.Stdin)

	instructions := ck.ReadInstructions(scanner)

	for _, instruction := range instructions {
		if !ck.IsValidInstruction(instruction) {
			utils.PrintError()
			return
		}

		err := stack.Execute(instruction, &a, &b)
		if err != nil {
			utils.PrintError()
			return
		}
	}

	if utils.IsStackSorted(a.Values) && len(b.Values) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
