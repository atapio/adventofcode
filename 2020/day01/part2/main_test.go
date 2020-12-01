package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAnswer(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name  string
		Input []int

		Expected int
	}{
		{
			Name:     "example",
			Input:    []int{1721, 979, 366, 299, 675, 1456},
			Expected: 241861950,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			input := data{input: tc.Input}

			got := findAnswer(input)

			assert.Equal(t, tc.Expected, got)

		})
	}
}
