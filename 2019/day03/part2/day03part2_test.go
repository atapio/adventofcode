package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAnswer(t *testing.T) {
	//t.Parallel()

	testCases := []struct {
		Name  string
		Input []string

		Expected int
	}{
		{
			Name:     "1",
			Input:    []string{"R8,U5,L5,D3", "U7,R6,D4,L4"},
			Expected: 30,
		},
		{
			Name:     "2",
			Input:    []string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"},
			Expected: 610,
		},
		{
			Name:     "3",
			Input:    []string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"},
			Expected: 410,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			//t.Parallel()

			got := FindAnswer(tc.Input)

			assert.Equal(t, tc.Expected, got)
		})
	}
}
