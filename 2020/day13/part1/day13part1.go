package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Bus struct {
	Depart    int
	InService bool
}

func (b *Bus) WaitingTime(time int) int {
	wait := b.Depart - time%b.Depart
	return wait
}

func findAnswer(lines []string) (int, error) {
	buses := []Bus{}

	busLines := strings.Split(lines[1], ",")

	for _, l := range busLines {
		switch l {
		case "x":
			continue
		default:
			lineNumber, err := strconv.Atoi(l)
			if err != nil {
				return 0, err
			}
			buses = append(buses, Bus{Depart: lineNumber})
		}
	}

	departAt, err := strconv.Atoi(lines[0])
	if err != nil {
		return 0, err
	}

	minDepart := 100000
	answer := 0

	for _, b := range buses {
		w := b.WaitingTime(departAt)
		if w < minDepart {
			fmt.Printf("depart: %d, b: %d w: %d\n", departAt, b.Depart, w)
			minDepart = w
			answer = w * b.Depart
		}
	}

	return answer, nil
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

	answer, err := findAnswer(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("answer: %d\n", answer)
}
