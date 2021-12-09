package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func (c Coordinate) Neighbors() []Coordinate {
	x := c.x
	y := c.y

	n := []Coordinate{
		{x - 1, y},
		{x + 1, y},
		{x, y - 1},
		{x, y + 1},
	}

	return n
}

type Grid struct {
	heights map[Coordinate]int
}

func NewGrid() *Grid {
	return &Grid{
		heights: map[Coordinate]int{},
	}
}

func (g *Grid) Init(lines []string) error {
	for y, row := range lines {
		for x, col := range strings.Split(row, "") {
			h, err := strconv.Atoi(col)
			if err != nil {
				return err
			}

			g.heights[Coordinate{x, y}] = h
		}
	}

	return nil
}

func (g *Grid) LowPoints() []Coordinate {
	lows := []Coordinate{}

	for coord := range g.heights {
		neigbors := coord.Neighbors()
		height := g.heights[coord]
		isLow := true

		for _, n := range neigbors {
			if val, ok := g.heights[n]; ok {
				if val <= height {
					isLow = false
				}
				//fmt.Printf("c: %v %d, neighbor: %v %d isLow: %t\n", coord, height, n, val, isLow)
			}
		}
		if isLow {
			lows = append(lows, coord)
		}
	}

	return lows
}

func (g *Grid) Basins() []int {
	pointsLeft := map[Coordinate]int{}
	for k, v := range g.heights {
		// Locations of height 9 do not count as being in any basin
		if v < 9 {
			pointsLeft[k] = v
		}
	}

	lows := g.LowPoints()
	basins := make([]int, len(lows))

	for i, low := range lows {
		size := 1
		queue := []Coordinate{low}
		delete(pointsLeft, low)

		for len(queue) > 0 {
			fmt.Printf("queue: %v\n", queue)
			var p Coordinate
			p, queue = queue[0], queue[1:]

			for _, n := range p.Neighbors() {
				if _, ok := pointsLeft[n]; ok {
					size++
					queue = append(queue, n)
					// all other locations will always be part of exactly one basin.
					delete(pointsLeft, n)
				}
			}
		}
		basins[i] = size
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	return basins
}

func findAnswer(lines []string) (int, error) {
	answer := 1

	g := NewGrid()
	g.Init(lines)

	basins := g.Basins()

	for _, size := range basins[0:3] {
		answer *= size
	}

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
