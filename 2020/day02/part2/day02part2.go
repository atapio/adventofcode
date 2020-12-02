package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var lineFormat = regexp.MustCompile(`^(?P<Min>\d+)-(?P<Max>\d+) (?P<char>[a-z]): (?P<Password>[a-z]+)$`)

type Password struct {
	pos1     int
	pos2     int
	char     string
	password string
}

func (p *Password) isValid() bool {
	count := 0

	if p.password[p.pos1:p.pos1+1] == p.char {
		count++
	}
	if p.password[p.pos2:p.pos2+1] == p.char {
		count++
	}

	return count == 1
}

type data struct {
	input []Password
}

func findAnswer(d data) int {
	valid := 0
	for _, p := range d.input {
		if p.isValid() {
			valid++
		}
	}
	return valid
}

func processInput(lines []string) data {
	d := data{}
	for _, line := range lines {
		v, err := processLine(line)
		if err != nil {
			log.Fatal(err)
		}
		d.input = append(d.input, v)
	}

	return d
}

func processLine(line string) (Password, error) {
	m := lineFormat.FindStringSubmatch(line)

	pos1, err := strconv.Atoi(m[1])
	if err != nil {
		return Password{}, err
	}
	pos2, err := strconv.Atoi(m[2])
	if err != nil {
		return Password{}, err
	}

	p := Password{
		pos1:     pos1 - 1,
		pos2:     pos2 - 1,
		char:     m[3],
		password: m[4],
	}
	return p, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		input = append(input, l)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	d := processInput(input)

	answer := findAnswer(d)

	fmt.Printf("answer: %d\n", answer)
}
