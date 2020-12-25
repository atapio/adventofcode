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

// return list of primes less than N
func sieveOfEratosthenes(N int) (primes []int) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	return
}

func findAnswer(lines []string) (int, error) {
	answer := 0
	log.Printf("generating primes")
	primes := sieveOfEratosthenes(100000000)
	log.Printf("generation done %d", len(primes))

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
	/*
		for card.Transform(7) != card.publicKey {
			card.loopSize++
		}
	*/

	log.Printf("card: %v", card)

	/*
		prime := 0
		for door.Transform(7) != door.publicKey {

			if prime%1000 == 0 {
				log.Printf("prime %d %d", prime, primes[prime])
			}

			door.loopSize = primes[prime]
			prime++
		}
	*/

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
