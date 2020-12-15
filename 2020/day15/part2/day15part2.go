package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var maskFormat = regexp.MustCompile(`^mask = (?P<Value>[X01]+)$`)
var instructionFormat = regexp.MustCompile(`^mem\[(?P<Address>\d+)\] = (?P<Value>\d+)$`)

type Game struct {
	Numbers    map[int][]int
	LastNumber int
	Turn       int
}

func MakeGame(startingNumbers []int) *Game {
	g := &Game{
		Numbers: map[int][]int{},
	}

	for _, n := range startingNumbers {
		g.Turn++
		g.Numbers[n] = append(g.Numbers[n], g.Turn)
		g.LastNumber = n
	}

	return g
}

func (g *Game) SayNumber() int {
	if turns, ok := g.Numbers[g.LastNumber]; ok {
		g.Turn++
		switch len(turns) {
		case 0:
			panic("zero length")
		case 1:
			g.LastNumber = 0
		default:
			g.LastNumber = turns[len(turns)-1] - turns[len(turns)-2]
		}

		g.Numbers[g.LastNumber] = append(g.Numbers[g.LastNumber], g.Turn)
		return g.LastNumber
	}

	panic("number not found")
}

func findAnswer(lines []string) (int, error) {
	numbersStr := strings.Split(lines[0], ",")
	numbers := []int{}

	for _, n := range numbersStr {
		num, err := strconv.Atoi(n)
		if err != nil {
			return 0, err
		}

		numbers = append(numbers, num)
	}

	g := MakeGame(numbers)

	end := 30000000
	answer := 0

	for g.Turn < end {
		answer = g.SayNumber()
		if g.Turn%10000 == 0 {
			fmt.Printf("%d: %d\n", g.Turn, answer)
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
