package calculations

import (
	"calc/comparator"
	"calc/stack"
	"calc/types"

	"errors"
	"strconv"
)

// https://ru.wikipedia.org/wiki/%D0%9E%D0%B1%D1%80%D0%B0%D1%82%D0%BD%D0%B0%D1%8F_%D0%BF%D0%BE%D0%BB%D1%8C%D1%81%D0%BA%D0%B0%D1%8F_%D0%B7%D0%B0%D0%BF%D0%B8%D1%81%D1%8C

// Reverse Polish Notation
func SolveRPN(outQueue []types.Token) (float64, error) {

	var polishStack stack.Stack[float64]

	for _, token := range outQueue {

		switch token.Type {

		case types.BinaryOp:

			top, err := polishStack.Top()

			if err != nil {
				return 0, errors.New("Not enough operands")
			}

			rightOperand := top
			polishStack.Pop()

			top, err = polishStack.Top()

			if err != nil {
				return 0, errors.New("Not enough operands")
			}

			leftOperand := top
			polishStack.Pop()

			switch token.Value {
			case "*":

				result := leftOperand * rightOperand
				polishStack.Push(result)

			case "/":

				if comparator.Float64Compare(rightOperand, 0.0) {
					return 0, errors.New("Division by zero")
				}

				result := leftOperand / rightOperand
				polishStack.Push(result)

			case "-":

				result := leftOperand - rightOperand
				polishStack.Push(result)

			case "+":

				result := leftOperand + rightOperand
				polishStack.Push(result)
			}

		case types.UnaryOp:

			top, err := polishStack.Top()

			if err != nil {
				return 0, errors.New("Not enough operands")
			}

			operand := top
			polishStack.Pop()

			switch token.Value {
			case "+":
				result := operand
				polishStack.Push(result)

			case "-":
				result := -operand

				polishStack.Push(result)
			}

		default:

			res, parseErr := strconv.ParseFloat(token.Value, 64)

			if parseErr != nil {
				return 0, errors.New("Can't parse float")
			}

			polishStack.Push(res)
		}
	}

	top, err := polishStack.Top()

	if err != nil {
		return top, err
	}

	return top, nil
}
