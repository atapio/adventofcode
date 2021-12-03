package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Diagnostic struct {
	oxygen int
	co2    int
}

func NewDiagnostic() *Diagnostic {
	return &Diagnostic{}
}

func (d *Diagnostic) Process(input []string) {
	l := len(input[0])

	oxygen := input
	co2 := input

	for i := 0; i < l; i++ {
		ones, zeros := d.split(oxygen, i)

		oxygen = zeros

		if len(ones) >= len(zeros) {
			oxygen = ones
		}

		fmt.Printf("o2: %d, ones: %d zeros: %d\n", len(oxygen), len(ones), len(zeros))

		ones, zeros = d.split(co2, i)
		co2 = ones

		if len(zeros) <= len(ones) {
			co2 = zeros
		}
		fmt.Printf("co2: %d, ones: %d zeros: %d\n", len(co2), len(ones), len(zeros))
	}

	o2, err := strconv.ParseInt(oxygen[0], 2, 32)
	if err != nil {
		log.Fatal(err)
	}

	co2int, err := strconv.ParseInt(co2[0], 2, 32)
	if err != nil {
		log.Fatal(err)
	}

	d.oxygen = int(o2)
	d.co2 = int(co2int)
}

func (d *Diagnostic) split(input []string, pos int) ([]string, []string) {
	if len(input) == 1 {
		return input, input
	}

	ones := []string{}
	zeros := []string{}

	for _, line := range input {
		chars := strings.Split(line, "")
		if chars[pos] == "1" {
			ones = append(ones, line)
		} else {
			zeros = append(zeros, line)
		}
	}
	return ones, zeros
}

func findAnswer(lines []string) (int, error) {
	answer := 0

	d := NewDiagnostic()
	d.Process(lines)

	answer = d.co2 * d.oxygen

	return answer, nil
}

func parseFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	input := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		input = append(input, l)
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}
	return input, nil
}

func main() {
	input, err := parseFile("input")
	if err != nil {
		log.Fatal(err)
	}

	answer, err := findAnswer(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("answer: %d\n", answer)
}
