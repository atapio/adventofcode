package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Map struct {
	width  int
	height int
	land   []string
}

func (m *Map) LandAt(x, y int) rune {
	x = x % m.width

	return []rune(m.land[y])[x]
}

func (m *Map) HasTree(x, y int) bool {
	switch m.LandAt(x, y) {
	case '#':
		return true
	}

	return false
}

type data struct {
	input *Map
}

func findAnswer(d data) int {
	trees := 0

	x := 0
	dX := 3
	dY := 1

	for y := 0; y < d.input.height; y = y + dY {
		if d.input.HasTree(x, y) {
			trees++
		}
		x = x + dX
	}

	return trees
}

func processInput(lines []string) data {
	d := data{
		input: &Map{
			width:  len(lines[0]),
			height: len(lines),
			land:   lines,
		},
	}

	return d
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

	d := processInput(input)

	answer := findAnswer(d)

	fmt.Printf("answer: %d\n", answer)
}
