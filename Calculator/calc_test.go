package main

import (
	"calc/calculations"
	"calc/utils"
	"testing"
)

func TestCalcAddition(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{"addition", "2+55", 57},
		{"addition", "2+4+6", 12},
		{"addition", "0+0", 0},
		{"addition", "2+0+9+0", 11},
	}

	for i, test := range tests {
		result, error := calculations.Calculate(test.input)

		if error != nil {
			t.Fatal("Fatal error in calc function, test:", test.name, i, error)
		}

		if !utils.Float64Compare(result, test.expected) {
			t.Errorf("Error. Test: %v, %v. Expected %v, got %v", test.name, i, test.expected, result)
		}
	}

}

func TestCalcSubtraction(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{"subtraction", "10-1", 9},
		{"subtraction", "10-3-4", 3},
		{"subtraction", "0-0", 0},
		{"subtraction", "10-101", -91},
	}

	for i, test := range tests {
		result, error := calculations.Calculate(test.input)

		if error != nil {
			t.Fatal("Fatal error in calc function, test:", test.name, i, error)
		}

		if !utils.Float64Compare(result, test.expected) {
			t.Errorf("Error. Test: %v, %v. Expected %v, got %v", test.name, i, test.expected, result)
		}
	}
}

func TestCalcMultiplication(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{"multiplication", "3*4", 12},
		{"multiplication", "3*0", 0},
		{"multiplication", "0*4", 0},
		{"multiplication", "0*0", 0},
		{"multiplication", "199*1", 199},
		{"multiplication", "2*3*4", 24},
		{"multiplication", "1*2*3*4*5*6*0", 0},
		{"multiplication", "1*2*3*4*5", 120},
	}

	for i, test := range tests {
		result, error := calculations.Calculate(test.input)

		if error != nil {
			t.Fatal("Fatal error in calc function, test:", test.name, i, error)
		}

		if !utils.Float64Compare(result, test.expected) {
			t.Errorf("Error. Test: %v, %v. Expected %v, got %v", test.name, i, test.expected, result)
		}
	}
}

func TestCalcDivision(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{"division", "6/2", 3},
		{"division", "120/5", 24},
		{"division", "103/1", 103},
		{"division", "20/2/5", 2},
		{"division", "1/1", 1},
		{"division", "0/28", 0},
		{"division", "10/3", float64(10.0 / 3.0)},
		{"division", "1763/98", float64(1763.0 / 98.0)},
	}

	for i, test := range tests {
		result, error := calculations.Calculate(test.input)

		if error != nil {
			t.Fatal("Fatal error in calc function, test:", test.name, i, error)
		}

		if !utils.Float64Compare(result, test.expected) {
			t.Errorf("Error. Test: %v, %v. Expected %v, got %v", test.name, i, test.expected, result)
		}
	}
}

func TestCalcParentheses(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{"parentheses", "(2+2)*2", 8},
		{"parentheses", "(10)", 10},
		{"parentheses", "2*(3+1)*4", 32},
		{"parentheses", "((2+3)*4)/2", 10},
		{"parentheses", "(1+2)-3", 0},
		{"parentheses", "(7-100)", -93},
	}

	for i, test := range tests {
		result, error := calculations.Calculate(test.input)

		if error != nil {
			t.Fatal("Fatal error in calc function, test:", test.name, i, error)
		}

		if !utils.Float64Compare(result, test.expected) {
			t.Errorf("Error. Test: %v, %v. Expected %v, got %v", test.name, i, test.expected, result)
		}
	}
}

func TestCalcPriority(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{"priority", "2+2*2", 6},
		{"priority", "2*2+2", 6},
		{"priority", "24/6*10", 40},
		{"priority", "24*6/10", float64(24.0 * 6.0 / 10.0)},
	}

	for i, test := range tests {
		result, error := calculations.Calculate(test.input)

		if error != nil {
			t.Fatal("Fatal error in calc function, test:", test.name, i, error)
		}

		if !utils.Float64Compare(result, test.expected) {
			t.Errorf("Error. Test: %v, %v. Expected %v, got %v", test.name, i, test.expected, result)
		}
	}
}

func TestCalcError(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedError string
	}{
		{"unbalancedParentheses", "(7*4+9)+87)", "Missing ("},
		{"unbalancedParentheses", ")(5+7)/443+73", "Missing ("},
		{"unbalancedParentheses", "(54*3-8", "Missing )"},
		{"unbalancedParentheses", "7+322)", "Missing ("},
		{"unbalancedParentheses", ")443+22(", "Missing ("},

		{"notEnoughOperands", "+", "Not enough operands"},
		{"notEnoughOperands", "//", "Not enough operands"},
		{"notEnoughOperands", "+4-5*2/10+43", "Not enough operands"},
		{"notEnoughOperands", "/4-5*2/10+43", "Not enough operands"},
		{"notEnoughOperands", "4-5*2/10+43/", "Not enough operands"},
		{"notEnoughOperands", "2++53", "Not enough operands"},
		{"notEnoughOperands", "2+", "Not enough operands"},

		{"notEnoughOperators", "1 2", "Not enough operators"},
		{"notEnoughOperators", "1 2 3", "Not enough operators"},
		{"notEnoughOperators", "1 0 0", "Not enough operators"},
		{"notEnoughOperators", "0 0 0", "Not enough operators"},

		{"divisionByZero", "15 / 0", "Division by zero"},
		{"divisionByZero", "2+55*4/2/0*4", "Division by zero"},
		{"divisionByZero", "11 / (6/6 - 1)", "Division by zero"},
		{"divisionByZero", "(18.6*0.24 + (3.0+2.0/9.0)*(15.0/38.0)) / (7.0/(7.0+1.0/2.0) - ((5.0+1.0/15.0)*3.0-(12.0+19.0/30.0))/2.75)", "Division by zero"},

		{"unknownSymbol", "a", "Unknown symbol"},
		{"unknownSymbol", "a  b  c", "Unknown symbol"},
		{"unknownSymbol", "a + e - 3 .de q", "Unknown symbol"},
		{"unknownSymbol", "1+2/4*55 - i &; s", "Unknown symbol"},

		{"empryExpression", "", "Empty expression"},
		{"empryExpression", " ", "Empty expression"},
		{"empryExpression", "  ", "Empty expression"},
		{"empryExpression", "                ", "Empty expression"},

		{"secondDot", "1.. + 5", "Second dot in number"},
		{"secondDot", "2.20.1 - 4", "Second dot in number"},
		{"secondDot", "..11 * 2", "Second dot in number"},
		{"secondDot", "4 - 2.2.2", "Second dot in number"},
		{"secondDot", "7 / 6.6 / 1.1.1", "Second dot in number"},
	}

	for _, test := range tests {
		_, error := calculations.Calculate(test.input)

		if error.Error() != test.expectedError {
			t.Error("Unexpected error:", error)
		}

	}
}

func TestCalcComplicated(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{"complicated",
			"((1+2)*(3+4))/(2+3)",
			((1.0 + 2.0) * (3.0 + 4.0)) / (2.0 + 3.0)},
		{"complicated",
			"(2.0 + 1.0*4.0 - 8.0*((16.0+14.0*2.0)/(38.0-3.0))) / (4.0 - 5.0*11.0/6.0 + 3.0*5.0)",
			(2.0 + 1.0*4.0 - 8.0*((16.0+14.0*2.0)/(38.0-3.0))) / (4.0 - 5.0*11.0/6.0 + 3.0*5.0)},
		{"complicated",
			"(18.6*0.24 + (3.0+2.0/9.0)*(15.0/38.0)) / (7.0/(7.0+1.0/2.0) - ((5.0+1.0/15.0)*3.0-(12.0+19.0/30.0))/2.7)",
			(18.6*0.24 + (3.0+2.0/9.0)*(15.0/38.0)) / (7.0/(7.0+1.0/2.0) - ((5.0+1.0/15.0)*3.0-(12.0+19.0/30.0))/2.7)},
		{"complicated",
			"3.1/3.0*(3.0*3.0)*3.0/3.0+3.0",
			3.1/3.0*(3.0*3.0)*3.0/3.0 + 3.0},
	}

	for i, test := range tests {
		result, error := calculations.Calculate(test.input)

		if error != nil {
			t.Fatal("Fatal error in calc function, test:", test.name, i, error)
		}

		if !utils.Float64Compare(result, test.expected) {
			t.Errorf("Error. Test: %v, %v. Expected %v, got %v", test.name, i, test.expected, result)
		}
	}
}

func TestCalcSpaces(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{"spaces", "  1 + 2  ", 3},
		{"spaces", "   (  1  + 2  )  *  (  3  + 4  )   ", 21},
		{"spaces", "   2    +  2   * ( 55  + 100     / 50    -  55)", 6},
	}

	for i, test := range tests {
		result, error := calculations.Calculate(test.input)

		if error != nil {
			t.Fatal("Fatal error in calc function, test:", test.name, i, error)
		}

		if !utils.Float64Compare(result, test.expected) {
			t.Errorf("Error. Test: %v, %v. Expected %v, got %v", test.name, i, test.expected, result)
		}
	}
}
