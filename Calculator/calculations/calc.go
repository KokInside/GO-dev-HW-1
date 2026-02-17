package calculations

import (
	"calc/stack"
	"calc/types"
	"calc/utils"
	"calc/validators"

	"errors"
	"strconv"
	"strings"
	"unicode"
)

func Calculate(str string) (float64, error) {

	// спарсить токены из переданной строки
	tokens, parseError := parseTokens(&str)

	if parseError != nil {
		return 0, parseError
	}

	validationError := validators.ValidateTokens(&tokens)

	if validationError != nil {
		return 0, validationError
	}

	// инфиксная нотация -> Польская нотация
	var outQueue []string = infixToPolish(&tokens)

	// получить результат из польской нотации
	result, calculationError := solveRPN(&outQueue)

	if calculationError != nil {
		return 0, calculationError
	}

	return result, nil

}

func priority(token string) int {
	switch token {
	case "*", "/":
		return 2

	case "+", "-":
		return 1
	}

	return 0
}

func solveRPN(outQueue *[]string) (float64, error) {

	var polishStack stack.Stack[float64]

	for _, token := range *outQueue {

		if token == "*" || token == "/" || token == "+" || token == "-" {

			rightOperand := polishStack.Top()
			polishStack.Pop()

			leftOperand := polishStack.Top()
			polishStack.Pop()

			switch token {
			case "*":
				var result float64 = leftOperand * rightOperand
				polishStack.Push(result)

			case "/":
				if utils.Float64Compare(rightOperand, 0.0) {
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
		} else {
			res, _ := strconv.ParseFloat(token, 64)
			polishStack.Push(res)
		}
	}

	return polishStack.Top(), nil
}

func infixToPolish(tokens *[]types.Token) []string {

	var operatorStack stack.Stack[string]

	var outQueue []string

	for _, token := range *tokens {

		if token.Type == types.Number {
			outQueue = append(outQueue, token.Value)
			continue
		}

		if token.Type == types.Operator {
			for operatorStack.Size() != 0 && priority(operatorStack.Top()) >= priority(token.Value) {
				outQueue = append(outQueue, operatorStack.Top())
				operatorStack.Pop()
			}

			operatorStack.Push(token.Value)

			continue
		}

		if token.Type == types.Parenthesis {
			if token.Value == "(" {
				operatorStack.Push(token.Value)
			} else {
				for operatorStack.Top() != "(" {
					outQueue = append(outQueue, operatorStack.Top())
					operatorStack.Pop()
				}
				operatorStack.Pop()
			}

			continue
		}
	}

	for operatorStack.Size() != 0 {
		outQueue = append(outQueue, operatorStack.Top())
		operatorStack.Pop()
	}

	return outQueue
}

func parseTokens(str *string) ([]types.Token, error) {

	var tokens []types.Token

	var currentNumber strings.Builder
	var hasDot bool

	for i, symbol := range *str {

		if unicode.IsDigit(symbol) {

			currentNumber.WriteRune(symbol)

			if i == len(*str)-1 {

				tokens = append(tokens, types.Token{Type: types.Number, Value: currentNumber.String()})
			}

			continue
		}

		if symbol == '.' {
			currentNumber.WriteRune(symbol)

			if hasDot == true {

				return []types.Token{}, errors.New("Second dot in number")
			}

			hasDot = true

			continue
		}

		if currentNumber.Len() != 0 {
			tokens = append(tokens, types.Token{Type: types.Number, Value: currentNumber.String()})

			hasDot = false
			currentNumber.Reset()
		}

		if unicode.IsSpace(symbol) {
			continue
		}

		switch symbol {

		case '*', '/', '+', '-':

			tokens = append(tokens, types.Token{Type: types.Operator, Value: string(symbol)})

		case '(', ')':

			tokens = append(tokens, types.Token{Type: types.Parenthesis, Value: string(symbol)})

		default:
			return []types.Token{}, errors.New("Unknown symbol")
		}

	}

	return tokens, nil
}
