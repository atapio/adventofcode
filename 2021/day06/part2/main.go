package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const MaxCycle = 8
const CycleStart = 6

type Population struct {
	cycle []int
	day   int
}

func NewPopulation() *Population {
	return &Population{
		cycle: make([]int, MaxCycle+1),
	}
}

func (p *Population) NextDay() {
	p.day++

	today, rest := p.cycle[0], p.cycle[1:]

	p.cycle = rest
	for len(p.cycle) < MaxCycle+1 {
		p.cycle = append(p.cycle, 0)
	}

	p.cycle[CycleStart] += today
	p.cycle[MaxCycle] += today
}

func (p *Population) Count() int {
	c := 0
	for _, dc := range p.cycle {
		c += dc
	}
	return c
}

func findAnswer(lines []string) (int, error) {
	population := NewPopulation()

	for _, day := range strings.Split(lines[0], ",") {
		d, err := strconv.Atoi(day)
		if err != nil {
			return 0, err
		}
		population.cycle[d]++
	}

	for i := 0; i < 256; i++ {
		population.NextDay()
	}

	return population.Count(), nil
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
