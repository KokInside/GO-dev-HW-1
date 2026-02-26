package tokenizer

import (
	"calc/types"

	"errors"
	"strings"
	"unicode"
)

func ParseTokens(str string) ([]types.Token, error) {

	tokens := make([]types.Token, 0)

	var currentNumber strings.Builder

	var hasDot bool

	for i, symbol := range str {

		if unicode.IsDigit(symbol) {

			currentNumber.WriteRune(symbol)

			if i == len(str)-1 {

				// неявное умножение, если после ')' идёт число
				if len(tokens) != 0 && tokens[len(tokens)-1].Value == ")" {
					tokens = append(tokens, types.Token{Type: types.BinaryOp, Value: "*"})
				}

				tokens = append(tokens, types.Token{Type: types.Number, Value: currentNumber.String()})
			}

			continue
		}

		if symbol == '.' {
			currentNumber.WriteRune(symbol)

			if hasDot {

				return nil, errors.New("Second dot in number")
			}

			if i == len(str)-1 {

				tokens = append(tokens, types.Token{Type: types.Number, Value: currentNumber.String()})
			}

			hasDot = true

			continue
		}

		if currentNumber.Len() != 0 {

			// неявное умножение, если после ')' идёт число
			if len(tokens) != 0 && tokens[len(tokens)-1].Value == ")" {
				tokens = append(tokens, types.Token{Type: types.BinaryOp, Value: "*"})
			}

			tokens = append(tokens, types.Token{Type: types.Number, Value: currentNumber.String()})

			hasDot = false
			currentNumber.Reset()
		}

		if unicode.IsSpace(symbol) {
			continue
		}

		switch symbol {

		case '+', '-':

			curToken := types.Token{Value: string(symbol)}

			if len(tokens) == 0 ||
				tokens[len(tokens)-1].Type == types.UnaryOp ||
				tokens[len(tokens)-1].Type == types.BinaryOp ||
				tokens[len(tokens)-1].Value == "(" {

				curToken.Type = types.UnaryOp
			} else {

				curToken.Type = types.BinaryOp
			}

			tokens = append(tokens, curToken)

		case '*', '/':

			tokens = append(tokens, types.Token{Type: types.BinaryOp, Value: string(symbol)})

		case '(', ')':

			// неявное умножение, если после числа или ')' стоит '('
			if symbol == '(' && len(tokens) != 0 {
				if tokens[len(tokens)-1].Type == types.Number ||
					tokens[len(tokens)-1].Value == ")" {

					tokens = append(tokens, types.Token{Type: types.BinaryOp, Value: "*"})
				}
			}

			tokens = append(tokens, types.Token{Type: types.Parenthesis, Value: string(symbol)})

		default:
			return nil, errors.New("Unknown symbol")
		}
	}

	return tokens, nil
}

func Priority(token types.Token) int {

	if token.Type == types.UnaryOp {
		return 3
	}

	if token.Type == types.BinaryOp {

		switch token.Value {

		case "*", "/":
			return 2

		case "+", "-":
			return 1
		}
	}

	return 0
}
