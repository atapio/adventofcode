package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const MaxCycle = 8
const CycleStart = 6

type Swarm struct {
	positions []int
}

func NewSwarm() *Swarm {
	return &Swarm{
		positions: []int{},
	}
}

func (s *Swarm) MoveTo(position int) int {
	total := 0
	for _, p := range s.positions {
		total += abs(p - position)
	}

	return total
}

func (s *Swarm) Median() int {
	sort.Ints(s.positions)

	l := len(s.positions)
	median := s.positions[l/2]

	if l/2%2 == 0 {
		m := float64(s.positions[l/2]) + float64(s.positions[l/2-1])/2
		median = int(math.Round(m))
		fmt.Printf("median: %d, %f\n", median, m)
	}

	return median
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findAnswer(lines []string) (int, error) {
	swarm := NewSwarm()
	for _, day := range strings.Split(lines[0], ",") {
		d, err := strconv.Atoi(day)
		if err != nil {
			return 0, err
		}
		swarm.positions = append(swarm.positions, d)
	}

	sort.Ints(swarm.positions)
	max := swarm.positions[len(swarm.positions)-1]

	min := swarm.MoveTo(0)

	for i := 1; i < max; i++ {
		move := swarm.MoveTo(i)
		if move < min {
			min = move
			fmt.Printf("pos %d, move: %d\n", i, move)
		}
	}
	return min, nil
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
