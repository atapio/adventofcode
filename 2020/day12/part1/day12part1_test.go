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
				"F10",
				"N3",
				"F7",
				"R90",
				"F11",
			},
			Expected: 25,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			got, err := findAnswer(tc.Input)
			require.NoError(t, err)

			assert.Equal(t, tc.Expected, got)

		})
	}
}
