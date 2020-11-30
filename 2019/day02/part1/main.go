package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func exec(pc int, program []int) ([]int, bool) {
	op := program[pc]
	if op == 99 {
		return program, true
	}

	v1 := program[program[pc+1]]
	v2 := program[program[pc+2]]
	r := program[pc+3]

	switch op {
	case 1:
		program[r] = v1 + v2
	case 2:
		program[r] = v1 * v2
	}

	return program, false
}

func compute(programStr string, replace bool) ([]int, error) {
	opcodes := strings.Split(programStr, ",")

	program := make([]int, len(opcodes))
	for i, opcode := range opcodes {
		parsed, err := strconv.Atoi(opcode)
		if err != nil {
			return []int{}, err
		}
		program[i] = parsed
	}

	if replace {
		program[1] = 12
		program[2] = 2
	}

	pc := 0
	halt := false
	for !halt {
		program, halt = exec(pc, program)
		pc = pc + 4
	}

	return program, nil
}

func processLine(line string, current []int) ([]int, error) {
	return compute(line, true)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	prog := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		prog, err = processLine(scanner.Text(), prog)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result: %v\n", prog)
}
