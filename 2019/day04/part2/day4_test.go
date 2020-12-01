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
			Input:    112233,
			Expected: true,
		},
		{
			Input:    123444,
			Expected: false,
		},
		{
			Input:    111122,
			Expected: true,
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
