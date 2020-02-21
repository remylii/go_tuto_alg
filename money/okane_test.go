package main

import (
	"testing"
)

func TestOkaneSuccess(t *testing.T) {
	cases := okaneSuccessProvidor()

	for _, testCase := range cases {
		res := okane(testCase.Argument)

		for key, expectedValue := range testCase.Expected {
			value, ok := res[key]
			if ok == false {
				t.Fatalf("failed. key: %s doesn't exists.", key)
			}

			if value != expectedValue {
				t.Fatalf("failed. key: `%s` is `%d`. should be returned `%d`.", key, value, expectedValue)
			}
		}
	}
}

type OkaneCase struct {
	Argument int
	Expected map[string]int
}

func okaneSuccessProvidor() []OkaneCase {
	cases := []OkaneCase{
		{
			Argument: 19999,
			Expected: map[string]int{"10000": 1, "5000": 1, "1000": 4, "500": 1, "100": 4, "50": 1, "10": 4, "5": 1, "1": 4},
		},
		{
			Argument: 10001,
			Expected: map[string]int{"10000": 1, "5000": 0, "1000": 0, "500": 0, "100": 0, "50": 0, "10": 0, "5": 0, "1": 1},
		},
		{
			Argument: 999999,
			Expected: map[string]int{"10000": 99, "5000": 1, "1000": 4, "500": 1, "100": 4, "50": 1, "10": 4, "5": 1, "1": 4},
		},
	}

	return cases
}
