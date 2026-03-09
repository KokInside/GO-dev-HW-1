package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	name     string
	input    string
	expected float64
}

type errorTestCase struct {
	name          string
	input         string
	expectedError string
}

func TestCalc(t *testing.T) {
	noexceptTests := [][]testcase{
		testAddition,
		testSubtraction,
		testMultiplication,
		testDivision,
		testUnary,
		tesеImplicitMultiplication,
		testParentheses,
		testPriority,
		testComplicated,
		testSpaces,
	}

	assert := assert.New(t)

	// Not error test cases
	for _, tests := range noexceptTests {
		for i, test := range tests {
			result, err := Calculate(test.input)

			if err != nil {
				t.Logf("Fatal error: %s. Test: %s, № %v", err.Error(), test.name, i)
				t.FailNow()
			}
			assert.InDeltaf(test.expected, result, 1.0e-9, "Error. Test: %s, %v.\nExpected: %v\nGot:      %v\n", test.name, i, test.expected, result)
		}
	}

	// Error test cases
	for i, test := range errortests {
		_, err := Calculate(test.input)

		assert.EqualErrorf(err, test.expectedError, "Expected error: \"%s\", got: nil. Test: %s, № %v", test.expectedError, test.name, i)
	}
}

// Test Cases

var testAddition []testcase = []testcase{
	{"addition", "2+55", 57},
	{"addition", "2+4+6", 12},
	{"addition", "0+0", 0},
	{"addition", "2+0+9+0", 11},
}

var testSubtraction []testcase = []testcase{
	{"subtraction", "10-1", 9},
	{"subtraction", "10-3-4", 3},
	{"subtraction", "0-0", 0},
	{"subtraction", "10-101", -91},
}

var testMultiplication []testcase = []testcase{
	{"multiplication", "3*4", 12},
	{"multiplication", "3*0", 0},
	{"multiplication", "0*4", 0},
	{"multiplication", "0*0", 0},
	{"multiplication", "199*1", 199},
	{"multiplication", "2*3*4", 24},
	{"multiplication", "1*2*3*4*5*6*0", 0},
	{"multiplication", "1*2*3*4*5", 120},
}

var testDivision []testcase = []testcase{
	{"division", "6/2", 3},
	{"division", "120/5", 24},
	{"division", "103/1", 103},
	{"division", "20/2/5", 2},
	{"division", "1/1", 1},
	{"division", "0/28", 0},
	{"division", "10/3", 10.0 / 3.0},
	{"division", "2/3", 2.0 / 3.0},
	{"division", "1763/98", 1763.0 / 98.0},
}

var testParentheses []testcase = []testcase{
	{"parentheses", "(2+2)*2", 8},
	{"parentheses", "(10)", 10},
	{"parentheses", "2*(3+1)*4", 32},
	{"parentheses", "((2+3)*4)/2", 10},
	{"parentheses", "(1+2)-3", 0},
	{"parentheses", "(7-100)", -93},
	{"parentheses", "(-1)*(-1)", 1.0},
	{"parentheses", "2*(3+4)", 14.0},
}

var testUnary []testcase = []testcase{
	{"unary", "-4", -4.0},
	{"unary", "+6", 6.0},
	{"unary", "-3-5", -8.0},
	{"unary", "--1", 1.0},
	{"unary", "+-+55-+100", -155.0},
	{"unary", "-(1+2)", -3.0},
	{"unary", "----10", 10.0},
	{"unary", "-(-5)", 5.0},
	{"unary", "1+-2", -1.0},
	{"unary", "2-+1", 1.0},
	{"unary", "4--3", 7.0},
	{"unary", "7++3", 10.0},
	{"unary", "2*-3", -6.0},
	{"unary", "-2*-3", 6.0},
	{"unary", "6/-2", -3.0},
	{"unary", "-(2+3)*4", -20.0},
	{"unary", "-2/3", -2.0 / 3.0},
	{"unary", "+4-5*2/10+43", 46.0},
}

var tesеImplicitMultiplication []testcase = []testcase{
	{"implicitMultiplacation", "(-1)(-1)", 1.0},
	{"implicitMultiplacation", "2(3+4)", 14.0},
	{"implicitMultiplacation", "(2+3)(4+5)", 45.0},
	{"implicitMultiplacation", "(2+3)5", 25.0},
	{"implicitMultiplacation", "5(2+3)", 25.0},
	{"implicitMultiplacation", "(-2)(3)", -6.0},
	{"implicitMultiplacation", "(-2)(-3)", 6.0},
	{"implicitMultiplacation", "-(2)(3)", -6.0},
	{"implicitMultiplacation", "-(2+3)(4)", -20.0},
}

var testPriority []testcase = []testcase{
	{"priority", "2+2*2", 6},
	{"priority", "2*2+2", 6},
	{"priority", "24/6*10", 40},
	{"priority", "24*6/10", float64(24.0 * 6.0 / 10.0)},
}

var testComplicated []testcase = []testcase{
	{"complicated",
		"((1+2)*(3+4))/(2+3)",
		4.2},
	{"complicated",
		"(2 + 1*4 - 8*((16+14*2)/(38-3))) / (4 - 5*11/6 + 3*5)",
		(2.0 + 1.0*4.0 - 8.0*((16.0+14.0*2.0)/(38.0-3.0))) / (4.0 - 5.0*11.0/6.0 + 3.0*5.0)},
	{"complicated",
		"(18.6*0.24 + (3+2/9)*(15/38)) / (7/(7+1/2) - ((5+1/15)*3-(12+19/30))/2.7)",
		(18.6*0.24 + (3.0+2.0/9.0)*(15.0/38.0)) / (7.0/(7.0+1.0/2.0) - ((5.0+1.0/15.0)*3.0-(12.0+19.0/30.0))/2.7)},
	{"complicated",
		"3.1/3*(3*3)*3/3+3",
		3.1/3.0*(3.0*3.0)*3.0/3.0 + 3.0},
	{"complicated", "-(-11-(1*20/2)-11/2*3)", 37.5},
	{"complicated", "(-3.5+4.2*(2.8--5.1)/1.5-(6.3/-2.1)*(-3+7))*((-(1.2+3.4)*2.5--4.8/2.2)+((5.6--2.3)*-3.1)/(-2*1.7))--(2.5*-2.1)+(-3.7)*(-4.2)/((8.1-3.2*2)+-5.4)+(-(9.3/-3.1))*(2.2+-1.1)-((7.5*-0.5)-(-2.2*3.3))/(-1.5)+--3.14*(-2.71)/(1.5--2.5)-((-5)*(3.2-1.1))+---2.5-(-3.5+4.2*(2.8--5.1)/1.5-(6.3/-2.1)*(-3+7))*((-(1.2+3.4)*2.5--4.8/2.2)+((5.6--2.3)*-3.1)/(-2*1.7))--(2.5*-2.1)+(-3.7)*(-4.2)/((8.1-3.2*2)+-5.4)+(-(9.3/-3.1))*(2.2+-1.1)-((7.5*-0.5)-(-2.2*3.3))/(-1.5)+--3.14*(-2.71)/(1.5--2.5)-((-5)*(3.2-1.1))+---2.5", 4.1253},
	{"complicated", "((-3.5+4.2*(2.8--5.1)/1.5 - (6.3/-2.1)*(-3+7) + ( -(1.2+3.4)*2.5 - -4.8/2.2 ) * ( (5.6--2.3)*-3.1 / (-2*1.7) ) - - (2.5*-2.1) + (-3.7)*(-4.2)/((8.1-3.2*2)+-5.4) + (-(9.3/-3.1))*(2.2+-1.1) - ((7.5*-0.5)-(-2.2*3.3))/(-1.5) + --3.14*(-2.71)/(1.5--2.5) - ((-5)*(3.2-1.1)) + ---2.5 * (3.14 - 1.59) / ( -2.71 + 1.1 ) - ( -1.23 * 4.56 ) / ( 7.89 - -2.34 ) + ( ( -2.5 * 3.7 ) - ( 4.8 / -1.2 ) ) * ( - ( 6.7 - 2.3 ) ) - ( - ( 8.9 / 1.5 ) ) * ( 2.3 * -4.5 ) + ( -3.14 * 2.71 ) / ( -1.5 + 2.5 ) - ( ( -7.5 * 0.5 ) - ( -2.2 * 3.3 ) ) / ( -1.5 ) + --3.14 * ( -2.71 ) / ( 1.5 - -2.5 ) - ( -5 * ( 3.2 - 1.1 ) ) + ---2.5 + ( 1.23 * 4.56 ) / ( 7.89 + 2.34 ) - ( 3.14 * 2.71 ) * ( -1.5 ) + ( 5.6 / -2.3 ) - ( -3.7 * 4.2 ) + ( ( 8.1 - 3.2 ) * 2 ) / -1.5 - ( - ( 9.3 / 3.1 ) ) * ( 2.2 - 1.1 ) + ( 7.5 * -0.5 ) - ( -2.2 * 3.3 ) / ( -1.5 ) + 3.14 * -2.71 / ( 1.5 - -2.5 ) - ( -5 * 3.2 ) + ---2.9 * ( 2.71 - 1.23 ) / ( -3.14 + 1.5 ) - ( -4.56 / 1.2 ) + ( 6.78 * -0.5 ) / ( -1.23 + 2.34 ) - ( ( 3.14 * 2.71 ) - ( 1.59 / 0.5 ) ) * ( - ( 7.8 - 3.2 ) ) + ( - ( 9.1 / 1.7 ) ) * ( 2.2 * -1.1 ) - ( -3.3 * 4.4 ) / ( 5.5 - -2.2 ) + ( ( -6.6 * 1.1 ) - ( 2.2 / -0.5 ) ) * ( - ( 3.3 - 1.1 ) ) - ( - ( 4.4 / 1.1 ) ) * ( 1.2 * -3.3 ) + ( -2.2 * 3.3 ) / ( -1.1 + 2.2 ) - ( ( -7.7 * 0.7 ) - ( -1.1 * 2.2 ) ) / ( -1.1 ) + --2.2 * ( -3.3 ) / ( 1.1 - -2.2 ) - ( -2 * ( 3.3 - 1.1 ) ) + ---2.2) - ((-3.5+4.2*(2.8--5.1)/1.5 - (6.3/-2.1)*(-3+7) + ( -(1.2+3.4)*2.5 - -4.8/2.2 ) * ( (5.6--2.3)*-3.1 / (-2*1.7) ) - - (2.5*-2.1) + (-3.7)*(-4.2)/((8.1-3.2*2)+-5.4) + (-(9.3/-3.1))*(2.2+-1.1) - ((7.5*-0.5)-(-2.2*3.3))/(-1.5) + --3.14*(-2.71)/(1.5--2.5) - ((-5)*(3.2-1.1)) + ---2.5 * (3.14 - 1.59) / ( -2.71 + 1.1 ) - ( -1.23 * 4.56 ) / ( 7.89 - -2.34 ) + ( ( -2.5 * 3.7 ) - ( 4.8 / -1.2 ) ) * ( - ( 6.7 - 2.3 ) ) - ( - ( 8.9 / 1.5 ) ) * ( 2.3 * -4.5 ) + ( -3.14 * 2.71 ) / ( -1.5 + 2.5 ) - ( ( -7.5 * 0.5 ) - ( -2.2 * 3.3 ) ) / ( -1.5 ) + --3.14 * ( -2.71 ) / ( 1.5 - -2.5 ) - ( -5 * ( 3.2 - 1.1 ) ) + ---2.5 + ( 1.23 * 4.56 ) / ( 7.89 + 2.34 ) - ( 3.14 * 2.71 ) * ( -1.5 ) + ( 5.6 / -2.3 ) - ( -3.7 * 4.2 ) + ( ( 8.1 - 3.2 ) * 2 ) / -1.5 - ( - ( 9.3 / 3.1 ) ) * ( 2.2 - 1.1 ) + ( 7.5 * -0.5 ) - ( -2.2 * 3.3 ) / ( -1.5 ) + 3.14 * -2.71 / ( 1.5 - -2.5 ) - ( -5 * 3.2 ) + ---2.5 * ( 2.71 - 1.23 ) / ( -3.14 + 1.5 ) - ( -4.56 / 1.2 ) + ( 6.78 * -0.5 ) / ( -1.23 + 2.34 ) - ( ( 3.14 * 2.71 ) - ( 1.59 / 0.5 ) ) * ( - ( 7.8 - 3.2 ) ) + ( - ( 9.1 / 1.7 ) ) * ( 2.2 * -1.1 ) - ( -3.3 * 4.4 ) / ( 5.5 - -2.2 ) + ( ( -6.6 * 1.1 ) - ( 2.2 / -0.5 ) ) * ( - ( 3.3 - 1.1 ) ) - ( - ( 4.4 / 1.1 ) ) * ( 1.2 * -3.3 ) + ( -2.2 * 3.3 ) / ( -1.1 + 2.2 ) - ( ( -7.7 * 0.7 ) - ( -1.1 * 2.2 ) ) / ( -1.1 ) + --2.2 * ( -3.3 ) / ( 1.1 - -2.2 ) - ( -2 * ( 3.3 - 1.1 ) ) + ---2.2)))", 0.36097561},
}

var testSpaces []testcase = []testcase{
	{"spaces", "  1 + 2  ", 3},
	{"spaces", "   (  1  + 2  )  *  (  3  + 4  )   ", 21},
	{"spaces", "   2    +  2   * ( 55  + 100     / 50    -  55)", 6},
}

// Error Test Cases
var errortests []errorTestCase = []errorTestCase{
	{"unbalancedParentheses", "(7*4+9)+87)", "missing ("},
	{"unbalancedParentheses", ")(5+7)/443+73", "missing ("},
	{"unbalancedParentheses", "(54*3-8", "missing )"},
	{"unbalancedParentheses", "7+322)", "missing ("},
	{"unbalancedParentheses", ")443+22(", "missing ("},

	{"notEnoughOperands", "+", "not enough operands"},
	{"notEnoughOperands", "//", "not enough operands"},

	{"notEnoughOperands", "/4-5*2/10+43", "not enough operands"},
	{"notEnoughOperands", "4-5*2/10+43/", "not enough operands"},
	{"notEnoughOperands", "2+", "not enough operands"},

	{"notEnoughOperators", "1 2", "not enough operators"},
	{"notEnoughOperators", "1 2 3", "not enough operators"},
	{"notEnoughOperators", "1 0 0", "not enough operators"},
	{"notEnoughOperators", "0 0 0", "not enough operators"},

	{"divisionByZero", "15 / 0", "division by zero"},
	{"divisionByZero", "2+55*4/2/0*4", "division by zero"},
	{"divisionByZero", "11 / (6/6 - 1)", "division by zero"},
	{"divisionByZero", "(18.6*0.24 + (3.0+2.0/9.0)*(15.0/38.0)) / (7.0/(7.0+1.0/2.0) - ((5.0+1.0/15.0)*3.0-(12.0+19.0/30.0))/2.75)", "division by zero"},

	{"unknownSymbol", "a", "unknown symbol"},
	{"unknownSymbol", "a  b  c", "unknown symbol"},
	{"unknownSymbol", "a + e - 3 .de q", "unknown symbol"},
	{"unknownSymbol", "1+2/4*55 - i &; s", "unknown symbol"},

	{"empryExpression", "", "empty expression"},
	{"empryExpression", " ", "empty expression"},
	{"empryExpression", "  ", "empty expression"},
	{"empryExpression", "                ", "empty expression"},

	{"secondDot", "1.. + 5", "second dot in number"},
	{"secondDot", "2.20.1 - 4", "second dot in number"},
	{"secondDot", "..11 * 2", "second dot in number"},
	{"secondDot", "4 - 2.2.2", "second dot in number"},
	{"secondDot", "7 / 6.6 / 1.1.1", "second dot in number"},
}
