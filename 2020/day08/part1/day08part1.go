package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var ErrInfiniteLoop = errors.New("infinite loop")

type Instruction struct {
	Count    int
	Op       string
	Argument int
}

func NewInstruction(line string) (*Instruction, error) {
	l := strings.Split(line, " ")
	arg, err := strconv.Atoi(l[1])
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s %d\n", l[0], arg)
	return &Instruction{
		Op:       l[0],
		Argument: arg,
	}, nil
}

func (i *Instruction) Exec() (int, int, error) {
	fmt.Printf("i: %v\n", i)
	i.Count++

	switch i.Op {
	case "nop":
		return 1, 0, nil
	case "jmp":
		return i.Argument, 0, nil
	case "acc":
		return 1, i.Argument, nil
	}
	return 0, 0, nil
}

type Program struct {
	Instructions []*Instruction
	Accumulator  int
	Counter      int
}

func NewProgram(lines []string) (*Program, error) {
	p := &Program{}

	for _, l := range lines {
		i, err := NewInstruction(l)
		if err != nil {
			return nil, err
		}
		p.Instructions = append(p.Instructions, i)
	}

	return p, nil
}

func (p *Program) Step() (bool, error) {
	fmt.Printf("step: %d %d\n", p.Counter, p.Accumulator)
	next := p.Instructions[p.Counter]
	if next.Count > 0 {
		return false, ErrInfiniteLoop
	}
	pc, acc, err := next.Exec()
	if err != nil {
		return false, err
	}
	p.Accumulator += acc
	p.Counter += pc

	return true, nil
}

func (p *Program) Run() (int, error) {
	cont, err := p.Step()
	for cont {
		if err != nil {
			return p.Accumulator, err
		}
		cont, err = p.Step()
	}
	return p.Accumulator, nil
}

type data struct {
	input *Program
}

func findAnswer(d data) (int, error) {
	return d.input.Run()
}

func processInput(lines []string) (data, error) {
	d := data{}
	p, err := NewProgram(lines)
	if err != nil {
		return d, err
	}
	d.input = p

	return d, err
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

	d, err := processInput(input)
	if err != nil {
		log.Fatal(err)
	}

	answer, err := findAnswer(d)
	if err != nil {
		if err != ErrInfiniteLoop {
			log.Fatal(err)
		}
	}

	fmt.Printf("answer: %d\n", answer)
}
