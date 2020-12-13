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
				"939",
				"7,13,x,x,59,x,31,19",
			},
			Expected: 295,
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
