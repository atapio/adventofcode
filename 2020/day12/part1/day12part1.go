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
	position  Point
	direction int
}

func NewShip() *Ship {
	return &Ship{direction: 90}
}

func (s *Ship) Turn(angle int) {
	s.direction = (s.direction + angle + 360) % 360
}

func (s *Ship) Move(direction string, distance int) {
	switch direction {
	case "N":
		s.position.y += distance
	case "S":
		s.position.y -= distance
	case "E":
		s.position.x += distance
	case "W":
		s.position.x -= distance
	case "L":
		s.Turn(-distance)
	case "R":
		s.Turn(distance)
	case "F":
		switch s.direction {
		case 0:
			s.Move("N", distance)
		case 90:
			s.Move("E", distance)
		case 180:
			s.Move("S", distance)
		case 270:
			s.Move("W", distance)
		}
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
		fmt.Printf("move: %s pos: %d %d dir: %d\n", line, s.position.x, s.position.y, s.direction)
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
