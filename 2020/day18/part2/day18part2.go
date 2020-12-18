package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/scanner"
)

type Expression struct {
	op    string
	Value int
}

func (e *Expression) AddInputString(input string) {
	i, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalf("err %v", err)
	}
	e.AddInput(i)
}

func (e *Expression) AddInput(i int) {
	e.Value = e.calculate(i)
}

func (e *Expression) AddOp(op string) {
	e.op = op
}

func (e Expression) calculate(input int) int {
	switch e.op {
	case "+":
		return e.Value + input
	case "-":
		return e.Value - input
	case "*":
		return e.Value * input
	case "":
		return input
	}

	log.Fatalf("unknown op %s", e.op)
	return 0
}

type Parser struct {
	scanner scanner.Scanner
}

var s = scanner.Scanner{Mode: scanner.ScanInts}

func MakeParser() *Parser {
	p := &Parser{}
	p.scanner = scanner.Scanner{Mode: scanner.ScanInts}
	return p
}

func (p Parser) Parse(input string) int {
	p.scanner.Init(strings.NewReader(input))

	fmt.Printf("in: %s\n", input)
	result := p.parseExpression()
	fmt.Printf("out: %d\n", result)

	return result
}

func (p Parser) parseExpression() int {
	mults := []int{}

	e := Expression{}
	nesting := 0
loop:
	for {
		fmt.Printf("exp: %v, nest %d, mults: %v\n", e, nesting, mults)
		switch t := p.scanner.Scan(); t {
		case scanner.EOF:
			break loop
		case scanner.Int:
			if nesting == 0 {
				fmt.Printf("int: %s\n", p.scanner.TokenText())
				e.AddInputString(p.scanner.TokenText())
			}
		case '(':
			if nesting == 0 {
				e.AddInput(p.parseExpression())
			}
			nesting++
		case ')':
			if nesting == 0 {
				break loop
			}
			nesting--
		case '*':
			if nesting == 0 {
				fmt.Printf("op: %s\n", string(t))
				mults = append(mults, e.Value)
				e = Expression{}
			}
		case '+':
			if nesting == 0 {
				fmt.Printf("op: %s\n", string(t))
				e.AddOp(string(t))
			}
		}
	}
	val := e.Value
	for _, a := range mults {
		val *= a
	}
	return val
}

func findAnswer(lines []string) (int, error) {
	answer := 0
	for _, line := range lines {
		parser := MakeParser()
		answer += parser.Parse(line)
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
