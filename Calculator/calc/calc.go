package calc

import (
	"calc/calc/calculations"
	"calc/tokenizer"
	"calc/types"
	"calc/validators"
)

func Calculate(expression string) (float64, error) {

	tokens, parseError := tokenizer.ParseTokens(expression)

	if parseError != nil {
		return 0, parseError
	}

	if err := validators.ValidateTokens(tokens); err != nil {
		return 0, err
	}

	// инфиксная нотация -> Польская нотация
	var outQueue []types.Token = calculations.InfixToPolish(tokens)

	// получить результат из польской нотации
	result, calculationError := calculations.SolveRPN(outQueue)

	if calculationError != nil {
		return 0, calculationError
	}

	return result, nil
}
