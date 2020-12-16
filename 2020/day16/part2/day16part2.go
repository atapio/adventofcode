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
	Fields         map[int][]string
	FieldOrder     map[string]int
	PossibleFields []map[string]bool
	FieldNames     []string
}

func MakeTicketFormat() *TicketFormat {
	t := &TicketFormat{
		Fields:     map[int][]string{},
		FieldOrder: map[string]int{},
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
		t.FieldOrder[fieldName] = -1
	}
	for i := r2s; i <= r2e; i++ {
		t.Fields[i] = append(t.Fields[i], fieldName)
	}

	return nil
}

func (t *TicketFormat) IsTicketValid(fields []int) bool {
	for _, field := range fields {
		if len(t.Fields[field]) == 0 {
			return false
		}
	}
	return true
}

func (t *TicketFormat) FillPossibleFields(tickets [][]int) {
	mappedFields := map[string]bool{}
	for fieldName, value := range t.FieldOrder {
		if value != -1 {
			mappedFields[fieldName] = true
		}
	}

	t.PossibleFields = []map[string]bool{}

	l := len(tickets[0])
	for i := 0; i < l; i++ {
		t.PossibleFields = append(t.PossibleFields, map[string]bool{})
		first := true
		for _, ticket := range tickets {
			if first {
				first = false
				for _, valid := range t.Fields[ticket[i]] {
					if !mappedFields[valid] {
						t.PossibleFields[i][valid] = true
					}
				}
			}

		loop:
			for field := range t.PossibleFields[i] {
				for _, valid := range t.Fields[ticket[i]] {
					if valid == field {
						continue loop
					}
				}
				t.PossibleFields[i][field] = false
			}

		}
	}
}

func (t *TicketFormat) DetermineFieldOrder(tickets [][]int) []string {
	for fieldName := range t.FieldOrder {
		t.FieldNames = append(t.FieldNames, fieldName)
	}

	mapped := false
	for !mapped {
		mapped = true

		for col, fields := range t.PossibleFields {
			last := ""
			count := 0
			for name, possible := range fields {
				if possible {
					fmt.Printf("possible: %d %s\n", col, name)
					last = name
					count++
				}
			}
			if count == 1 {
				t.FieldOrder[last] = col
				mapped = false
			}
		}

		fmt.Printf("%v\n", t.PossibleFields)
		t.FillPossibleFields(tickets)
	}

	for field, pos := range t.FieldOrder {
		t.FieldNames[pos] = field
	}

	return t.FieldNames
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

	your := false
	yourTicket := []int{}
	for _, line := range lines {
		if !your && line != "your ticket:" {
			continue
		}
		if !your && line == "your ticket:" {
			your = true
			continue
		}
		fmt.Printf("found your\n")

		fieldsStr := strings.Split(line, ",")
		for _, f := range fieldsStr {
			field, err := strconv.Atoi(f)
			if err != nil {
				return 0, err
			}
			yourTicket = append(yourTicket, field)
		}
		break
	}

	nearby := false
	validTickets := [][]int{}

	for _, line := range lines {
		if !nearby && line != "nearby tickets:" {
			continue
		}
		if !nearby && line == "nearby tickets:" {
			nearby = true
			continue
		}
		fmt.Printf("found nearby\n")

		fields := []int{}

		fieldsStr := strings.Split(line, ",")
		for _, f := range fieldsStr {
			field, err := strconv.Atoi(f)
			if err != nil {
				return 0, err
			}
			fields = append(fields, field)
		}
		if !t.IsTicketValid(fields) {
			continue
		}
		validTickets = append(validTickets, fields)
	}

	t.FillPossibleFields(validTickets)

	t.DetermineFieldOrder(validTickets)

	for i, name := range t.FieldNames {
		fmt.Printf("order %s: %d\n", name, i)
	}

	answer := 1
	for i, field := range t.FieldNames {
		fmt.Printf("field: %s, value %d\n", field, yourTicket[i])
		if strings.HasPrefix(field, "departure") {
			answer *= yourTicket[i]
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
