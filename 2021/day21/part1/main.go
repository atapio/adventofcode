package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Die interface {
	Roll() int
}

type DeterministicDie struct {
	Sides int
	Rolls int
}

func (d *DeterministicDie) Roll() int {
	roll := (d.Rolls % d.Sides) + 1
	d.Rolls++
	fmt.Printf("roll: %d\n", roll)
	return roll
}

type Player struct {
	position int
	Score    int
}

type DiceGame struct {
	winningScore int
	players      []*Player
	die          Die
}

func NewDiceGame() *DiceGame {
	return &DiceGame{
		players: []*Player{},
	}
}

func (g *DiceGame) AddPlayer(position int) {
	g.players = append(g.players, &Player{position: position})
}

func (g *DiceGame) Loser() Player {
	loser := g.players[0]
	for _, p := range g.players {
		if p.Score < loser.Score {
			loser = p
		}
	}
	return *loser
}

func (g *DiceGame) Play() {
	for {
		for i, p := range g.players {
			rolls := 0
			for i := 0; i < 3; i++ {
				rolls += g.die.Roll()
			}
			p.position = (p.position + rolls) % 10
			if p.position == 0 {
				p.position = 10
			}
			p.Score += p.position
			fmt.Printf("player %d: r: %d pos: %d score: %d\n", i+1, rolls, p.position, p.Score)
			if p.Score >= g.winningScore {
				return
			}
		}
	}
}

func findAnswer(lines []string) (int, error) {
	die := &DeterministicDie{Sides: 100}
	g := NewDiceGame()
	g.winningScore = 1000
	g.die = die

	p1 := strings.TrimPrefix(lines[0], "Player 1 starting position: ")
	pos1, err := strconv.Atoi(p1)
	if err != nil {
		return 0, err
	}
	g.AddPlayer(pos1)

	p2 := strings.TrimPrefix(lines[1], "Player 2 starting position: ")
	pos2, err := strconv.Atoi(p2)
	if err != nil {
		return 0, err
	}
	g.AddPlayer(pos2)

	g.Play()

	answer := g.Loser().Score * die.Rolls

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
