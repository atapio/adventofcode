package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type BoardingPass struct {
	Row    int
	Column int
}

func (b *BoardingPass) SeatID() int {
	return b.Row*8 + b.Column
}

func NewBoardingPass(data string) *BoardingPass {
	b := &BoardingPass{}

	rowData := data[0:7]
	columnData := data[7:10]

	for _, f := range strings.Split(rowData, "") {
		b.Row = b.Row << 1
		if f == "B" {
			b.Row++
		}
	}

	for _, f := range strings.Split(columnData, "") {
		b.Column = b.Column << 1
		if f == "R" {
			b.Column++
		}
	}
	fmt.Printf("%s %s %+v\n", rowData, columnData, b)

	return b
}

type data struct {
	input []*BoardingPass
}

func findAnswer(d data) int {
	max := 0

	for _, p := range d.input {
		if id := p.SeatID(); id > max {
			max = id
		}
	}

	return max
}

func processInput(lines []string) data {
	d := data{}

	for _, line := range lines {
		d.input = append(d.input, NewBoardingPass(line))
	}

	return d
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

	d := processInput(input)

	answer := findAnswer(d)

	fmt.Printf("answer: %d\n", answer)
}
