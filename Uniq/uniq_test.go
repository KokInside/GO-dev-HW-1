package main

import (
	"slices"
	"testing"
	"uniq/testcases"
	"uniq/uniq"
)

func TestUniq(t *testing.T) {

	allTests := [][]testcases.TestCase{
		testcases.TestsNoParemeters,
		testcases.Tests_c_Flag,
		testcases.Tests_d_Flag,
		testcases.Tests_u_Flag,
		testcases.Tests_i_Flag,
		testcases.Tests_f_Flag,
		testcases.Tests_s_Flag,
	}

	for _, tests := range allTests {

		for i, test := range tests {

			result, err := uniq.Uniq(test.Input, test.Options)

			if err != nil {
				t.Logf("Fatal error: %s. Test: %v, № %v", err.Error(), test.Name, i)
				t.FailNow()
			}

			if !slices.Equal(result, test.Expected) {
				t.Errorf("Error: test: %v, № %v.\nExpected: %v\nGet     : %v\n", test.Name, i, test.Expected, result)
			}
		}
	}
}

func TestUniqError(t *testing.T) {
	allErrorTests := [][]testcases.ErrorTestCase{
		testcases.TestDU_FlagsError,
		testcases.TestNegativeF_FlagError,
		testcases.TestNegativeS_FlagError,
	}

	for _, tests := range allErrorTests {
		for i, test := range tests {

			_, err := uniq.Uniq(test.Input, test.Options)

			if err == nil {
				t.Errorf("Expected error, got nil. Test: %s, № %v", test.Name, i)
			}

			if err.Error() != test.ExpectedError {
				t.Errorf("Expected error: %s, got %s.\nTest: %s, № %v", test.ExpectedError, err.Error(), test.Name, i)
			}
		}
	}
}
