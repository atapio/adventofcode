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
	dots  map[Coordinate]int
	sizeX int
	sizeY int
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
	if line.x != 0 {
		p.sizeX = line.x
	}
	if line.y != 0 {
		p.sizeY = line.y
	}
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

func (p *Paper) Draw() {
	for y := 0; y < p.sizeY; y++ {
		for x := 0; x < p.sizeX; x++ {
			if _, ok := p.dots[Coordinate{x, y}]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
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

	for _, fold := range folds {
		p.Fold(fold)
	}

	p.Draw()

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
