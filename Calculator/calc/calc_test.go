package calc

import (
	"calc/comparator"
	"testing"
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

	allTests := [][]testcase{
		testAddition,
		testSubtraction,
		testMultiplication,
		testDivision,
		testParentheses,
		testPriority,
		testComplicated,
		testSpaces,
	}

	for _, tests := range allTests {

		for i, test := range tests {

			result, err := Calculate(test.input)

			if err != nil {
				t.Logf("Fatal error: %s. Test: %s, № %v", err.Error(), test.name, i)
				t.FailNow()
			}

			if !comparator.Float64Compare(result, test.expected) {
				t.Errorf("Error. Test: %s, %v.\nExpected: %v\nGet:      %v\n", test.name, i, test.expected, result)
			}
		}
	}
}

func TestCalcError(t *testing.T) {

	for _, test := range Errortests {
		_, error := Calculate(test.input)

		if error.Error() != test.expectedError {
			t.Error("Unexpected error:", error)
		}

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
	{"division", "10/3", float64(10.0 / 3.0)},
	{"division", "1763/98", float64(1763.0 / 98.0)},
}

var testParentheses []testcase = []testcase{
	{"parentheses", "(2+2)*2", 8},
	{"parentheses", "(10)", 10},
	{"parentheses", "2*(3+1)*4", 32},
	{"parentheses", "((2+3)*4)/2", 10},
	{"parentheses", "(1+2)-3", 0},
	{"parentheses", "(7-100)", -93},
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

var testSpaces []testcase = []testcase{
	{"spaces", "  1 + 2  ", 3},
	{"spaces", "   (  1  + 2  )  *  (  3  + 4  )   ", 21},
	{"spaces", "   2    +  2   * ( 55  + 100     / 50    -  55)", 6},
}

// Error Test Cases

var Errortests []errorTestCase = []errorTestCase{
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
