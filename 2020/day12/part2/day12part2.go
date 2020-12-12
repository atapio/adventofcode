package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Direction struct {
	dir int
}

func (d *Direction) Turn(angle int) {
	d.dir = (d.dir + angle + 360) % 360
}

type Point struct {
	x int
	y int
}

func (p Point) distance(p2 Point) int {
	return Abs(p.x-p2.x) + Abs(p.y-p2.y)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Ship struct {
	waypoint Point
	position Point
}

func NewShip() *Ship {
	return &Ship{
		waypoint: Point{x: 10, y: 1},
	}
}

func Sign(a int) int {
	if a < 0 {
		return -1
	}
	if a > 0 {
		return 1
	}
	return 0
}

func (s *Ship) Turn(angle int) {
	direction := (angle + 360) % 360

	orig := Point{x: s.waypoint.x, y: s.waypoint.y}

	switch direction {
	case 0:
		// do nothing
	case 90:
		s.waypoint.y = -orig.x
		s.waypoint.x = orig.y
	case 180:
		s.waypoint.y = -orig.y
		s.waypoint.x = -orig.x
	case 270:
		s.waypoint.y = orig.x
		s.waypoint.x = -orig.y
	}
}

func (s *Ship) Move(direction string, distance int) {
	switch direction {
	case "N":
		s.waypoint.y += distance
	case "S":
		s.waypoint.y -= distance
	case "E":
		s.waypoint.x += distance
	case "W":
		s.waypoint.x -= distance
	case "L":
		s.Turn(-distance)
	case "R":
		s.Turn(distance)
	case "F":
		s.position.x += s.waypoint.x * distance
		s.position.y += s.waypoint.y * distance
	}
}

func findAnswer(lines []string) (int, error) {
	s := NewShip()

	for _, line := range lines {
		cmd := line[0:1]
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			return 0, err
		}
		s.Move(cmd, value)
		fmt.Printf("move: %s pos: %d %d wp: %d %d\n", line, s.position.x, s.position.y, s.waypoint.x, s.waypoint.y)
	}
	return Point{}.distance(s.position), nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		input = append(input, l)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	answer, err := findAnswer(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("answer: %d\n", answer)
}
