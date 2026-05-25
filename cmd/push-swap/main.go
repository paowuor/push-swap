package main

import (
	"fmt"
	"os"
	"strings"

    "push-swap/internal/algorithms"
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

	if len(numbers) == 0 || utils.IsSorted(numbers) {
		return
	}

	normalized := algorithms.Normalize(numbers)

	a := models.Stack{
		Name:   "a",
		Values: normalized,
	}

	b := models.Stack{
		Name: "b",
	}

	var ops []string

	size := len(a.Values)

	if size == 2 {
		algorithms.SortTwo(&a, &ops)
	} else if size == 3 {
		algorithms.SortThree(&a, &ops)
	} else if size <= 5 {
		algorithms.SortFive(&a, &b, &ops)
	} else {
		algorithms.RadixSort(&a, &b, &ops)
	}

	fmt.Println(strings.Join(ops, "\n"))

	if len(ops) > 0 {
		fmt.Println()
	}		
}
