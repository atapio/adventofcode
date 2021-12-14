package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func findAnswer(lines []string) (int, error) {
	answer := 0

	template, ruleLines := lines[0], lines[1:]
	rules := map[string]string{}

	for _, rule := range ruleLines {
		if rule == "" {
			continue
		}
		ruleParts := strings.Split(rule, " -> ")
		rules[ruleParts[0]] = ruleParts[1]
	}

	polymer := template

	for i := 0; i < 10; i++ {
		elements := strings.Split(polymer, "")
		nextPolymer := []string{elements[0]}

		for j := 1; j < len(elements); j++ {
			prev := nextPolymer[len(nextPolymer)-1]
			curr := elements[j]
			nextPolymer = append(nextPolymer, rules[prev+curr])
			nextPolymer = append(nextPolymer, curr)
		}

		polymer = strings.Join(nextPolymer, "")
		fmt.Printf("step %d: %s\n", i, polymer)
	}

	counts := map[string]int{}
	for _, c := range strings.Split(polymer, "") {
		counts[c]++
	}

	freqs := []int{}
	for _, c := range counts {
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
