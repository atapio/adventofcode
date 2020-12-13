package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type EEResult struct {
	d int
	a int
	b int
}

func ExtendedEuclid(m int, n int) EEResult {
	re := EEResult{}
	if m == 0 {
		re.d = n
		re.a = 0
		re.b = 1
		return re
	}

	re = ExtendedEuclid(n%m, m)
	d := re.d
	a := re.b - (n/m)*re.a
	b := re.a
	re.d = d
	re.a = a
	re.b = b
	return re
}

func findAnswer(input []string) (int, error) {
	busLines := strings.Split(input[1], ",")

	lines := map[int]int{}
	multiple := 1

	for i, l := range busLines {
		if l == "x" {
			continue
		}

		lineNumber, err := strconv.Atoi(l)
		if err != nil {
			return 0, err
		}
		lines[i] = lineNumber
		multiple = multiple * (i + 1)
	}

	// oh noes
	//return existenseConstruction(lines), nil

	return findMinimum(lines), nil
}

func findMinimum(lines map[int]int) int {
	var a1, n1 int
	first := true

	answer := 0
	step := 1
	var keys []int
	for k := range lines {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	fmt.Printf("keys %v\n", keys)
	fmt.Printf("map %v\n", lines)

	for _, a2 := range keys {
		n2 := lines[a2]
		if first {
			first = false
			n1 = n2
			a1 = a2
			step = n1
			continue
		}

		found := false

		mult := 1

		fmt.Printf("1: %d + %d\n", n1, a1)
		fmt.Printf("2: %d + %d\n", n2, a2)
		fmt.Printf("mult: %d step %d answer %d\n", mult, step, answer)

		for !found {
			d1 := answer + (step*mult + a1)
			d2 := (d1 + a2)

			fmt.Printf("d1: %d d2: %d mult: %d m1: %d m2: %d\n", d1, d2, mult, d1%n1, d2%n2)

			if (d1%n1 == 0) &&
				(d2%n2 == 0) {
				fmt.Printf("found %d\n", mult)
				answer = d1
				step *= n2
				break
			}
			if mult > 10000 {
				panic(mult)
			}
			mult++
		}
	}
	return answer
}

func existenseConstruction(lines map[int]int) int {
	solution := 0
	a1 := 0
	n1 := 0
	first := true

	for a2, n2 := range lines {
		if first {
			first = false
			n1 = n2
			a1 = (n1 - a2) % n1
			continue
		}

		a2 = (n2 - a2) % n2

		fmt.Printf("a1 mod n1: %d mod %d\n", a1, n1)
		fmt.Printf("a2 mod n2: %d mod %d\n", a2, n2)

		ee := ExtendedEuclid(n1, n2)

		i1 := ee.b * n1
		i2 := ee.a * n2

		fmt.Printf("foo: %d\n", i1+i2)

		solution = i1*a1 + i2*a2

		a1 = a1 + a2
		n1 = n1 * n2

		//solution = ee.a*x*lines[prevMod] + ee.b*prev*lines[x]
		//solution = ee.a*x*lines[x] + ee.b*prev*lines[prevMod]
		fmt.Printf("ee %v sol: %d mul: %d\n", ee, solution, n1)
	}

	/*
		positive := solution
		n := 0

		for positive < minimum {
			n++
			positive = solution + n*n1
			fmt.Printf("n: %d, sol: %d, n1 %d pos: %d\n", n, solution, n1, positive)
		}
	*/

	return solution
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
