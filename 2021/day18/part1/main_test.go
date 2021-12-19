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
		/*
				{
					Name:     "simple1",
					Input:    []string{"[9,1]"},
					Expected: 29,
				},
				{
					Name:     "simple2",
					Input:    []string{"[1,9]"},
					Expected: 21,
				},
				{
					Name:     "simple2.1",
					Input:    []string{"[[9,1],[1,9]]"},
					Expected: 129,
				},
				{
					Name:     "simple3",
					Input:    []string{"[[1,2],[[3,4],5]]"},
					Expected: 143,
				},
				{
					Name:     "simple4",
					Input:    []string{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
					Expected: 1384,
				},
				{
					Name:     "explode1",
					Input:    []string{"[[[[[9,8],1],2],3],4]"},
					Expected: 548,
				},
				{
					Name:     "explode2",
					Input:    []string{"[7,[6,[5,[4,[3,2]]]]]"},
					Expected: 285,
				},
			{
				Name: "sum1",
				Input: []string{
					"[[[[4,3],4],4],[7,[[8,4],9]]]",
					"[1,1]",
				},
				Expected: 1384,
			},
					{
						Name: "sum2",
						Input: []string{
							"[1,1]",
							"[2,2]",
							"[3,3]",
							"[4,4]",
						},
						Expected: 445,
					},
				{
					Name: "sum3",
					Input: []string{
						"[1,1]",
						"[2,2]",
						"[3,3]",
						"[4,4]",
						"[5,5]",
					},
					Expected: 791,
				},
					{
						Name: "sum4",
						Input: []string{
							"[1,1]",
							"[2,2]",
							"[3,3]",
							"[4,4]",
							"[5,5]",
							"[6,6]",
						},
						Expected: 1137,
					},
		*/
		{
			Name:     "example1",
			Expected: 4140,
		},
		/*
			{
				Name:     "example2",
				Expected: 4140,
			},
		*/
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
