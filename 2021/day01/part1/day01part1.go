package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func findAnswer(lines []string) (int, error) {
	answer := 0

	prevDepth := 0

	for n := range lines {
		depth, err := strconv.Atoi(lines[n])
		if err != nil {
			return 0, err
		}
		if n == 0 || depth <= prevDepth {
			prevDepth = depth
			continue
		}

		answer++
		prevDepth = depth
	}

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
