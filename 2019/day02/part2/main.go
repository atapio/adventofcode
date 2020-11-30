package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func exec(pc int, program []int) ([]int, int, bool) {
	c := 0
	op := program[pc]
	if op == 99 {
		return program, 1, true
	}

	v1 := program[program[pc+1]]
	v2 := program[program[pc+2]]
	r := program[pc+3]

	switch op {
	case 1:
		program[r] = v1 + v2
		c = 4
	case 2:
		program[r] = v1 * v2
		c = 4
	default:
		fmt.Printf("unknown opcode %d", op)
	}

	return program, c, false
}

func runProg(program []int) []int {
	p := make([]int, len(program))
	copy(p, program)
	pc := 0
	ic := 0
	halt := false
	for !halt {
		p, ic, halt = exec(pc, p)
		pc = pc + ic
	}

	return p
}

func iterate(program []int, target int) (int, int) {
	for noun := 0; noun < 100; noun++ {
		program[1] = noun
		for verb := 0; verb < 100; verb++ {
			program[2] = verb
			fmt.Printf("%d/%d\n", noun, verb)
			p := runProg(program)
			if p[0] == target {
				return noun, verb
			}
		}
	}
	return -1, -1
}

func compute(programStr string, target int) ([]int, error) {
	opcodes := strings.Split(programStr, ",")

	program := make([]int, len(opcodes))
	for i, opcode := range opcodes {
		parsed, err := strconv.Atoi(opcode)
		if err != nil {
			return []int{}, err
		}
		program[i] = parsed
	}

	noun := 0
	verb := 0
	if target > 0 {
		noun, verb = iterate(program, target)
		fmt.Printf("found: %d\n", 100*noun+verb)
	}

	return program, nil
}

func processLine(line string, target int) ([]int, error) {
	return compute(line, target)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_, err = processLine(scanner.Text(), 19690720)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
