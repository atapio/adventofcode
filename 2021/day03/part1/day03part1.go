package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Diagnostic struct {
	total int
	ones  []int
}

func NewDiagnostic(width int) *Diagnostic {
	return &Diagnostic{
		ones: make([]int, width),
	}
}

func (d *Diagnostic) Next(input string) {
	chars := strings.Split(input, "")
	d.total++
	for i, b := range chars {
		if b == "1" {
			d.ones[i]++
		}
	}
}

func (d *Diagnostic) Gamma() int {
	g := 0

	for i := range d.ones {
		g = g << 1

		if d.ones[i] > d.total/2 {
			g++
		}
	}
	return g
}

func (d *Diagnostic) Epsilon() int {
	e := 0

	for i := range d.ones {
		e = e << 1

		if d.ones[i] < d.total/2 {
			e++
		}
	}
	return e
}

func findAnswer(lines []string) (int, error) {
	answer := 0

	d := NewDiagnostic(len(lines[0]))

	for n := range lines {
		d.Next(lines[n])
		fmt.Printf("total %d ones: %v\n", d.total, d.ones)
	}

	answer = d.Gamma() * d.Epsilon()

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
