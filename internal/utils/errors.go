package utils

import (
	"fmt"
	"os"
)

func PrintError() {
	fmt.Fprintln(os.Stderr, "Error")
}
