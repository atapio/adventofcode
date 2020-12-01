package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type data struct {
	input []int
}

func findAnswer(stars data) int {
	target := 2020
	for i := 0; i < len(stars.input); i++ {
		for j := 0; j < len(stars.input); j++ {
			if stars.input[i]+stars.input[j] == target {
				return stars.input[i] * stars.input[j]
			}
		}
	}
	return -1
}

func processInput(scanner *bufio.Scanner) data {
	stars := data{}
	for scanner.Scan() {
		v, err := processLine(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		stars.input = append(stars.input, v)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return stars
}

func processLine(line string) (int, error) {
	star, err := strconv.Atoi(line)
	if err != nil {
		return 0, err
	}
	return star, err
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	stars := processInput(scanner)

	answer := findAnswer(stars)

	fmt.Printf("answer: %d\n", answer)
}
