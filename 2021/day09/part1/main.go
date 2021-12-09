package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func (g *Grid) LowPoints() []int {
	lows := []int{}

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
			lows = append(lows, height)
		}
	}

	return lows
}

func findAnswer(lines []string) (int, error) {
	answer := 0

	g := NewGrid()
	g.Init(lines)

	lows := g.LowPoints()

	for _, low := range lows {
		answer += low + 1
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
