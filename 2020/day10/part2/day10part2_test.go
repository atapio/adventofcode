package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
				"16",
				"10",
				"15",
				"5",
				"1",
				"11",
				"7",
				"19",
				"6",
				"12",
				"4",
			},
			Expected: 7,
		},
		{
			Name: "example2",
			Input: []string{
				"28",
				"33",
				"18",
				"42",
				"31",
				"14",
				"46",
				"20",
				"48",
				"47",
				"24",
				"23",
				"49",
				"45",
				"19",
				"38",
				"39",
				"11",
				"1",
				"32",
				"25",
				"35",
				"8",
				"17",
				"7",
				"9",
				"4",
				"2",
				"34",
				"10",
				"3",
			},
			Expected: 19208,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			d, err := processInput(tc.Input)
			require.NoError(t, err)

			got, err := findAnswer(d)
			require.NoError(t, err)

			assert.Equal(t, tc.Expected, got)

		})
	}
}
