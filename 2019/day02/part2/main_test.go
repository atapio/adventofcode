package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompute(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name  string
		Input string

		Expected []int
	}{
		{
			Name:     "1",
			Input:    "1,0,0,0,99",
			Expected: []int{2, 0, 0, 0, 99},
		},
		{
			Name:     "2",
			Input:    "2,3,0,3,99",
			Expected: []int{2, 3, 0, 6, 99},
		},
		{
			Name:     "3",
			Input:    "2,4,4,5,99,0",
			Expected: []int{2, 4, 4, 5, 99, 9801},
		},
		{
			Name:     "4",
			Input:    "1,1,1,4,99,5,6,0,99",
			Expected: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			got, err := compute(tc.Input, false)
			require.NoError(t, err)

			assert.Equal(t, tc.Expected, got)
		})
	}
}
