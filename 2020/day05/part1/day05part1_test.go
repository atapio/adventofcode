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
			Name: "all",
			Input: []string{
				"FBFBBFFRLR",
				"BFFFBBFRRR",
				"FFFBBBFRRR",
				"BBFFBBFRLL",
			},
			Expected: 820,
		},
		{
			Name: "example1",
			Input: []string{
				"FBFBBFFRLR",
			},
			Expected: 357,
		},
		{
			Name: "example2",
			Input: []string{
				"BFFFBBFRRR",
			},
			Expected: 567,
		},
		{
			Name: "example3",
			Input: []string{
				"FFFBBBFRRR",
			},
			Expected: 119,
		},
		{
			Name: "example4",
			Input: []string{
				"BBFFBBFRLL",
			},
			Expected: 820,
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
