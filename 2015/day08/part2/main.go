package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var hexFormat = regexp.MustCompile(`\\x..`)

func findAnswer(lines []string) (int, error) {
	answer := 0

	for _, line := range lines {
		answer -= len(line)
		escaped := strings.ReplaceAll(line, "\"", "XX")
		escaped = strings.ReplaceAll(escaped, "\\", "XX")
		escaped = fmt.Sprintf("\"%s\"", escaped)

		answer += len(escaped)

		fmt.Printf("o: %s\nu: %s\n", line, escaped)

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
