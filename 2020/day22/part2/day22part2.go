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

func (d Deck) Copy(cards int) *Deck {
	copy := &Deck{
		name: d.name,
	}
	for i := 0; i < cards; i++ {
		copy.cards = append(copy.cards, d.cards[i])

	}
	return copy
}

func (d *Deck) PlayCard() (int, bool) {
	c := d.cards[0]
	d.cards = d.cards[1:len(d.cards)]
	fmt.Printf("%s plays: %d\n", d.name, c)
	return c, len(d.cards) >= c
}

func (d *Deck) Play(d2 *Deck, game int) (bool, []int) {
	round := 0

	playedRounds := map[string]bool{}

	var winner *Deck

	for {
		round++
		if len(d.cards) == 0 {
			winner = d2
			break
		}
		if len(d2.cards) == 0 {
			winner = d
			break
		}

		if _, played := playedRounds[d.String()+d2.String()]; played {
			winner = d
			break
		}

		playedRounds[d.String()+d2.String()] = true

		fmt.Printf("-- Round %d (Game %d)--\n", round, game)

		fmt.Printf("%s\n", d.String())
		fmt.Printf("%s\n", d2.String())

		c1, r1 := d.PlayCard()
		c2, r2 := d2.PlayCard()

		if r1 && r2 {
			copy1 := d.Copy(c1)
			copy2 := d2.Copy(c2)
			w1, _ := copy1.Play(copy2, game+1)

			if w1 {
				d.cards = append(d.cards, c1, c2)
				fmt.Printf("%s wins the round!\n\n", d.name)
			} else {
				d2.cards = append(d2.cards, c2, c1)
				fmt.Printf("%s wins the round!\n\n", d.name)
			}
			continue
		}

		if c1 > c2 {
			d.cards = append(d.cards, c1, c2)
			fmt.Printf("%s wins the round!\n\n", d.name)
		} else {
			d2.cards = append(d2.cards, c2, c1)
			fmt.Printf("%s wins the round!\n\n", d2.name)
		}
	}

	fmt.Printf("== Post-game results (Game %d) ==\n", game)
	fmt.Printf("%s\n", d.String())
	fmt.Printf("%s\n", d2.String())

	return d == winner, winner.cards
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

	_, cards := p1.Play(p2, 1)

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
