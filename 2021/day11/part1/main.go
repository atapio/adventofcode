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
		{x - 1, y - 1},
		{x - 1, y},
		{x - 1, y + 1},
		{x, y - 1},
		{x, y + 1},
		{x + 1, y - 1},
		{x + 1, y},
		{x + 1, y + 1},
	}

	return n
}

type Octopus struct {
	energy  int
	flashed bool
}

type Grid struct {
	octopi map[Coordinate]*Octopus
}

func NewGrid() *Grid {
	return &Grid{
		octopi: map[Coordinate]*Octopus{},
	}
}

func (g *Grid) Init(lines []string) error {
	for y, row := range lines {
		for x, col := range strings.Split(row, "") {
			e, err := strconv.Atoi(col)
			if err != nil {
				return err
			}

			g.octopi[Coordinate{x, y}] = &Octopus{energy: e}
		}
	}

	return nil
}

func (g *Grid) Step() int {
	flashes := 0
	for _, o := range g.octopi {
		o.energy++
		o.flashed = false
	}

	for {
		f := g.Flash()
		flashes += f
		if f == 0 {
			break
		}
	}

	for _, o := range g.octopi {
		if o.energy > 9 {
			o.energy = 0
			o.flashed = false
		}
	}

	return flashes
}

func (g *Grid) Flash() int {
	flashes := 0

	for c, o := range g.octopi {
		if !o.flashed && o.energy > 9 {
			o.flashed = true
			flashes++

			for _, n := range c.Neighbors() {
				if neighbor, ok := g.octopi[n]; ok {
					neighbor.energy++
				}
			}
		}
	}

	return flashes
}

func findAnswer(lines []string) (int, error) {
	answer := 0

	g := NewGrid()
	g.Init(lines)

	for step := 1; step <= 100; step++ {
		answer += g.Step()
		fmt.Printf("step: %d flashes: %d\n", step, answer)
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
