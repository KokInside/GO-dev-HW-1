package calculations

import (
	"calc/stack"
	"calc/tokenizer"
	"calc/types"
)

// https://ru.wikipedia.org/wiki/%D0%90%D0%BB%D0%B3%D0%BE%D1%80%D0%B8%D1%82%D0%BC_%D1%81%D0%BE%D1%80%D1%82%D0%B8%D1%80%D0%BE%D0%B2%D0%BE%D1%87%D0%BD%D0%BE%D0%B9_%D1%81%D1%82%D0%B0%D0%BD%D1%86%D0%B8%D0%B8

// Infix notation -> Polish notation
func InfixToPolish(tokens []types.Token) []types.Token {

	var operatorStack stack.Stack[types.Token]

	var outQueue []types.Token

	for _, token := range tokens {

		if token.Type == types.Number {

			outQueue = append(outQueue, token)
			continue
		}

		if token.Type == types.BinaryOp {

			for top, _ := operatorStack.Top(); operatorStack.Size() != 0 && tokenizer.Priority(top) >= tokenizer.Priority(token); top, _ = operatorStack.Top() {

				outQueue = append(outQueue, top)
				operatorStack.Pop()
			}

			operatorStack.Push(token)

			continue
		}

		if token.Type == types.UnaryOp {

			operatorStack.Push(token)

			continue
		}

		if token.Type == types.Parenthesis {

			if token.Value == "(" {
				operatorStack.Push(token)
			} else {

				for top, _ := operatorStack.Top(); top.Value != "("; top, _ = operatorStack.Top() {

					outQueue = append(outQueue, top)
					operatorStack.Pop()
				}

				operatorStack.Pop()
			}

			continue
		}
	}

	for operatorStack.Size() != 0 {

		top, _ := operatorStack.Top()

		outQueue = append(outQueue, top)
		operatorStack.Pop()
	}

	return outQueue
}
