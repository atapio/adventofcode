package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var fieldFormat = regexp.MustCompile(`^([a-z ]+): (?P<Range1s>[\d]+)-(?P<Range1e>[\d]+) or (?P<Range2s>[\d]+)-(?P<Range2e>[\d]+)$`)

type TicketFormat struct {
	Fields map[int][]string
}

func MakeTicketFormat() *TicketFormat {
	t := &TicketFormat{
		Fields: map[int][]string{},
	}

	return t
}

func (t *TicketFormat) AddField(line string) error {
	m := fieldFormat.FindStringSubmatch(line)
	fmt.Printf("line: %s\nregex %v\n", line, m)
	fieldName := m[1]
	r1s, err := strconv.Atoi(m[2])
	if err != nil {
		return err
	}
	r1e, err := strconv.Atoi(m[3])
	if err != nil {
		return err
	}
	r2s, err := strconv.Atoi(m[4])
	if err != nil {
		return err
	}
	r2e, err := strconv.Atoi(m[5])
	if err != nil {
		return err
	}
	fmt.Printf("add field: %s %d-%d %d-%d\n", fieldName, r1s, r1e, r2s, r2e)

	for i := r1s; i <= r1e; i++ {
		t.Fields[i] = append(t.Fields[i], fieldName)
	}
	for i := r2s; i <= r2e; i++ {
		t.Fields[i] = append(t.Fields[i], fieldName)
	}

	return nil
}

func (t *TicketFormat) IsFieldValid(field int) bool {
	return len(t.Fields[field]) > 0
}

func (t *TicketFormat) IsTicketValid(fields []int) (bool, int) {
	for _, field := range fields {
		valid := t.IsFieldValid(field)
		if !valid {
			return false, field
		}
	}
	return true, 0
}

func findAnswer(lines []string) (int, error) {
	// fields
	t := MakeTicketFormat()

	for _, line := range lines {
		if line == "" {
			break
		}
		err := t.AddField(line)
		if err != nil {
			return 0, err
		}
	}

	answer := 0
	nearby := false
	for _, line := range lines {
		if !nearby && line != "nearby tickets:" {
			continue
		}
		if !nearby && line == "nearby tickets:" {
			nearby = true
			continue
		}
		fmt.Printf("found nearby\n")

		fieldsStr := strings.Split(line, ",")
		for _, f := range fieldsStr {
			field, err := strconv.Atoi(f)
			if err != nil {
				return 0, err
			}

			if !t.IsFieldValid(field) {
				fmt.Printf("not valid %d\n", field)

				answer += field
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
	// remove last empty line if it exists
	if input[len(input)-1] == "" {
		input = input[:len(input)-2]
	}
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
