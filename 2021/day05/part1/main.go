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

var commandFormat = regexp.MustCompile(`^(\d+,\d+) -> (\d+,\d+)$`)

type Coordinate struct {
	x int
	y int
}

func (c Coordinate) PointsIn(other Coordinate) []Coordinate {
	points := []Coordinate{}

	minX := min(c.x, other.x)
	maxX := max(c.x, other.x)
	minY := min(c.y, other.y)
	maxY := max(c.y, other.y)

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			points = append(points, Coordinate{x, y})
		}
	}
	fmt.Printf("points %v->%v: %v\n", c, other, points)
	return points
}

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
	lines map[Coordinate]int
}

func NewGrid() *Grid {
	return &Grid{lines: map[Coordinate]int{}}
}

func (g *Grid) DrawLine(c1 Coordinate, c2 Coordinate) {
	for _, l := range c1.PointsIn(c2) {
		g.lines[l]++
	}
}

func (g Grid) Overlaps() int {
	count := 0
	for _, b := range g.lines {
		fmt.Printf("line: %v\n", b)
		if b > 1 {
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

		from, to := m[1], m[2]

		c1, err := ParseCoordinate(from)
		if err != nil {
			return 0, err
		}
		c2, err := ParseCoordinate(to)
		if err != nil {
			return 0, err
		}

		if c1.x == c2.x || c1.y == c2.y {
			grid.DrawLine(c1, c2)
		}
	}

	answer = grid.Overlaps()

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
