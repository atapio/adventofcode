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

var foldFormat = regexp.MustCompile(`^fold along ([xy])=(\d+)$`)

type Coordinate struct {
	x int
	y int
}

type Paper struct {
	dots map[Coordinate]int
}

func NewPaper() *Paper {
	return &Paper{
		dots: map[Coordinate]int{},
	}
}

func (p *Paper) Init(lines []string) error {
	for _, line := range lines {
		coords := strings.Split(line, ",")

		x, err := strconv.Atoi(coords[0])
		if err != nil {
			return err
		}

		y, err := strconv.Atoi(coords[1])
		if err != nil {
			return err
		}
		p.dots[Coordinate{x, y}]++
	}

	return nil
}

func (p *Paper) Dots() int {
	count := 0
	for _, c := range p.dots {
		if c > 0 {
			count++
		}
	}
	return count
}

func (p *Paper) Fold(line Coordinate) {
	for c := range p.dots {
		if line.x != 0 {
			if c.x > line.x {
				p.dots[Coordinate{2*line.x - c.x, c.y}]++
				delete(p.dots, c)
			}
		}

		if line.y != 0 {
			if c.y > line.y {
				p.dots[Coordinate{c.x, 2*line.y - c.y}]++
				delete(p.dots, c)
			}
		}
	}
}

func findAnswer(lines []string) (int, error) {
	answer := 0

	folds := []Coordinate{}
	coords := []string{}

	for _, line := range lines {
		m := foldFormat.FindStringSubmatch(line)
		if m != nil {
			c := Coordinate{}
			fold, err := strconv.Atoi(m[2])
			if err != nil {
				return 0, err
			}

			if m[1] == "x" {
				c.x = fold
			}
			if m[1] == "y" {
				c.y = fold
			}

			folds = append(folds, c)
			continue
		}

		coords = append(coords, line)
	}
	p := NewPaper()
	p.Init(coords)

	p.Fold(folds[0])
	answer = p.Dots()

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
