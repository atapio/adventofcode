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

func validPaths(jolts []int) int {
	pathCount := map[int]int{0: 1}

	for _, j := range jolts {
		pathCount[j] = pathCount[j-1] + pathCount[j-2] + pathCount[j-3]
		fmt.Printf("path: %d:%d\n", j, pathCount[j])
	}

	last := jolts[len(jolts)-1]

	return pathCount[last]
}

func findAnswer(d data) (int, error) {
	count := validPaths(d.input)

	return count, nil
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
