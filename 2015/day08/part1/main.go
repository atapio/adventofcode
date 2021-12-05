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
		answer += len(line)
		unescaped := strings.Trim(line, "\"")
		//unescaped = strings.ReplaceAll(unescaped, "\\\\", "\\")
		unescaped = strings.ReplaceAll(unescaped, "\\\\", "X")
		//unescaped = strings.ReplaceAll(unescaped, "\\\"", "\"")
		unescaped = strings.ReplaceAll(unescaped, "\\\"", "X")
		unescaped = hexFormat.ReplaceAllString(unescaped, "X")

		answer -= len(unescaped)

		fmt.Printf("o: %s\nu: %s\n", line, unescaped)

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
