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

		Expected int
	}{
		{
			Name:     "example1",
			Input:    []string{"D2FE28"},
			Expected: 6,
		},
		{
			Name:     "example2",
			Input:    []string{"38006F45291200"},
			Expected: 9,
		},
		{
			Name:     "example3",
			Input:    []string{"EE00D40C823060"},
			Expected: 14,
		},
		{
			Name:     "example4",
			Input:    []string{"8A004A801A8002F478"},
			Expected: 16,
		},
		{
			Name:     "example5",
			Input:    []string{"620080001611562C8802118E34"},
			Expected: 12,
		},
		{
			Name:     "example6",
			Input:    []string{"C0015000016115A2E0802F182340"},
			Expected: 23,
		},
		{
			Name:     "example7",
			Input:    []string{"A0016C880162017C3686B18A3D4780"},
			Expected: 31,
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
