package uniq

import (
	"testing"
	"uniq/options"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name     string
	input    []string
	expected []string
	options  options.Options
}

type errortestCase struct {
	name          string
	input         []string
	options       options.Options
	expectedError string
}

func TestUniq(t *testing.T) {

	allTests := [][]testCase{
		testsNoParemeters,
		testCountFlag,
		testsRepeatedFlag,
		testsUniqueFlag,
		testsIgnoreCaseFlag,
		testsSkipFieldsFlag,
		testsSkipCharsFlag,

		testsNoParemetersRus,
		testsCountFlagRus,
		testsRepeatedFlagRus,
		testsUniqueFlagRus,
		testsIgnoreCaseFlagRus,
		testsSkipFieldsFlagRus,
		testsSkipCharsFlagRus,
	}

	assert := assert.New(t)

	for _, tests := range allTests {
		for i, test := range tests {
			result, err := Uniq(test.input, test.options)

			if assert.NoErrorf(err, "Fatal error. Test: %v, № %v", test.name, i) {
				assert.Equalf(test.expected, result, "Error: test: %v, № %v.\nExpected: %v\nGot     : %v\n", test.name, i, test.expected, result)
			}
		}
	}
}

func TestUniqError(t *testing.T) {
	allErrorTests := [][]errortestCase{
		testCountRepeatedUniqueFlagsError,
		testSkipFieldsFlagError,
		testSkipCharsFlagError,
	}

	assert := assert.New(t)

	for _, tests := range allErrorTests {
		for i, test := range tests {
			_, err := Uniq(test.input, test.options)

			assert.EqualErrorf(err, test.expectedError, "Expected error: \"%s\", got \"%s\".\nTest: %s, № %v", test.expectedError, err.Error(), test.name, i)
		}
	}
}

// Test cases

var testsNoParemeters []testCase = []testCase{
	{
		name: "noParameters",
		input: []string{
			"Today we gathered moss for my uncle's wedding.",
			"Today we gathered moss for my uncle's wedding.",
			"Today we gathered moss for my uncle's wedding.",
			"hello",
			"Today we gathered moss for my uncle's wedding.",
			"",
			" ",
			"hello",
			"hello",
			"hi",
			"hello",
			"Today we gathered moss for my uncle's wedding.",
		},
		expected: []string{
			"Today we gathered moss for my uncle's wedding.",
			"hello",
			"Today we gathered moss for my uncle's wedding.",
			"",
			" ",
			"hello",
			"hi",
			"hello",
			"Today we gathered moss for my uncle's wedding.",
		},
	},
	{
		name: "noParameters",
		input: []string{
			"a",
			"a",
			"a ",
			"a    ",
			"b",
			"b",
			"a",
			"a",
			"	",
			"	",
			" ",
			"a",
		},
		expected: []string{
			"a",
			"a ",
			"a    ",
			"b",
			"a",
			"	",
			" ",
			"a",
		},
	},
}

var testCountFlag []testCase = []testCase{
	{
		name: "-c flag",
		input: []string{
			"",
			" ",
			"He learned the important lesson that a picnic at the beach on a windy day is a bad idea.",
			"He learned the important lesson that a picnic at the beach on a windy day is a bad idea.",
			"",
			"The book is in front of the table.",
			"The book is in front of the table.",
			"The book is in front of the table.",
			"The book is in front of the table.",
			"The book is in front of the table.",
			"",
			"",
			"",
			"",
			" ",
			" ",
			" ",
			" ",
			"hello",
			"hi",
			"hello",
			"hi",
			"hello",
			"",
		},
		expected: []string{
			"    1 ",
			"    1  ",
			"    2 He learned the important lesson that a picnic at the beach on a windy day is a bad idea.",
			"    1 ",
			"    5 The book is in front of the table.",
			"    4 ",
			"    4  ",
			"    1 hello",
			"    1 hi",
			"    1 hello",
			"    1 hi",
			"    1 hello",
			"    1 ",
		},
		options: options.Options{Count: true},
	},
	{
		name: "-c flag",
		input: []string{
			"alpha beta gamma",
			"alpha beta gamma",
			"ALPHA beta gamma",
			"alpha  beta gamma",
			"alpha beta gamma",
			"alpha beta gamma",
			"",
			"",
			" ",
			"",
			"omega",
			"omega",
		},
		expected: []string{
			"    2 alpha beta gamma",
			"    1 ALPHA beta gamma",
			"    1 alpha  beta gamma",
			"    2 alpha beta gamma",
			"    2 ",
			"    1  ",
			"    1 ",
			"    2 omega",
		},
		options: options.Options{Count: true},
	},
	{
		name: "-c flag",
		input: []string{
			"line",
			"line",
			"line",
			"line ",
			"line ",
			"   ",
			"   ",
			"line",
			"line",
		},
		expected: []string{
			"    3 line",
			"    2 line ",
			"    2    ",
			"    2 line",
		},
		options: options.Options{Count: true},
	},
	{
		name: "-c flag",
		input: []string{
			"       ",
			"       ",
			"	",
			" ",
			"single",
			"single",
			"	 ",
		},
		expected: []string{
			"    2        ",
			"    1 	",
			"    1  ",
			"    2 single",
			"    1 	 ",
		},
		options: options.Options{Count: true},
	},
}

var testsRepeatedFlag []testCase = []testCase{
	{
		name: "-d flag",
		input: []string{
			"",
			" ",
			"He learned the important lesson that a picnic at the beach on a windy day is a bad idea.",
			"He learned the important lesson that a picnic at the beach on a windy day is a bad idea.",
			"",
			"The book is in front of the table.",
			"The book is in front of the table.",
			"The book is in front of the table.",
			"The book is in front of the table.",
			"The book is in front of the table.",
			"",
			"",
			"",
			"",
			" ",
			" ",
			" ",
			" ",
			"hello",
			"hi",
			"hello",
			"hi",
			"hello",
			"",
		},
		expected: []string{
			"He learned the important lesson that a picnic at the beach on a windy day is a bad idea.",
			"The book is in front of the table.",
			"",
			" ",
		},
		options: options.Options{Repeated: true},
	},
	{
		name: "-d flag",
		input: []string{
			"one",
			"one",
			"two",
			"three",
			"two",
			"two",
			"two",
			"One",
			"one",
		},
		expected: []string{
			"one",
			"two",
		},
		options: options.Options{Repeated: true},
	},
	{
		name: "-d flag",
		input: []string{
			"x",
			"x",
			"y",
			"z",
			"z",
			"z",
			"y",
			"y",
			"x",
			"x",
		},
		expected: []string{
			"x",
			"z",
			"y",
			"x",
		},
		options: options.Options{Repeated: true},
	},
}

var testsUniqueFlag []testCase = []testCase{
	{
		name: "-u flag",
		input: []string{
			"",
			" ",
			"He learned the important lesson that a picnic at the beach on a windy day is a bad idea.",
			"He learned the important lesson that a picnic at the beach on a windy day is a bad idea.",
			"",
			"The book is in front of the table.",
			"The book is in front of the table.",
			"The book is in front of the table.",
			"The book is in front of the table.",
			"The book is in front of the table.",
			"",
			"",
			"",
			"",
			" ",
			" ",
			" ",
			" ",
			"hello",
			"hi",
			"hello",
			"hi",
			"hello",
			"",
		},
		expected: []string{
			"",
			" ",
			"",
			"hello",
			"hi",
			"hello",
			"hi",
			"hello",
			"",
		},
		options: options.Options{Unique: true},
	},
	{
		name: "-u flag",
		input: []string{
			"first",
			"second",
			"second",
			"third",
			"fourth",
			"fourth",
			"fifth",
			"second",
		},
		expected: []string{
			"first",
			"third",
			"fifth",
			"second",
		},
		options: options.Options{Unique: true},
	},
	{
		name: "-u flag",
		input: []string{
			"apple",
			"banana",
			"apple",
			"banana",
			"cherry",
			"cherry",
			"apple",
		},
		expected: []string{
			"apple",
			"banana",
			"apple",
			"banana",
			"apple",
		},
		options: options.Options{Unique: true},
	},
}

var testsIgnoreCaseFlag []testCase = []testCase{
	{
		name: "-i flag",
		input: []string{
			"",
			" ",
			"he LeARneD The iMPORtANt LessON tHaT a Picnic At tHE BEaCh on a WINDY day iS A bAD IDEa.",
			"he learnED tHE ImPoRtant lEsSON tHaT A piCNic aT tHE beACH on a windY dAY Is a bad IdEa.",
			"",
			"THe BoOk Is iN FroNt Of tHe TABlE.",
			"THe booK is in frOnt of tHE TabLe.",
			"tHE BOoK is in FRoNT oF tHe TAbLe.",
			"the boOk is iN frONT Of tHe tAblE.",
			"tHE BoOk IS iN front OF the tablE.",
			"",
			"i LIked THeiR fiRst tWO albUmS BuT CHAnged my Mind aFTeR THat CHaRITy gIG.",
			" ",
			"hELLO",
			"hELLO",
			"hi",
			"hello",
			"i liked THEiR fIrst TWO aLBuMS BUt cHANgED mY miND afteR thAt cHARity gig.",
			" ",
			" ",
			"hEllO",
			"hi",
			"HelLO",
			"hI",
			"hELlo",
			"",
		},
		expected: []string{
			"",
			" ",
			"he LeARneD The iMPORtANt LessON tHaT a Picnic At tHE BEaCh on a WINDY day iS A bAD IDEa.",
			"",
			"THe BoOk Is iN FroNt Of tHe TABlE.",
			"",
			"i LIked THeiR fiRst tWO albUmS BuT CHAnged my Mind aFTeR THat CHaRITy gIG.",
			" ",
			"hELLO",
			"hi",
			"hello",
			"i liked THEiR fIrst TWO aLBuMS BUt cHANgED mY miND afteR thAt cHARity gig.",
			" ",
			"hEllO",
			"hi",
			"HelLO",
			"hI",
			"hELlo",
			"",
		},
		options: options.Options{IgnoreCase: true},
	},
}

var testsSkipFieldsFlag []testCase = []testCase{
	{
		name: "-f flag",
		input: []string{
			"12.03.25 Today we gathered moss for my uncle's wedding.",
			"13.03.25 Today we gathered moss for my uncle's wedding.",
			"14.03.25 Today we gathered moss for my uncle's wedding.",
			"15.03.25 hello",
			"16.03.25 Today we gathered moss for my uncle's wedding.",
			"17.03.25 ",
			"18.03.25  ",
			"19.03.25 hello",
			"20.03.25 hello",
			"21.03.25 hi",
			"22.03.25 hello",
			"23.03.25 Today we gathered moss for my uncle's wedding.",
		},
		expected: []string{
			"12.03.25 Today we gathered moss for my uncle's wedding.",
			"15.03.25 hello",
			"16.03.25 Today we gathered moss for my uncle's wedding.",
			"17.03.25 ",
			"18.03.25  ",
			"19.03.25 hello",
			"21.03.25 hi",
			"22.03.25 hello",
			"23.03.25 Today we gathered moss for my uncle's wedding.",
		},
		options: options.Options{SkipFields: 1},
	},
	{
		name: "-f flag",
		input: []string{
			"2026-02-17 INFO started process A",
			"2026-02-17 INFO started process A",
			"2026-02-17 DEBUG started process A",
			"2026-02-17 INFO started process B",
			"2026-02-18 INFO started process A",
			"INFO started process A",
		},
		expected: []string{
			"2026-02-17 INFO started process A",
			"2026-02-17 DEBUG started process A",
			"2026-02-17 INFO started process B",
			"2026-02-18 INFO started process A",
			"INFO started process A",
		},
		options: options.Options{SkipFields: 1},
	},
	{
		name: "-f flag",
		input: []string{
			"A BB C",
			"UO BBB C",
			"XYZ 11Y Z",
			"OOO 9 Z",
		},
		expected: []string{
			"A BB C",
			"XYZ 11Y Z",
		},
		options: options.Options{SkipFields: 2},
	},
}

var testsSkipCharsFlag []testCase = []testCase{
	{
		name: "-s flag",
		input: []string{
			"ABCfoobar",
			"ABDfoobar",
			"ABEfoobar",
			"ABCfoobar",
			"XYZfoobar",
			"XYafoobar",
		},
		expected: []string{
			"ABCfoobar",
		},
		options: options.Options{SkipChars: 3},
	},
}

// Rus tests

var testsNoParemetersRus []testCase = []testCase{
	{
		name: "noParametersRus",
		input: []string{
			"яблоко",
			"яблоко",
			"Яблоко",
			"яблоко",
			"банан",
			"банан",
			"",
			"",
			"груша",
		},
		expected: []string{
			"яблоко",
			"Яблоко",
			"яблоко",
			"банан",
			"",
			"груша",
		},
	},
	{
		name: "noParametersRus",
		input: []string{
			" пробел",
			"пробел",
			"пробел ",
			"пробел ",
			" пробел",
		},
		expected: []string{
			" пробел",
			"пробел",
			"пробел ",
			" пробел",
		},
	},
}

var testsCountFlagRus []testCase = []testCase{
	{
		name: "-c flag Rus",
		input: []string{
			"раз",
			"раз",
			"два",
			"три",
			"три",
			"три",
			"два",
		},
		expected: []string{
			"    2 раз",
			"    1 два",
			"    3 три",
			"    1 два",
		},
		options: options.Options{Count: true},
	},
	{
		name: "-c flag Rus",
		input: []string{
			"Тест",
			"тест",
			"тест",
			"ТЕСТ",
			"тест",
		},
		expected: []string{
			"    1 Тест",
			"    2 тест",
			"    1 ТЕСТ",
			"    1 тест",
		},
		options: options.Options{Count: true},
	},
}

var testsRepeatedFlagRus []testCase = []testCase{
	{
		name: "-d flag Rus",
		input: []string{
			"яблоко",
			"яблоко",
			"банан",
			"банан",
			"банан",
			"вишня",
			"вишня",
			"яблоко",
			"банан",
			"киви",
		},
		expected: []string{
			"яблоко",
			"банан",
			"вишня",
		},
		options: options.Options{Repeated: true},
	},
	{
		name: "-d flag Rus",
		input: []string{
			"яблоко",
			"яблоко",
			"банан",
			"банан",
			"банан",
			"вишня",
			"груша",
			"груша",
			"дыня",
		},
		expected: []string{
			"яблоко",
			"банан",
			"груша",
		},
		options: options.Options{Repeated: true},
	},
}

var testsUniqueFlagRus []testCase = []testCase{
	{
		name: "-u flag Rus",
		input: []string{
			"один",
			"один",
			"два",
			"три",
			"три",
			"четыре",
			"пять",
			"пять",
		},
		expected: []string{
			"два",
			"четыре",
		},
		options: options.Options{Unique: true},
	},
	{
		name: "-u flag Rus",
		input: []string{
			"яблоко",
			"яблоко",
			"банан",
			"банан",
			"банан",
			"вишня",
			"вишня",
			"яблоко",
			"банан",
			"киви",
			"манго",
			"манго",
		},
		expected: []string{
			"яблоко",
			"банан",
			"киви",
		},
		options: options.Options{Unique: true},
	},
}

var testsIgnoreCaseFlagRus []testCase = []testCase{
	{
		name: "-i flag Rus",
		input: []string{
			"Яблоко",
			"яблоко",
			"Банан",
			"БАНАН",
			"банан",
			"Вишня",
			"ВИШНЯ",
			"яблоко",
			"банан",
		},
		expected: []string{
			"Яблоко",
			"Банан",
			"Вишня",
			"яблоко",
			"банан",
		},
		options: options.Options{IgnoreCase: true},
	},
}

var testsSkipFieldsFlagRus []testCase = []testCase{
	{
		name: "-f flag Rus",
		input: []string{
			"1 альфа бета",
			"2 альфа бета",
			"3 Альфа Бета",
			"4 гамма",
			"4 гамма",
		},
		expected: []string{
			"1 альфа бета",
		},
		options: options.Options{SkipFields: 300},
	},
}

var testsSkipCharsFlagRus []testCase = []testCase{
	{
		name: "-s flag Rus",
		input: []string{
			"2025-01-01 яблоко",
			"2024-12-31 яблоко",
			"2025-01-01 Яблоко",
			"2025-01-02 банан",
			"2025-01-02 банан",
		},
		expected: []string{
			"2025-01-01 яблоко",
			"2025-01-02 банан",
		},
		options: options.Options{SkipChars: 15},
	},
}

// Error cases

var testCountRepeatedUniqueFlagsError []errortestCase = []errortestCase{
	{
		name: "Unique -c, -d, -u",
		input: []string{
			"Some input",
			"Some output",
		},
		options:       options.Options{Repeated: true, Unique: true},
		expectedError: "only one flag of -c, -d and -u can be used",
	},
	{
		name: "Unique -c, -d, -u",
		input: []string{
			"Some input",
			"Some output",
		},
		options:       options.Options{Count: true, Repeated: true, Unique: true, IgnoreCase: true},
		expectedError: "only one flag of -c, -d and -u can be used",
	},
	{
		name: "Unique -c, -d, -u",
		input: []string{
			"Some input",
			"Some output",
		},

		options:       options.Options{Repeated: true, Unique: true, IgnoreCase: true, Count: true},
		expectedError: "only one flag of -c, -d and -u can be used",
	},
}

var testSkipFieldsFlagError []errortestCase = []errortestCase{
	{
		name: "Negative -f flag",
		input: []string{
			"Some text",
			"some text",
			"one more text",
			"what text ?",
		},
		options:       options.Options{SkipFields: -11},
		expectedError: "[num_fields] must be positive",
	},
	{
		name: "Negative -f flag",
		input: []string{
			"Some text",
			"some text",
			"one more text",
			"test textxxx",
		},
		options:       options.Options{SkipFields: -1, Count: true, IgnoreCase: true},
		expectedError: "[num_fields] must be positive",
	},
}

var testSkipCharsFlagError []errortestCase = []errortestCase{
	{
		name: "Negative -s flag",
		input: []string{
			"Some text",
			"some text",
			"one more text",
		},
		options:       options.Options{SkipChars: -11},
		expectedError: "[num_chars] must be positive",
	},
	{
		name: "Negative -s flag",
		input: []string{
			"Some text",
			"some text",
			"one more text",
		},
		options:       options.Options{SkipChars: -1, Repeated: true},
		expectedError: "[num_chars] must be positive",
	},
}
