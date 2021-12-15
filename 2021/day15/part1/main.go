package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func (c Point) Neighbors(diagonals bool) []Point {
	x := c.x
	y := c.y

	n := []Point{
		{x - 1, y},
		{x + 1, y},
		{x, y - 1},
		{x, y + 1},
	}
	if diagonals {
		n = append(n, Point{x - 1, y - 1})
		n = append(n, Point{x - 1, y + 1})
		n = append(n, Point{x + 1, y - 1})
		n = append(n, Point{x + 1, y + 1})
	}

	return n
}

type Grid struct {
	risks map[Point]int
	maxX  int
	maxY  int
}

func NewGrid() *Grid {
	return &Grid{
		risks: map[Point]int{},
	}
}

func (g *Grid) Neighbors(point Point, diagonals bool) []Point {
	neighbors := []Point{}

	for _, neigbor := range point.Neighbors(diagonals) {
		if _, ok := g.risks[neigbor]; ok {
			neighbors = append(neighbors, neigbor)
		}
	}

	return neighbors
}

func (g *Grid) Init(lines []string) error {
	for y, row := range lines {
		if y > g.maxY {
			g.maxY = y
		}
		for x, col := range strings.Split(row, "") {
			if x > g.maxX {
				g.maxX = x
			}
			level, err := strconv.Atoi(col)
			if err != nil {
				return err
			}

			g.risks[Point{x, y}] = level
		}
	}

	return nil
}

func (g *Grid) Distance(from Point, to Point) int {
	distances := map[Point]int{from: 0}
	visited := map[Point]bool{}
	queue := map[Point]int{from: 0}

	for len(queue) > 0 {
		curr, _ := Min(queue)
		delete(queue, curr)
		fmt.Printf("current: %v\n", curr)
		visited[curr] = true
		if curr == to {
			break
		}

		for _, n := range g.Neighbors(curr, false) {
			if _, ok := visited[n]; ok {
				continue
			}
			dist := g.risks[n]
			newCost := distances[curr] + dist
			oldCost, ok := distances[n]

			if !ok || newCost < oldCost {
				distances[n] = newCost
				queue[n] = newCost
			}
		}

	}

	return distances[to]
}

func Min(m map[Point]int) (Point, int) {
	point := Point{}
	min := -1
	for p, v := range m {
		if v < min || min == -1 {
			point = p
			min = v
		}
	}
	return point, m[point]
}

func findAnswer(lines []string) (int, error) {
	answer := 0

	g := NewGrid()
	g.Init(lines)

	answer = g.Distance(Point{0, 0}, Point{g.maxX, g.maxY})

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
