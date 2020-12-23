package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Cup struct {
	value    int
	pickedUp bool
}

type Game struct {
	current *ring.Ring
	byValue map[int]*ring.Ring
	len     int
}

func NewGame(cups string, size int) *Game {
	g := &Game{
		current: ring.New(size),
		len:     size,
		byValue: map[int]*ring.Ring{},
	}

	item := g.current

	for _, cup := range strings.Split(cups, "") {
		val, err := strconv.Atoi(cup)
		if err != nil {
			log.Fatalf("error %s", err)
		}

		item.Value = Cup{value: val}
		g.byValue[val] = item
		item = item.Next()
	}

	rest := len(cups)
	for rest < size {
		rest++
		item.Value = Cup{value: rest}
		g.byValue[rest] = item
		item = item.Next()
	}

	return g
}

func PrintRing(r *ring.Ring) {
	cupsStr := []string{}
	item := r
	for i := 0; i < r.Len(); i++ {
		cupsStr = append(cupsStr, fmt.Sprintf("%d", item.Value.(Cup).value))
		item = item.Next()
	}
	fmt.Printf("cups: %s\n", strings.Join(cupsStr, " "))
}

func (g *Game) PickUp() *ring.Ring {
	pickup := g.current.Unlink(3)

	loop := pickup
	for i := 0; i < 3; i++ {
		val := loop.Value.(Cup).value
		delete(g.byValue, val)
		loop = loop.Next()
	}

	return pickup
}

func (g Game) SelectDestination() *ring.Ring {
	target := g.current.Value.(Cup).value - 1
	if target == 0 {
		target = g.len
	}

	//fmt.Printf("current %v, target %d\n", g.current.Value, target)

	//PrintRing(search)

	max := g.len

	for {
		search, found := g.byValue[target]

		if found {
			return search
		}

		if target == max {
			target--
		}

		target = (target - 1 + g.len) % g.len
		if target == 0 {
			target = max
		}
	}
}

func (g Game) PlaceCups(pickup *ring.Ring, destination *ring.Ring) {
	len := pickup.Len()
	destination.Link(pickup)
	loop := pickup
	for i := 0; i < len; i++ {
		cup := loop.Value.(Cup)
		cup.pickedUp = false
		loop.Value = cup

		g.byValue[cup.value] = loop
		loop = loop.Next()
	}
}

func (g *Game) Play(moves int) int {
	move := 0

	for move < moves {
		move++
		if move%100000 == 1 {
			fmt.Printf("-- move %d --\n", move)
		}

		pickup := g.PickUp()
		destination := g.SelectDestination()
		g.PlaceCups(pickup, destination)
		g.current = g.current.Next()
	}

	return 0
}

func findAnswer(lines []string) (int, error) {
	answer := 0

	cups := lines[0]
	rounds, err := strconv.Atoi(lines[2])
	if err != nil {
		return 0, err
	}

	g := NewGame(cups, 1000000)
	answer = g.Play(rounds)

	fmt.Printf("game played\n")

	a := g.byValue[1]
	a = a.Next()
	fmt.Printf("first: %v\n", a.Value.(Cup))
	answer = a.Value.(Cup).value
	a = a.Next()
	fmt.Printf("second: %v\n", a.Value.(Cup))
	answer *= a.Value.(Cup).value

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
