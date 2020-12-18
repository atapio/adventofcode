package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindAnswer(t *testing.T) {
	//t.Parallel()

	testCases := []struct {
		Name  string
		Input []string

		Expected int
	}{
		{
			Name:     "example1",
			Input:    []string{"1 + 2 * 3 + 4 * 5 + 6"},
			Expected: 71,
		},
		{
			Name:     "example1.1",
			Input:    []string{"1 + (2 * 3) + (4 * (5 + 6))"},
			Expected: 51,
		},
		{
			Name: "example2",
			Input: []string{
				"2 * 3 + (4 * 5)",
			},
			Expected: 26,
		},
		{
			Name: "example-all",
			Input: []string{
				"2 * 3 + (4 * 5)",
				"5 + (8 * 3 + 9 + 3 * 4 * 3)",
				"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
				"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			},
			Expected: 26 + 437 + 12240 + 13632,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			//t.Parallel()
			input := tc.Input

			if len(input) == 0 {
				var err error
				input, err = parseFile(fmt.Sprintf("%s.txt", tc.Name))
				require.NoError(t, err)
			}

			got, err := findAnswer(input)
			require.NoError(t, err)

			assert.Equal(t, tc.Expected, got)

		})
	}
}
