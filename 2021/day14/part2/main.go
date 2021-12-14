package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Polymer struct {
	rules    map[string]string
	pairs    map[string]int
	elements map[string]int
}

func NewPolymer() *Polymer {
	return &Polymer{
		rules:    map[string]string{},
		pairs:    map[string]int{},
		elements: map[string]int{},
	}
}

func (p *Polymer) Polymerize(polymer string) {
	elements := strings.Split(polymer, "")
	p.elements[elements[0]]++
	for i := 1; i < len(elements); i++ {
		pair := elements[i-1] + elements[i]
		p.elements[elements[i]]++
		p.pairs[pair]++
	}
}

func (p *Polymer) Step() {
	pairs := []string{}
	for pair := range p.pairs {
		pairs = append(pairs, pair)
	}

	nextPairs := map[string]int{}
	p.elements = map[string]int{}

	for pair, count := range p.pairs {
		elements := strings.Split(pair, "")
		if val, ok := p.rules[pair]; ok {
			//p.elements[elements[0]] += count
			p.elements[val] += count
			p.elements[elements[1]] += count
			nextPairs[elements[0]+val] += count
			nextPairs[val+elements[1]] += count
		} else {
			panic(fmt.Errorf("pair not found %s\n", pair))
		}
	}
	p.pairs = nextPairs
}

func findAnswer(lines []string) (int, error) {
	answer := 0

	template, ruleLines := lines[0], lines[1:]
	p := NewPolymer()

	for _, rule := range ruleLines {
		if rule == "" {
			continue
		}
		ruleParts := strings.Split(rule, " -> ")
		p.rules[ruleParts[0]] = ruleParts[1]
	}

	polymer := template

	p.Polymerize(polymer)
	fmt.Printf("pairs: %v\n", p.pairs)
	fmt.Printf("elements: %v\n", p.elements)
	for i := 1; i <= 40; i++ {
		p.Step()
		fmt.Printf("step: %d pairs: %v\n", i, p.pairs)
		fmt.Printf("step: %d elements: %v\n", i, p.elements)
	}

	freqs := []int{}
	for _, c := range p.elements {
		freqs = append(freqs, c)
	}

	sort.Ints(freqs)
	min := freqs[0]
	max := freqs[len(freqs)-1]

	answer = max - min

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
