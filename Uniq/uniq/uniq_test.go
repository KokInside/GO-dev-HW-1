package uniq

import (
	"slices"
	"testing"
	"uniq/options"
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
		tests_c_Flag,
		tests_d_Flag,
		tests_u_Flag,
		tests_i_Flag,
		tests_f_Flag,
		tests_s_Flag,
	}

	for _, tests := range allTests {

		for i, test := range tests {

			result, err := Uniq(test.input, test.options)

			if err != nil {
				t.Logf("Fatal error: %s. Test: %v, № %v", err.Error(), test.name, i)
				t.FailNow()
			}

			if !slices.Equal(result, test.expected) {
				t.Errorf("Error: test: %v, № %v.\nexpected: %v\nGet     : %v\n", test.name, i, test.expected, result)
			}
		}
	}
}

func TestUniqError(t *testing.T) {
	allErrorTests := [][]errortestCase{
		testDU_FlagsError,
		testNegativeF_FlagError,
		testNegativeS_FlagError,
	}

	for _, tests := range allErrorTests {
		for i, test := range tests {

			_, err := Uniq(test.input, test.options)

			if err == nil {
				t.Errorf("expected error, got nil. Test: %s, № %v", test.name, i)
			}

			if err.Error() != test.expectedError {
				t.Errorf("expected error: %s, got %s.\nTest: %s, № %v", test.expectedError, err.Error(), test.name, i)
			}
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

var tests_c_Flag []testCase = []testCase{
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
		options: options.Options{C: true},
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
		options: options.Options{C: true},
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
		options: options.Options{C: true},
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
		options: options.Options{C: true},
	},
}

var tests_d_Flag []testCase = []testCase{
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
		options: options.Options{D: true},
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
		options: options.Options{D: true},
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
		options: options.Options{D: true},
	},
}

var tests_u_Flag []testCase = []testCase{
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
		options: options.Options{U: true},
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
		options: options.Options{U: true},
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
		options: options.Options{U: true},
	},
}

var tests_i_Flag []testCase = []testCase{
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
		options: options.Options{I: true},
	},
}

var tests_f_Flag []testCase = []testCase{
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
		options: options.Options{F: 1},
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
		options: options.Options{F: 1},
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
		options: options.Options{F: 2},
	},
}

var tests_s_Flag []testCase = []testCase{
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
		options: options.Options{S: 3},
	},
}

// Error cases

var testDU_FlagsError []errortestCase = []errortestCase{
	{
		name: "-d and -u flags used simultaneously",
		input: []string{
			"Some input",
			"Some output",
		},
		options:       options.Options{D: true, U: true},
		expectedError: "Flags -d and -u can't be used simultaneously.",
	},
	{
		name: "-d and -u flags used simultaneously",
		input: []string{
			"Some input",
			"Some output",
		},
		options:       options.Options{D: true, U: true, I: true},
		expectedError: "Flags -d and -u can't be used simultaneously.",
	},
	{
		name: "-d and -u flags used simultaneously",
		input: []string{
			"Some input",
			"Some output",
		},

		options:       options.Options{D: true, U: true, I: true, C: true},
		expectedError: "Flags -d and -u can't be used simultaneously.",
	},
}

var testNegativeF_FlagError []errortestCase = []errortestCase{
	{
		name: "Negative -f flag",
		input: []string{
			"Some text",
			"some text",
			"one more text",
		},
		options:       options.Options{F: -11},
		expectedError: "[num_fields] must be positive",
	},
	{
		name: "Negative -f flag",
		input: []string{
			"Some text",
			"some text",
			"one more text",
		},
		options:       options.Options{F: -1, C: true, U: true, I: true},
		expectedError: "[num_fields] must be positive",
	},
}

var testNegativeS_FlagError []errortestCase = []errortestCase{
	{
		name: "Negative -s flag",
		input: []string{
			"Some text",
			"some text",
			"one more text",
		},
		options:       options.Options{S: -11},
		expectedError: "[num_chars] must be positive",
	},
	{
		name: "Negative -f flag",
		input: []string{
			"Some text",
			"some text",
			"one more text",
		},
		options:       options.Options{S: -1, C: true, D: true},
		expectedError: "[num_chars] must be positive",
	},
}
