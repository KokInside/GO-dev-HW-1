package calc

import (
	"calc/calc/calculations"
	"calc/tokenizer"
	"calc/validators"
)

func Calculate(expression string) (float64, error) {
	tokens, err := tokenizer.ParseTokens(expression)
	if err != nil {
		return 0, err
	}

	if err := validators.ValidateTokens(tokens); err != nil {
		return 0, err
	}

	// инфиксная нотация -> Польская нотация
	outQueue := calculations.InfixToPolish(tokens)

	// получить результат из польской нотации
	result, err := calculations.SolveRPN(outQueue)
	if err != nil {
		return 0, err
	}

	return result, nil
}
