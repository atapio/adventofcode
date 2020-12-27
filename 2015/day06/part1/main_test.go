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
			Name: "example1",
			Input: []string{
				"turn on 0,0 through 999,999",
			},
			Expected: 1000000,
		},
		{
			Name: "example2",
			Input: []string{
				"turn on 0,0 through 999,999",
				"toggle 0,0 through 999,0",
			},
			Expected: 1000000 - 1000,
		},
		{
			Name: "example2",
			Input: []string{
				"turn on 0,0 through 999,999",
				"toggle 0,0 through 999,0",
				"turn off 499,499 through 500,500",
			},
			Expected: 1000000 - 1000 - 4,
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
