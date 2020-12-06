package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAnswer(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name  string
		Input []string

		Expected int
	}{
		{
			Name: "example1",
			Input: []string{
				"abcx",
				"abcy",
				"abcz",
			},
			Expected: 6,
		},
		{
			Name: "example1",
			Input: []string{
				"abc",
				"",
				"a",
				"b",
				"c",
				"",
				"ab",
				"ac",
				"",
				"a",
				"a",
				"a",
				"a",
				"",
				"b",
			},
			Expected: 11,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			d := processInput(tc.Input)
			got := findAnswer(d)

			assert.Equal(t, tc.Expected, got)

		})
	}
}
