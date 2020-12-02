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
			Name:     "example1",
			Input:    []string{"1-3 a: abcde"},
			Expected: 1,
		},
		{
			Name:     "example2",
			Input:    []string{"1-3 b: cdefg"},
			Expected: 0,
		},
		{
			Name:     "example3",
			Input:    []string{"2-9 c: ccccccccc"},
			Expected: 0,
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
