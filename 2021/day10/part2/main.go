package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

/*
   ): 3 points.
   ]: 57 points.
   }: 1197 points.
   >: 25137 points.
*/
func ErrorScore(char string) int {
	switch char[0] {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	}

	return 0
}

func IncompleteScore(stack []string) int {
	score := 0
	for i := len(stack) - 1; i >= 0; i-- {
		score *= 5
		switch stack[i] {
		case "(":
			score += 1
		case "[":
			score += 2
		case "{":
			score += 3
		case "<":
			score += 4
		}
	}

	return score
}

func findAnswer(lines []string) (int, error) {
	scores := []int{}

	for _, line := range lines {
		corrupt := false

		stack := []string{}
		fmt.Printf("line: %s\n", line)

		for _, c := range strings.Split(line, "") {
			var pop string
			l := len(stack)
			switch c {
			case "(":
				stack = append(stack, c)
			case "[":
				stack = append(stack, c)
			case "{":
				stack = append(stack, c)
			case "<":
				stack = append(stack, c)
			case ")":
				stack, pop = stack[:l-1], stack[l-1]
				fmt.Printf("stack: %v, char: %s pop: %s\n", stack, c, pop)
				if pop != "(" {
					corrupt = true
					break
				}
			case "]":
				stack, pop = stack[:l-1], stack[l-1]
				fmt.Printf("stack: %v, char: %s pop: %s\n", stack, c, pop)
				if pop != "[" {
					corrupt = true
					break
				}
			case "}":
				stack, pop = stack[:l-1], stack[l-1]
				fmt.Printf("stack: %v, char: %s pop: %s\n", stack, c, pop)
				if pop != "{" {
					corrupt = true
					break
				}
			case ">":
				stack, pop = stack[:l-1], stack[l-1]
				fmt.Printf("stack: %v, char: %s pop: %s\n", stack, c, pop)
				if pop != "<" {
					corrupt = true
					break
				}
			}
		}
		if !corrupt && len(stack) > 0 {
			score := IncompleteScore(stack)
			scores = append(scores, score)
			fmt.Printf("incomplete: %v, score: %d\n", stack, score)
		}
	}

	sort.Ints(scores)
	answer := scores[len(scores)/2]

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
