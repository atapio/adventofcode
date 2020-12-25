package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const modulo = 20201227

type PKI struct {
	publicKey int
	loopSize  int
}

func (p PKI) Transform(subject int) int {
	val := 1
	for i := 0; i < p.loopSize; i++ {
		val *= subject
		val = val % modulo
	}
	return val
}

func (p *PKI) DetermineLoopSize(subject int) {
	val := 1
	p.loopSize = 0
	for {
		p.loopSize++
		val *= subject
		val = val % modulo
		if val == p.publicKey {
			return
		}
	}
}

func findAnswer(lines []string) (int, error) {
	answer := 0

	cardPubKey, err := strconv.Atoi(lines[0])
	if err != nil {
		return 0, err
	}
	doorPubKey, err := strconv.Atoi(lines[1])
	if err != nil {
		return 0, err
	}

	card := PKI{publicKey: cardPubKey}
	door := PKI{publicKey: doorPubKey}

	card.DetermineLoopSize(7)

	log.Printf("card: %v", card)

	door.DetermineLoopSize(7)
	log.Printf("door: %v", door)

	doorEnc := card.Transform(door.publicKey)
	cardEnc := door.Transform(card.publicKey)

	log.Printf("card: %d", cardEnc)
	log.Printf("door: %d", doorEnc)

	answer = doorEnc

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
	input, err := parseFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	answer, err := findAnswer(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("answer: %d\n", answer)
}
