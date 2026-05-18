package parser

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func ParseArguments(args []string) ([]int, error) {
	if len(args) == 0 {
		return []int{}, nil
	}

	var numbers []int

	joined := strings.Join(args, " ")
	fields := strings.Fields(joined)

	for _, field := range fields {
		num64, err := strconv.ParseInt(field, 10, 64)
		if err != nil {
			return nil, errors.New("invalid integer")
		}

		if num64 < math.MinInt32 || num64 > math.MaxInt32 {
			return nil, errors.New("integer overflow")
		}

		numbers = append(numbers, int(num64))
	}

	if HasDuplicates(numbers) {
		return nil, errors.New("duplicate numbers")
	}

	return numbers, nil
}
