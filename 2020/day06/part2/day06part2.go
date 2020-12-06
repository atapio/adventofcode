package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type CustomsGroup struct {
	Persons int
	Answers map[string]int
}

func NewCustomsGroup() *CustomsGroup {
	cg := &CustomsGroup{
		Answers: map[string]int{},
	}

	return cg
}

func (g *CustomsGroup) AnswerCount() int {
	count := 0
	for _, v := range g.Answers {
		if v == g.Persons {
			count++
		}
	}

	return count
}

func (g *CustomsGroup) AddPerson(data string) {
	g.Persons++

	answers := strings.Split(data, "")
	for _, a := range answers {
		g.Answers[a]++
	}
}

type data struct {
	input []*CustomsGroup
}

func findAnswer(d data) int {
	count := 0

	for _, g := range d.input {
		count = count + g.AnswerCount()
	}

	return count
}

func processInput(lines []string) data {
	d := data{}

	cg := NewCustomsGroup()
	d.input = append(d.input, cg)

	for _, line := range lines {
		switch line {
		case "":
			cg = NewCustomsGroup()
			d.input = append(d.input, cg)
		default:
			cg.AddPerson(line)
		}
	}

	return d
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
