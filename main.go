package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/davidlanouette/ValidParentheses/validator"
)

func main() {
	fmt.Println(len(os.Args))
	if (len(os.Args)) <= 1 {
		fmt.Println("You must specify an input string")
		return
	}
	input := os.Args[1]

	// remove any leading or trainling quotes
	input = strings.TrimLeft(input, `"`)
	input = strings.TrimRight(input, `"`)

	isValid := validator.ValidateParens(input)
	fmt.Printf("%v\n", isValid)
}
