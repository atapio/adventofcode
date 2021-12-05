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

		Expected map[string]uint
	}{
		{
			Name: "example1",
			Expected: map[string]uint{
				"d": 72,
				"e": 507,
				"f": 492,
				"g": 114,
				"h": 65412,
				"i": 65079,
				"x": 123,
				"y": 456,
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			//t.Parallel()
			input := tc.Input

			if len(input) == 0 {
				var err error
				input, err = parseFile(fmt.Sprintf("../%s.txt", tc.Name))
				require.NoError(t, err)
			}

			got, err := findAnswer(input)
			require.NoError(t, err)

			assert.Equal(t, tc.Expected, got)

		})
	}
}
