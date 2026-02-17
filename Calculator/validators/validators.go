package validators

import (
	"calc/types"
	"calc/utils"
	"errors"
	"strconv"
)

func ValidateTokens(tokens *[]types.Token) error {

	empryError := emptyExpression(tokens)

	if empryError != nil {
		return empryError
	}

	parenthesesError := unbalancedParentheses(tokens)

	if parenthesesError != nil {
		return parenthesesError
	}

	operandsError := notEnoughOperands(tokens)

	if operandsError != nil {
		return operandsError
	}

	operatorsError := notEnoughOperators(tokens)

	if operatorsError != nil {
		return operatorsError
	}

	divisionByZeroError := divisionByZero(tokens)

	if divisionByZeroError != nil {
		return divisionByZeroError
	}

	return nil
}

// Validators:

// несбалансированные скобки
func unbalancedParentheses(tokens *[]types.Token) error {

	var open, close int

	for _, token := range *tokens {

		switch token.Value {
		case "(":
			open++

		case ")":
			close++

		}

		if close > open {
			return errors.New("Missing (")
		}

	}

	if open == close {
		return nil
	} else {
		return errors.New("Missing )")
	}
}

// несколько операторов подряд
func notEnoughOperands(tokens *[]types.Token) error {

	var isPrevOperator bool

	for _, token := range *tokens {

		if token.Type == types.Operator {

			if isPrevOperator == true {
				return errors.New("Not enough operands")
			}

			isPrevOperator = true

		} else {
			isPrevOperator = false
		}
	}

	// первый и последний токены не должны быть операторами
	if (*tokens)[0].Type == types.Operator || (*tokens)[len(*tokens)-1].Type == types.Operator {
		return errors.New("Not enough operands")
	}

	return nil
}

func notEnoughOperators(tokens *[]types.Token) error {

	var isPrevOperand bool

	for _, token := range *tokens {

		if token.Type == types.Number {

			if isPrevOperand == true {
				return errors.New("Not enough operators")
			}

			isPrevOperand = true

		} else {
			isPrevOperand = false
		}
	}

	return nil
}

// деление на ноль до этапа раскрытия скобок
func divisionByZero(tokens *[]types.Token) error {

	var isPrevDivisor bool

	for _, token := range *tokens {

		if isPrevDivisor == true && token.Type == types.Number {

			v, err := strconv.ParseFloat(token.Value, 64)

			if err != nil {
				return errors.New("Can't parse float")
			}

			if utils.Float64Compare(v, 0) {

				return errors.New("Division by zero")
			}
		}

		if token.Value == "/" {
			isPrevDivisor = true
		} else {
			isPrevDivisor = false
		}
	}

	return nil
}

func emptyExpression(tokens *[]types.Token) error {
	if len(*tokens) == 0 {
		return errors.New("Empty expression")
	}
	return nil
}
