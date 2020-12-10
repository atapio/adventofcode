package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type data struct {
	input []int
}

func findAnswer(d data) (int, error) {
	differences := map[int]int{}

	prev := 0

	for _, jolts := range d.input {
		fmt.Printf("%d %d\n", jolts, jolts-prev)
		differences[jolts-prev]++
		prev = jolts
	}
	return differences[1] * differences[3], nil
}

func processInput(lines []string) (data, error) {
	d := data{}

	max := 0
	for _, l := range lines {
		v, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		if max < v {
			max = v
		}
		d.input = append(d.input, v)
	}

	d.input = append(d.input, max+3)

	sort.Ints(d.input)

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

	answer, err := findAnswer(d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("answer: %d\n", answer)
}
