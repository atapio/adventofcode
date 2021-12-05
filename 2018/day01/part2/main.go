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
	freq := 0
	frequencies := map[int]bool{
		0: true,
	}
	found := false

	for !found {
		for _, line := range lines {
			n, err := strconv.Atoi(line)
			if err != nil {
				return 0, err
			}

			freq += n

			if _, ok := frequencies[freq]; ok {
				answer = freq
				found = true
				break
			}
			frequencies[freq] = true
		}
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

	fmt.Printf("answer: %v\n", answer)
}
