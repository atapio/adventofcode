package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var tileFormat = regexp.MustCompile(`(se|e|sw|nw|ne|w)`)

type Coordinate struct {
	x int
	y int
}

func (c Coordinate) Move(direction string) Coordinate {
	switch direction {
	case "ne":
		return Coordinate{c.x + 1, c.y + 1}
	case "e":
		return Coordinate{c.x + 2, c.y}
	case "se":
		return Coordinate{c.x + 1, c.y - 1}
	case "sw":
		return Coordinate{c.x - 1, c.y - 1}
	case "w":
		return Coordinate{c.x - 2, c.y}
	case "nw":
		return Coordinate{c.x - 1, c.y + 1}
	}

	log.Fatalf("invalid direction %s", direction)
	return Coordinate{}
}

type Floor struct {
	tiles map[Coordinate]bool
}

func NewFloor() *Floor {
	f := &Floor{
		tiles: map[Coordinate]bool{},
	}

	return f
}

func (f *Floor) FlipTile(c Coordinate) {
	log.Printf("flip %v %t", c, f.tiles[c])
	f.tiles[c] = !f.tiles[c]
}

func (f *Floor) CountBlack() int {
	count := 0
	for _, black := range f.tiles {
		if black {
			count++
		}
	}

	return count
}

func findAnswer(lines []string) (int, error) {
	answer := 0

	floor := NewFloor()

	for _, line := range lines {
		m := tileFormat.FindAllStringSubmatch(line, -1)

		c := Coordinate{}

		log.Printf("line %s", line)
		for _, match := range m {
			c = c.Move(match[0])
			log.Printf("match %s coord %v", match[0], c)
		}

		/*
			for i := 1; i < len(m); i++ {
				c.Move(m[i])
				log.Printf("coordinate %v", c)
			}
		*/

		floor.FlipTile(c)
		log.Printf("black: %d", floor.CountBlack())
	}
	log.Printf("floor: %v", floor.tiles)

	answer = floor.CountBlack()

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
