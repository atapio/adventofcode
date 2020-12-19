package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"text/scanner"
)

type RegexpBuilder struct {
	rules map[string]string
}

func (r RegexpBuilder) Build() string {
	return "^" + r.BuildRule("0") + "$"
}

func (r RegexpBuilder) BuildRule(key string) string {
	rule := r.rules[key]

	s := scanner.Scanner{Mode: scanner.ScanInts | scanner.ScanStrings}
	s.Init(strings.NewReader(rule))

	output := "("

	for {
		switch t := s.Scan(); t {
		case scanner.EOF:
			return output + ")"
		case scanner.String:
			output += strings.ReplaceAll(s.TokenText(), "\"", "")
		case scanner.Int:
			output += r.BuildRule(s.TokenText())
		case '|':
			output += "|"
		}
	}
}

func (r RegexpBuilder) AddRule(input string) {
	rule := strings.Split(input, ":")
	r.rules[rule[0]] = rule[1][1:]
}

func NewRegexpBuilder() *RegexpBuilder {
	r := &RegexpBuilder{rules: map[string]string{}}
	return r
}

func findAnswer(lines []string) (int, error) {
	answer := 0
	b := NewRegexpBuilder()

	rules := true
	var re *regexp.Regexp

	for _, line := range lines {
		if rules {
			if line == "" {
				rules = false
				re = regexp.MustCompile(b.Build())
				continue
			}
			b.AddRule(line)
			continue
		}

		if re.MatchString(line) {
			answer++
		}
	}

	fmt.Printf("rules: %v\n", b.rules)
	fmt.Printf("regex: %v\n", re)

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
	// remove last empty line if it exists
	/*
		if input[len(input)-1] == "" {
			input = input[:len(input)-2]
		}
	*/
	return input, nil
}

func main() {
	input, err := parseFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	answer, err := findAnswer(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("answer: %d\n", answer)
}
