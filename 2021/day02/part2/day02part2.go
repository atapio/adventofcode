package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Submarine struct {
	pos   int
	depth int
	aim   int
}

func (s Submarine) Move(direction string, speed int) Submarine {
	switch direction {
	case "forward":
		return Submarine{pos: s.pos + speed, depth: s.depth + speed*s.aim, aim: s.aim}
	case "down":
		return Submarine{pos: s.pos, depth: s.depth, aim: s.aim + speed}
	case "up":
		return Submarine{pos: s.pos, depth: s.depth, aim: s.aim - speed}
	}

	log.Fatalf("invalid direction %s", direction)
	return Submarine{}
}

func findAnswer(lines []string) (int, error) {
	answer := 0

	sub := Submarine{}

	for n := range lines {
		parts := strings.Split(lines[n], " ")
		speed, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}
		sub = sub.Move(parts[0], speed)
		fmt.Printf("%s %v\n", lines[n], sub)
	}

	answer = sub.pos * sub.depth

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
