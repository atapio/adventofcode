package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuelNeeded(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name  string
		Input int

		Expected int
	}{
		{
			Name:     "14",
			Input:    14,
			Expected: 2,
		},
		{
			Name:     "1969",
			Input:    1969,
			Expected: 966,
		},
		{
			Name:     "100756",
			Input:    100756,
			Expected: 50346,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			got := fuelNeeded(tc.Input)

			assert.Equal(t, tc.Expected, got)

		})
	}
}
