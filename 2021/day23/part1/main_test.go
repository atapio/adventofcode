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
			Expected: 12521,
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

func TestFromRoomToHallway(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		Name   string
		Room   int
		Burrow *Burrow
		Goal   *Burrow

		Expected       *Burrow
		ExpectedLength int
	}{
		{
			Name: "simple",
			Room: 0,
			Burrow: &Burrow{
				cost:    0,
				hallway: []Amphipod{"", "", "", "", "", "", "", "", "", "", ""},
				rooms: [][]Amphipod{
					{"B", "A"},
					{"C", "B"},
					{"D", "C"},
					{"A", "D"},
				},
			},
			Goal: &Burrow{
				hallway: []Amphipod{"", "", "", "", "", "", "", "", "", "", ""},
				rooms: [][]Amphipod{
					{"A", "A"},
					{"B", "B"},
					{"C", "C"},
					{"D", "D"},
				},
			},
			Expected: &Burrow{
				cost:    30,
				hallway: []Amphipod{"B", "", "", "", "", "", "", "", "", "", ""},
				rooms: [][]Amphipod{
					{"", "A"},
					{"C", "B"},
					{"D", "C"},
					{"A", "D"},
				},
			},
			ExpectedLength: 7,
		},
		{
			Name: "with obstacle",
			Room: 3,
			Burrow: &Burrow{
				cost:    0,
				hallway: []Amphipod{"", "A", "", "", "", "", "", "", "", "", ""},
				rooms: [][]Amphipod{
					{"C", "A"},
					{"D", "B"},
					{"", "C"},
					{"B", "D"},
				},
			},
			Goal: &Burrow{
				hallway: []Amphipod{"", "", "", "", "", "", "", "", "", "", ""},
				rooms: [][]Amphipod{
					{"A", "A"},
					{"B", "B"},
					{"C", "C"},
					{"D", "D"},
				},
			},
			Expected: &Burrow{
				cost:    60,
				hallway: []Amphipod{"", "A", "", "B", "", "", "", "", "", "", ""},
				rooms: [][]Amphipod{
					{"C", "A"},
					{"D", "B"},
					{"", "C"},
					{"", "D"},
				},
			},
			ExpectedLength: 5,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			got := tc.Burrow.FromRoomToHallway(tc.Room, tc.Goal)

			assert.Len(t, got, tc.ExpectedLength)
			assert.Equal(t, tc.Expected, got[0])

		})
	}
}

func TestFromHallwayToRoom(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		Name   string
		Room   int
		Burrow *Burrow
		Goal   *Burrow

		Expected       *Burrow
		ExpectedLength int
	}{
		{
			Name: "simple",
			Room: 0,
			Burrow: &Burrow{
				cost:    0,
				hallway: []Amphipod{"A", "", "", "B", "", "", "", "", "", "", ""},
				rooms: [][]Amphipod{
					{"", "A"},
					{"C", "B"},
					{"D", "C"},
					{"", "D"},
				},
			},
			Goal: &Burrow{
				hallway: []Amphipod{"", "", "", "", "", "", "", "", "", "", ""},
				rooms: [][]Amphipod{
					{"A", "A"},
					{"B", "B"},
					{"C", "C"},
					{"D", "D"},
				},
			},
			Expected: &Burrow{
				cost:    3,
				hallway: []Amphipod{"", "", "", "B", "", "", "", "", "", "", ""},
				rooms: [][]Amphipod{
					{"A", "A"},
					{"C", "B"},
					{"D", "C"},
					{"", "D"},
				},
			},
			ExpectedLength: 1,
		},
		{
			Name: "with obstacle",
			Room: 1,
			Burrow: &Burrow{
				cost:    1,
				hallway: []Amphipod{"B", "A", "", "", "B", "", "", "", "", "", ""},
				rooms: [][]Amphipod{
					{"C", "A"},
					{"", ""},
					{"D", "C"},
					{"", "D"},
				},
			},
			Goal: &Burrow{
				hallway: []Amphipod{"", "", "", "", "", "", "", "", "", "", ""},
				rooms: [][]Amphipod{
					{"A", "A"},
					{"B", "B"},
					{"C", "C"},
					{"D", "D"},
				},
			},
			Expected: &Burrow{
				cost:    21,
				hallway: []Amphipod{"B", "A", "", "", "", "", "", "", "", "", ""},
				rooms: [][]Amphipod{
					{"C", "A"},
					{"", "B"},
					{"D", "C"},
					{"", "D"},
				},
			},
			ExpectedLength: 1,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			got := tc.Burrow.FromHallwayToRoom(tc.Room, tc.Goal)

			assert.Len(t, got, tc.ExpectedLength)
			assert.Equal(t, tc.Expected, got[0])

		})
	}
}

func TestOpenRoom(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		Name   string
		Room   int
		Burrow *Burrow
		Goal   *Burrow

		Expected bool
	}{
		{
			Name: "open",
			Room: 0,
			Burrow: &Burrow{
				cost:    0,
				hallway: []Amphipod{"B", "", "", "", "", "", "", "", "", "", ""},
				rooms: [][]Amphipod{
					{"", "A"},
					{"C", "B"},
					{"D", "C"},
					{"A", "D"},
				},
			},
			Goal: &Burrow{
				hallway: []Amphipod{"", "", "", "", "", "", "", "", "", "", ""},
				rooms: [][]Amphipod{
					{"A", "A"},
					{"B", "B"},
					{"C", "C"},
					{"D", "D"},
				},
			},
			Expected: true,
		},
		{
			Name: "closed",
			Room: 0,
			Burrow: &Burrow{
				cost:    0,
				hallway: []Amphipod{"A", "", "", "", "", "", "", "", "", "", ""},
				rooms: [][]Amphipod{
					{"", "B"},
					{"C", "B"},
					{"D", "C"},
					{"A", "D"},
				},
			},
			Goal: &Burrow{
				hallway: []Amphipod{"", "", "", "", "", "", "", "", "", "", ""},
				rooms: [][]Amphipod{
					{"A", "A"},
					{"B", "B"},
					{"C", "C"},
					{"D", "D"},
				},
			},
			Expected: false,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			got := tc.Burrow.RoomOpen(tc.Room, tc.Goal)

			assert.Equal(t, tc.Expected, got)

		})
	}
}
