package calc

import (
	"calc/calc/calculations"
	"calc/tokenizer"
	"calc/validators"
)

func Calculate(str string) (float64, error) {

	// спарсить токены из переданной строки
	tokens, parseError := tokenizer.ParseTokens(&str)

	if parseError != nil {
		return 0, parseError
	}

	validationError := validators.ValidateTokens(tokens)

	if validationError != nil {
		return 0, validationError
	}

	// инфиксная нотация -> Польская нотация
	var outQueue []string = calculations.InfixToPolish(tokens)

	// получить результат из польской нотации
	result, calculationError := calculations.SolveRPN(outQueue)

	if calculationError != nil {
		return 0, calculationError
	}

	return result, nil
}
