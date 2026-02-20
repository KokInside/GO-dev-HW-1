package testcases

import (
	"uniq/options"
)

type ErrorTestCase struct {
	Name          string
	Input         []string
	Options       options.Options
	ExpectedError string
}

var TestDU_FlagsError []ErrorTestCase = []ErrorTestCase{
	{
		Name: "-d and -u flags used simultaneously",
		Input: []string{
			"Some input",
			"Some output",
		},
		Options:       options.Options{D: true, U: true},
		ExpectedError: "Flags -d and -u can't be used simultaneously.",
	},
	{
		Name: "-d and -u flags used simultaneously",
		Input: []string{
			"Some input",
			"Some output",
		},
		Options:       options.Options{D: true, U: true, I: true},
		ExpectedError: "Flags -d and -u can't be used simultaneously.",
	},
	{
		Name: "-d and -u flags used simultaneously",
		Input: []string{
			"Some input",
			"Some output",
		},

		Options:       options.Options{D: true, U: true, I: true, C: true},
		ExpectedError: "Flags -d and -u can't be used simultaneously.",
	},
}

var TestNegativeF_FlagError []ErrorTestCase = []ErrorTestCase{
	{
		Name: "Negative -f flag",
		Input: []string{
			"Some text",
			"some text",
			"one more text",
		},
		Options:       options.Options{F: -11},
		ExpectedError: "[num_fields] must be positive",
	},
	{
		Name: "Negative -f flag",
		Input: []string{
			"Some text",
			"some text",
			"one more text",
		},
		Options:       options.Options{F: -1, C: true, U: true, I: true},
		ExpectedError: "[num_fields] must be positive",
	},
}

var TestNegativeS_FlagError []ErrorTestCase = []ErrorTestCase{
	{
		Name: "Negative -s flag",
		Input: []string{
			"Some text",
			"some text",
			"one more text",
		},
		Options:       options.Options{S: -11},
		ExpectedError: "[num_chars] must be positive",
	},
	{
		Name: "Negative -f flag",
		Input: []string{
			"Some text",
			"some text",
			"one more text",
		},
		Options:       options.Options{S: -1, C: true, D: true},
		ExpectedError: "[num_chars] must be positive",
	},
}
