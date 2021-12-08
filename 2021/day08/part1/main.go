package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func findAnswer(lines []string) (int, error) {
	answer := 0

	for _, line := range lines {
		input := strings.Split(line, "|")
		//tests := input[0]
		data := input[1]

		for _, digit := range strings.Fields(data) {
			switch len(digit) {
			case 2:
				answer++
			case 3:
				answer++
			case 4:
				answer++
			case 7:
				answer++
			}
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

	fmt.Printf("answer: %d\n", answer)
}
