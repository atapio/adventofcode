package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Input int

		Expected bool
	}{
		{
			Input:    111111,
			Expected: true,
		},
		{
			Input:    223450,
			Expected: false,
		},
		{
			Input:    123789,
			Expected: false,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(fmt.Sprintf("case %d", tc.Input), func(t *testing.T) {
			t.Parallel()

			got := isValid(tc.Input)

			assert.Equal(t, tc.Expected, got)

		})
	}
}
