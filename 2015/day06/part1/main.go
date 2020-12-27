package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var commandFormat = regexp.MustCompile(`^(turn on|toggle|turn off) (\d+,\d+) through (\d+,\d+)$`)

type Coordinate struct {
	x int
	y int
}

func (c Coordinate) PointsIn(other Coordinate) []Coordinate {
	points := []Coordinate{}
	for x := c.x; x <= other.x; x++ {
		for y := c.y; y <= other.y; y++ {
			points = append(points, Coordinate{x, y})
		}
	}
	return points
}

func ParseCoordinate(input string) (Coordinate, error) {
	xy := strings.Split(input, ",")

	x, err := strconv.Atoi(xy[0])
	if err != nil {
		return Coordinate{}, err
	}
	y, err := strconv.Atoi(xy[1])
	if err != nil {
		return Coordinate{}, err
	}

	return Coordinate{x, y}, nil

}

type Grid struct {
	lights map[Coordinate]bool
}

func NewGrid() *Grid {
	return &Grid{lights: map[Coordinate]bool{}}
}

func (g *Grid) TurnOn(c1 Coordinate, c2 Coordinate) {
	for _, l := range c1.PointsIn(c2) {
		g.lights[l] = true
	}
}

func (g *Grid) TurnOff(c1 Coordinate, c2 Coordinate) {
	for _, l := range c1.PointsIn(c2) {
		delete(g.lights, l)
	}
}

func (g *Grid) Toggle(c1 Coordinate, c2 Coordinate) {
	for _, l := range c1.PointsIn(c2) {
		_, on := g.lights[l]
		if on {
			delete(g.lights, l)
		} else {
			g.lights[l] = true
		}
	}
}

func (g Grid) LightsOn() int {
	count := 0
	for _, on := range g.lights {
		if on {
			count++
		}
	}
	return count
}

func findAnswer(lines []string) (int, error) {
	answer := 0
	grid := NewGrid()

	for _, line := range lines {
		m := commandFormat.FindStringSubmatch(line)

		cmd, from, to := m[1], m[2], m[3]

		c1, err := ParseCoordinate(from)
		if err != nil {
			return 0, err
		}
		c2, err := ParseCoordinate(to)
		if err != nil {
			return 0, err
		}

		switch cmd {
		case "turn on":
			grid.TurnOn(c1, c2)
		case "turn off":
			grid.TurnOff(c1, c2)
		case "toggle":
			grid.Toggle(c1, c2)
		}
	}

	answer = grid.LightsOn()

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
	input, err := parseFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	answer, err := findAnswer(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("answer: %d\n", answer)
}
