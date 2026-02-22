package validators

import (
	"calc/comparator"
	"calc/types"

	"errors"
	"strconv"
)

func ValidateTokens(tokens []types.Token) error {

	if err := emptyExpression(tokens); err != nil {
		return err
	}

	if err := unbalancedParentheses(tokens); err != nil {
		return err
	}

	if err := notEnoughOperands(tokens); err != nil {
		return err
	}

	if err := notEnoughOperators(tokens); err != nil {
		return err
	}

	if err := divisionByZero(tokens); err != nil {
		return err
	}

	return nil
}

// Validators:

func emptyExpression(tokens []types.Token) error {
	if len(tokens) == 0 {
		return errors.New("Empty expression")
	}

	return nil
}

func unbalancedParentheses(tokens []types.Token) error {

	var open, close int

	for _, token := range tokens {

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

func notEnoughOperands(tokens []types.Token) error {

	var isPrevBinaryOp bool

	for _, token := range tokens {

		if token.Type == types.BinaryOp {

			if isPrevBinaryOp {
				return errors.New("Not enough operands")
			}

			isPrevBinaryOp = true

		} else {
			isPrevBinaryOp = false
		}
	}

	// Первый токен не может быть бинарным оператором
	// Последний токен не может быть любым оператором
	if tokens[0].Type == types.BinaryOp ||
		tokens[len(tokens)-1].Type == types.BinaryOp ||
		tokens[len(tokens)-1].Type == types.UnaryOp {
		return errors.New("Not enough operands")
	}

	return nil
}

func notEnoughOperators(tokens []types.Token) error {

	var isPrevOperand bool

	for _, token := range tokens {

		if token.Type == types.Number {

			if isPrevOperand {
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
func divisionByZero(tokens []types.Token) error {

	var isPrevDivisor bool

	for _, token := range tokens {

		if isPrevDivisor && token.Type == types.Number {

			v, err := strconv.ParseFloat(token.Value, 64)

			if err != nil {
				return errors.New("Can't parse float")
			}

			if comparator.Float64Compare(v, 0) {

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
