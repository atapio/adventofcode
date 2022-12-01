package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

const rooms = 4
const roomSize = 2
const hallwaySize = 11

var roomPosition = []int{2, 4, 6, 8}

var roomFormat = regexp.MustCompile(`^[ #]{3}([A-D])[ #]([A-D])[ #]([A-D])[ #]([A-D])[ #]+$`)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Amphipod string

func (a Amphipod) Cost() int {
	switch a {
	case "A":
		return 1
	case "B":
		return 10
	case "C":
		return 100
	case "D":
		return 1000
	}

	return 100000
}

type Burrow struct {
	hallway []Amphipod
	rooms   [][]Amphipod
	cost    int
}

func NewBurrow() *Burrow {
	b := &Burrow{
		hallway: make([]Amphipod, hallwaySize),
		rooms:   make([][]Amphipod, rooms),
	}
	for i := range b.rooms {
		b.rooms[i] = make([]Amphipod, roomSize)
	}
	return b
}

func (b *Burrow) Copy() *Burrow {
	newBurrow := NewBurrow()

	newBurrow.cost = b.cost

	copy(newBurrow.hallway, b.hallway)
	for i := range b.rooms {
		copy(newBurrow.rooms[i], b.rooms[i])
	}
	return newBurrow
}

func (b *Burrow) Equals(other *Burrow) bool {
	// skip hallway, skip size checks
	for room := range b.rooms {
		for pos := range b.rooms[room] {
			if b.rooms[room][pos] != other.rooms[room][pos] {
				return false
			}
		}
	}
	return true
}

func (b *Burrow) String() string {
	burrow := "#############\n"

	burrow += "#"
	for _, a := range b.hallway {
		label := string(a)
		if label == "" {
			label = "."
		}
		burrow += label
	}
	burrow += "#\n"

	for pos := 0; pos < len(b.rooms[0]); pos++ {
		switch pos {
		case 0:
			burrow += "###"
		case 1:
			burrow += "  #"
		}
		for _, room := range b.rooms {
			label := string(room[pos])
			if label == "" {
				label = "."
			}
			burrow += label + "#"
		}
		if pos == 0 {
			burrow += "##"
		}
		burrow += "\n"
	}
	burrow += "  #########"

	return burrow
}

// Solve solves the puzzle using Djikstra
func (b *Burrow) Solve(end *Burrow) int {
	distances := map[string]int{b.String(): 0}
	visited := map[string]bool{}
	queue := map[string]int{b.String(): 0}

	for len(queue) > 0 {
		curr, _ := Min(queue)
		delete(queue, curr)

		visited[curr] = true
		if curr == end.String() {
			break
		}

		for i := range b.rooms {
			for _, next := range b.FromRoomToHallway(i, end) {
				if _, ok := visited[next.String()]; ok {
					continue
				}

				dist := g.risks[n]
				newCost := distances[curr] + dist
				oldCost, ok := distances[n]
			}

		}
	}
}

func Min(m map[string]int) (string, int) {
	s := ""
	min := -1
	for k, v := range m {
		if v < min || min == -1 {
			s = k
			min = v
		}
	}
	return s, m[s]
}

func (b *Burrow) solve(end *Burrow, path []*Burrow, depth int) (int, bool) {
	//fmt.Printf("Solve %d: cost: %d\n%s\npath: %d\n", depth, b.cost, b, len(path))
	fmt.Printf("\nSolve %d: cost: %d\n%s\n", depth, b.cost, b)
	depth++

	/*
		if depth > 8 {
			fmt.Printf("reached max depth\n")
			return b.cost, true
		}
	*/

	if b.Equals(end) {
		return b.cost, true
	}

	path = append(path, b)
	nextPaths := []*Burrow{}

	// try moving amphipod from room to hallway
	for i := range b.rooms {
		paths := b.FromRoomToHallway(i, end)

		for _, next := range paths {
			visited := false

			for _, p := range path {
				if p.Equals(next) {
					visited = true
					break
				}
			}

			if visited {
				continue
			}

			nextPaths = append(nextPaths, paths...)
		}
	}

	for i := range b.OpenRooms(end) {
		paths := b.FromHallwayToRoom(i, end)

		for _, next := range paths {
			visited := false

			for _, p := range path {
				if p.Equals(next) {
					visited = true
					break
				}
			}

			if visited {
				continue
			}

			nextPaths = append(nextPaths, paths...)
		}
	}

	var best *Burrow
	for _, next := range nextPaths {
		_, ok := next.solve(end, path, depth)
		if ok {
			if best == nil {
				best = next
			}
			if next.cost < best.cost {
				best = next
			}
		}
	}

	if best != nil {
		return best.cost, true
	}

	return -1, false
}

func (b *Burrow) FromRoomToHallway(from int, goal *Burrow) []*Burrow {
	paths := []*Burrow{}

	amphipod := Amphipod("")
	stepsInRoom := 0
	for i, a := range b.rooms[from] {
		// pick first non empty, others cannot move
		if a != "" {
			amphipod = a
			stepsInRoom = i
			break
		}
	}

	if amphipod == "" {
		return paths
	}

	// no need to move anything out from an solved room
	if b.RoomOpen(from, goal) {
		return paths
	}

	//fmt.Printf("possible amphipod %s room: %d/%d\n", amphipod, from, stepsInRoom)

OUTER:
	for p, a := range b.hallway {
		// skip non empty hallway positions
		if a != "" {
			continue
		}

		// skip room entrances
		for _, roomPos := range roomPosition {
			if p == roomPos {
				continue OUTER
			}
		}

		// check if route is empty
		for i := min(p, roomPosition[from]); i < max(p, roomPosition[from]); i++ {
			if b.hallway[i] != "" {
				continue OUTER
			}
		}
		steps := max(p, roomPosition[from]) - min(p, roomPosition[from]) + stepsInRoom + 1
		cost := steps * amphipod.Cost()
		path := b.Copy()
		path.cost += cost
		path.rooms[from][stepsInRoom] = ""
		path.hallway[p] = amphipod

		//fmt.Println(path.String())

		paths = append(paths, path)
	}

	return paths
}

func (b *Burrow) FromHallwayToRoom(to int, goal *Burrow) []*Burrow {
	paths := []*Burrow{}
	// assume room is open
	// find suitable amphipod

	amphipod := Amphipod("")
	stepsInRoom := 0

	for i := len(b.rooms[to]) - 1; i >= 0; i-- {
		if b.rooms[to][i] == goal.rooms[to][i] {
			continue
		}
		amphipod = goal.rooms[to][i]
		stepsInRoom = i
		break
	}

	//fmt.Printf("amphi: %s/%d\n", amphipod, stepsInRoom)

OUTER:
	for pos, a := range b.hallway {

		if a != amphipod {
			continue
		}

		// check if route is empty
		for i := min(pos, roomPosition[to]); i < max(pos, roomPosition[to]); i++ {
			//fmt.Printf("route: %d/%s\n", i, b.hallway[i])
			if pos == i {
				continue
			}
			if b.hallway[i] != "" {
				continue OUTER
			}
		}

		//fmt.Printf("pos: %s/%d\n", a, pos)

		steps := max(pos, roomPosition[to]) - min(pos, roomPosition[to]) + stepsInRoom + 1
		cost := steps * amphipod.Cost()
		path := b.Copy()
		path.cost += cost
		path.rooms[to][stepsInRoom] = amphipod
		path.hallway[pos] = ""

		//fmt.Println(path.String())

		paths = append(paths, path)
	}

	return paths

}

func (b *Burrow) OpenRooms(goal *Burrow) []int {
	open := []int{}
	for i := 0; i < roomSize; i++ {
		if b.RoomOpen(i, goal) {
			open = append(open, i)
		}
	}
	return open
}

func (b *Burrow) RoomOpen(room int, goal *Burrow) bool {
	// check if room is solved from bottom up
	solved := true
	onlyEmpty := false
	for i := len(b.rooms[room]) - 1; i >= 0; i-- {
		if b.rooms[room][i] == goal.rooms[room][i] && !onlyEmpty {
			continue
		}
		if b.rooms[room][i] == "" {
			onlyEmpty = true
			continue
		}
		solved = false
		break
	}
	return solved
}

func findAnswer(lines []string) (int, error) {
	burrow := NewBurrow()

	if lines[0] != "#############" {
		return 0, fmt.Errorf("invalid first row")
	}
	if lines[1] != "#...........#" {
		return 0, fmt.Errorf("invalid second row")
	}
	if lines[4] != "  #########" {
		return 0, fmt.Errorf("invalid last row")
	}

	roomRow1 := roomFormat.FindStringSubmatch(lines[2])
	if len(roomRow1) != 5 {
		return 0, fmt.Errorf("invalid first room row")
	}
	for i := 1; i < 5; i++ {
		burrow.rooms[i-1][0] = Amphipod(roomRow1[i])
	}

	roomRow2 := roomFormat.FindStringSubmatch(lines[3])
	if len(roomRow2) != 5 {
		return 0, fmt.Errorf("invalid second room row")
	}

	for i := 1; i < 5; i++ {
		burrow.rooms[i-1][1] = Amphipod(roomRow2[i])
	}

	fmt.Println(burrow)

	goal := NewBurrow()
	goal.rooms[0] = []Amphipod{"A", "A"}
	goal.rooms[1] = []Amphipod{"B", "B"}
	goal.rooms[2] = []Amphipod{"C", "C"}
	goal.rooms[3] = []Amphipod{"D", "D"}

	fmt.Println(goal)

	answer := burrow.Solve(goal)

	return answer, nil
}

func parseFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	input := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		input = append(input, l)
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}
	return input, nil
}

func main() {
	input, err := parseFile("input")
	if err != nil {
		log.Fatal(err)
	}

	answer, err := findAnswer(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("answer: %d\n", answer)
}
