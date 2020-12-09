package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type XMAS struct {
	Preamble []int
	Length   int
}

func (x *XMAS) ValidateNumber(d int) bool {
	for i, v1 := range x.Preamble {
		for j, v2 := range x.Preamble {
			if i <= j {
				continue
			}

			if v1+v2 == d {
				return true
			}

		}
	}
	return false
}

func (x *XMAS) IsValid(d int) bool {
	if len(x.Preamble) == x.Length {
		b := x.ValidateNumber(d)
		if !b {
			return false
		}
		x.Preamble = x.Preamble[1:]
	}

	x.Preamble = append(x.Preamble, d)

	return true

}

type data struct {
	input []int
}

func findAnswer(d data, preamble int) (int, error) {
	x := XMAS{Length: preamble}
	for _, d := range d.input {
		valid := x.IsValid(d)
		if !valid {
			return d, nil
		}
	}
	return 0, fmt.Errorf("all valid")
}

func processInput(lines []string) (data, error) {
	d := data{}
	for _, l := range lines {
		v, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		d.input = append(d.input, v)
	}

	return d, nil
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

	answer, err := findAnswer(d, 25)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("answer: %d\n", answer)
}
