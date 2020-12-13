package main

import (
	"reflect"
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
			Name: "example1.1",
			Input: []string{
				"939",
				"7,13",
			},
			Expected: 77,
		},
		{
			Name: "example1.2",
			Input: []string{
				"939",
				"7,13,x,x,59",
			},
			Expected: 350,
		},
		{
			Name: "example1",
			Input: []string{
				"939",
				"7,13,x,x,59,x,31,19",
			},
			Expected: 1068781,
		},
		{
			Name: "example2",
			Input: []string{
				"939",
				"17,x,13,19",
			},
			Expected: 3417,
		},
		/*
			{
				Name: "example00",
				Input: []string{
					"939",
					"3,4",
				},
				Expected: 9,
			},
				{
					Name: "example3",
					Input: []string{
						"939",
						"67,7,59,61",
					},
					Expected: 754018,
				},
		*/
		/*
			{
				Name: "example4",
				Input: []string{
					"939",
					"67,x,7,59,61",
				},
				Expected: 779210,
			},
			{
				Name: "example5",
				Input: []string{
					"939",
					"67,7,x,59,61",
				},
				Expected: 1261476,
			},
			{
				Name: "example6",
				Input: []string{
					"939",
					"1789,37,47,1889",
				},
				Expected: 1202161486,
			},
		*/
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			//t.Parallel()

			got, err := findAnswer(tc.Input)
			require.NoError(t, err)

			assert.Equal(t, tc.Expected, got)

		})
	}
}

func TestExtendedEuclid(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want EEResult
	}{
		{
			name: "simple",
			args: args{m: 240, n: 46},
			want: EEResult{d: 2, a: -9, b: 47},
		},
		{
			name: "simple2",
			args: args{m: 3, n: 4},
			want: EEResult{d: 1, a: -1, b: 1},
		},
		{
			name: "simple3",
			args: args{m: 5, n: 3 * 4},
			want: EEResult{d: 1, a: 5, b: -2},
		},
		{
			name: "simple4",
			args: args{m: 1, n: 2},
			want: EEResult{d: 1, a: 1, b: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtendedEuclid(tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtendedEuclid() = %v, want %v", got, tt.want)
			}
		})
	}
}
