package tokenizer

import (
	"calc/types"

	"errors"
	"strings"
	"unicode"
)

func ParseTokens(str *string) ([]types.Token, error) {

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

				return nil, errors.New("Second dot in number")
			}

			if i == len(*str)-1 {

				tokens = append(tokens, types.Token{Type: types.Number, Value: currentNumber.String()})
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
			return nil, errors.New("Unknown symbol")
		}
	}

	return tokens, nil
}

func Priority(token string) int {
	switch token {
	case "*", "/":
		return 2

	case "+", "-":
		return 1
	}

	return 0
}
