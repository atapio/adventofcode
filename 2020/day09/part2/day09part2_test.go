package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindAnswer(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name     string
		Input    []string
		Preamble int

		Expected int
	}{
		{
			Name:     "example1",
			Preamble: 5,
			Input: []string{
				"35",
				"20",
				"15",
				"25",
				"47",
				"40",
				"62",
				"55",
				"65",
				"95",
				"102",
				"117",
				"150",
				"182",
				"127",
				"219",
				"299",
				"277",
				"309",
				"576",
			},
			Expected: 62,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			d, err := processInput(tc.Input)
			require.NoError(t, err)

			got, err := findAnswer(d, tc.Preamble)
			require.NoError(t, err)

			assert.Equal(t, tc.Expected, got)

		})
	}
}
