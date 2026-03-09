package main

import (
	"calc/calc"

	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		fmt.Println("Enter the expression in a such way: go run calc.go \"[your expression]\"")
		return
	}

	if len(args) > 2 {
		fmt.Println("Too much arguments")
		return
	}

	// parsed string
	expression := args[1]

	result, err := calc.Calculate(expression)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println(result)
}
