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
	return &Instruction{
		Op:       l[0],
		Argument: arg,
	}, nil
}

func (i *Instruction) Exec() (int, int, error) {
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
	Original     []*Instruction
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

	//p.Original = cip.Instructions
	p.Original = make([]*Instruction, len(p.Instructions))
	for i := range p.Instructions {
		p.Original[i] = p.Instructions[i]

	}

	return p, nil
}

func (p *Program) Step() (bool, error) {
	next := p.Instructions[p.Counter]
	fmt.Printf("step: %v %d %d\n", next, p.Counter, p.Accumulator)
	if next.Count > 0 {
		return true, ErrInfiniteLoop
	}
	pc, acc, err := next.Exec()
	if err != nil {
		return false, err
	}
	p.Accumulator += acc
	p.Counter += pc

	end := p.Counter == len(p.Instructions)

	return !end, nil
}

func (p *Program) GenerateVariation(variation int) error {
	fmt.Printf("variation: %d\n", variation)
	p.Counter = 0
	p.Accumulator = 0
	//p.Instructions = p.Original
	for i := range p.Original {
		p.Instructions[i] = p.Original[i]
		p.Instructions[i].Count = 0
	}

	currentVariation := 0

	for c, i := range p.Instructions {
		fmt.Printf("op: %s var: %d curvar: %d\n", i.Op, variation, currentVariation)
		if i.Op == "nop" {
			currentVariation++
			if currentVariation == variation {
				p.Instructions[c] = &Instruction{Op: "jmp", Argument: i.Argument}
				return nil
			}
		}
		if currentVariation >= variation {
			return fmt.Errorf("no more variations")
		}

		if i.Op == "jmp" {
			currentVariation++
			if currentVariation == variation {
				p.Instructions[c] = &Instruction{Op: "nop", Argument: i.Argument}
				return nil
			}
		}
		if currentVariation >= variation {
			return fmt.Errorf("no more variations")
		}
	}
	return fmt.Errorf("no more variations")
}

func (p *Program) Run() (int, error) {
	cont, err := p.Step()
	variation := 0
	for cont {
		fmt.Printf("c: %t err: %s\n", cont, err)
		if err != nil {
			if err != ErrInfiniteLoop {
				return p.Accumulator, err
			}
			variation++
			variationErr := p.GenerateVariation(variation)
			if variationErr != nil {
				return 0, variationErr
			}
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
