package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var lineFormat = regexp.MustCompile(`^(on|off) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)$`)

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

type Point struct {
	x int
	y int
	z int
}

type Grid struct {
	cubes map[Point]bool
	min   Point
	max   Point
}

func NewGrid() *Grid {
	return &Grid{
		cubes: map[Point]bool{},
	}
}

func (g *Grid) Switch(on bool, p1 Point, p2 Point) {
	from := Point{
		x: min(p1.x, p2.x),
		y: min(p1.y, p2.y),
		z: min(p1.z, p2.z),
	}
	to := Point{
		x: max(p1.x, p2.x),
		y: max(p1.y, p2.y),
		z: max(p1.z, p2.z),
	}

	fmt.Printf("switch: %t %v - %v\n", on, from, to)

	if (from.x < -50 || from.y < -50 || from.z < -50) ||
		(to.x > 50 || to.y > 50 || to.z > 50) {
		fmt.Println("skip")
		return
	}

	for x := from.x; x <= to.x; x++ {
		for y := from.y; y <= to.y; y++ {
			for z := from.z; z <= to.z; z++ {
				switch on {
				case true:
					g.cubes[Point{x, y, z}] = true
				case false:
					delete(g.cubes, Point{x, y, z})
				}
			}
		}
	}
	if on {
		g.min.x = min(g.min.x, from.x)
		g.min.y = min(g.min.y, from.y)
		g.min.z = min(g.min.z, from.z)
		g.max.x = max(g.max.x, to.x)
		g.max.y = max(g.max.y, to.y)
		g.max.z = max(g.max.z, to.z)
	}
}

func (g Grid) CountActive() int {
	count := 0
	for range g.cubes {
		count++
	}

	return count
}

func findAnswer(lines []string) (int, error) {
	var err error

	answer := 0
	g := NewGrid()
	for _, line := range lines {
		fmt.Println(line)
		m := lineFormat.FindStringSubmatch(line)
		if len(m) != 8 {
			return 0, fmt.Errorf("failed to parse input '%s'", lines[0])
		}

		on := false
		if m[1] == "on" {
			on = true
		}

		p1 := Point{}
		p2 := Point{}
		p1.x, err = strconv.Atoi(m[2])
		if err != nil {
			return 0, err
		}
		p2.x, err = strconv.Atoi(m[3])
		if err != nil {
			return 0, err
		}
		p1.y, err = strconv.Atoi(m[4])
		if err != nil {
			return 0, err
		}
		p2.y, err = strconv.Atoi(m[5])
		if err != nil {
			return 0, err
		}
		p1.z, err = strconv.Atoi(m[6])
		if err != nil {
			return 0, err
		}
		p2.z, err = strconv.Atoi(m[7])
		if err != nil {
			return 0, err
		}
		g.Switch(on, p1, p2)
		fmt.Printf("grid: min: %v max: %v active: %d\n", g.min, g.max, g.CountActive())
	}

	answer = g.CountActive()

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
