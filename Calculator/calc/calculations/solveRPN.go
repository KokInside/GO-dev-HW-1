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
			if polishStack.Size() < 2 {
				return 0, errors.New("Not enough operands")
			}

			rightOperand := polishStack.Top()
			polishStack.Pop()

			leftOperand := polishStack.Top()
			polishStack.Pop()

			switch token.Value {
			case "*":

				var result float64 = leftOperand * rightOperand
				polishStack.Push(result)

			case "/":

				if comparator.Float64Compare(rightOperand, 0.0) {
					return 0, errors.New("Division by zero")
				}

				var result float64 = leftOperand / rightOperand
				polishStack.Push(result)

			case "-":

				var result float64 = leftOperand - rightOperand
				polishStack.Push(result)

			case "+":

				var result float64 = leftOperand + rightOperand
				polishStack.Push(result)
			}

		case types.UnaryOp:
			if polishStack.Size() < 1 {
				return 0, errors.New("Not enough operands")
			}

			operand := polishStack.Top()
			polishStack.Pop()

			switch token.Value {
			case "+":
				var result float64 = operand
				polishStack.Push(result)

			case "-":
				var result float64 = -operand

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

	return polishStack.Top(), nil
}
