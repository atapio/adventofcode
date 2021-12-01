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
	windows := []int{0, 0, 0}
	windowSize := 3

	for n := range lines {
		depth, err := strconv.Atoi(lines[n])
		if err != nil {
			return 0, err
		}

		for i := 0; i < windowSize; i++ {
			windows[i] += depth
		}

		if n < windowSize {
			prevDepth = depth
			windows[n%windowSize] = 0
			continue
		}

		currDepth := windows[n%windowSize]
		windows[n%windowSize] = 0

		fmt.Printf("n: %d, prev: %d, curr: %d window: %v\n", n, prevDepth, currDepth, windows)

		if currDepth > prevDepth {
			answer++
		}
		prevDepth = currDepth
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
