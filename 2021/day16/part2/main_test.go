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
			Input:    []string{"C200B40A82"},
			Expected: 3,
		},
		{
			Name:     "example2",
			Input:    []string{"04005AC33890"},
			Expected: 54,
		},
		{
			Name:     "example3",
			Input:    []string{"880086C3E88112"},
			Expected: 7,
		},
		{
			Name:     "example4",
			Input:    []string{"CE00C43D881120"},
			Expected: 9,
		},
		{
			Name:     "example5",
			Input:    []string{"D8005AC2A8F0"},
			Expected: 1,
		},
		{
			Name:     "example6",
			Input:    []string{"F600BC2D8F"},
			Expected: 0,
		},
		{
			Name:     "example7",
			Input:    []string{"9C005AC2F8F0"},
			Expected: 0,
		},
		{
			Name:     "example8",
			Input:    []string{"9C0141080250320F1802104A08"},
			Expected: 1,
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
