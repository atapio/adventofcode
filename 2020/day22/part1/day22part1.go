package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Deck struct {
	name  string
	cards []int
}

func (d Deck) String() string {
	cardsStr := []string{}
	for _, c := range d.cards {
		cardsStr = append(cardsStr, fmt.Sprintf("%d", c))
	}
	return fmt.Sprintf("%s's deck: %s", d.name, strings.Join(cardsStr, ", "))
}

func (d *Deck) PlayCard() int {
	c := d.cards[0]
	d.cards = d.cards[1:len(d.cards)]
	fmt.Printf("%s plays: %d\n", d.name, c)
	return c
}

func (d *Deck) Play(d2 *Deck) []int {
	round := 0

	var winner *Deck

	for {
		round++
		if len(d.cards) == 0 {
			fmt.Printf("%s wins the round!", d2.name)
			winner = d2
			break
		}
		if len(d2.cards) == 0 {
			fmt.Printf("%s wins the round!", d.name)
			winner = d
			break
		}

		fmt.Printf("-- Round %d --\n", round)

		fmt.Printf("%s\n", d.String())
		fmt.Printf("%s\n", d2.String())

		c1, c2 := d.PlayCard(), d2.PlayCard()
		if c1 > c2 {
			d.cards = append(d.cards, c1, c2)
			fmt.Printf("%s wins the round!\n\n", d.name)
		} else {
			d2.cards = append(d2.cards, c2, c1)
			fmt.Printf("%s wins the round!\n\n", d2.name)
		}
	}

	fmt.Println("== Post-game results ==")
	fmt.Printf("%s\n", d.String())
	fmt.Printf("%s\n", d2.String())

	return winner.cards
}

func findAnswer(lines []string) (int, error) {
	answer := 0

	players := []*Deck{}
	var current *Deck

	for _, line := range lines {
		if strings.HasPrefix(line, "Player") {
			current = &Deck{
				name: strings.Trim(line, ":"),
			}
			players = append(players, current)
			continue
		}

		if line == "" {
			continue
		}

		card, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}

		current.cards = append(current.cards, card)
	}

	p1 := players[0]
	p2 := players[1]

	cards := p1.Play(p2)

	multiplier := len(cards)

	for _, c := range cards {
		answer += c * multiplier
		multiplier--
	}

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
