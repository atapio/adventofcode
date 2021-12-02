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
	z int
}

func (c Coordinate) Move(direction string, speed int) Coordinate {
	switch direction {
	case "forward":
		return Coordinate{c.x + speed, c.z}
	case "down":
		return Coordinate{c.x, c.z + speed}
	case "up":
		return Coordinate{c.x, c.z - speed}
	}

	log.Fatalf("invalid direction %s", direction)
	return Coordinate{}
}

func findAnswer(lines []string) (int, error) {
	answer := 0

	pos := Coordinate{}

	for n := range lines {
		parts := strings.Split(lines[n], " ")
		speed, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}
		pos = pos.Move(parts[0], speed)
	}

	answer = pos.x * pos.z

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
