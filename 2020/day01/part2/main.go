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
			if i == j {
				continue
			}

			for k := 0; k < len(stars.input); k++ {
				if i == k {
					continue
				}
				if j == k {
					continue
				}

				if stars.input[i]+stars.input[j]+stars.input[k] == target {
					return stars.input[i] * stars.input[j] * stars.input[k]
				}
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
