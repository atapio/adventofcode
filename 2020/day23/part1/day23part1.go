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
	first   *ring.Ring
	current *ring.Ring
	len     int
}

func NewGame(cups string) *Game {
	g := &Game{
		first: ring.New(len(cups)),
		len:   len(cups),
	}
	g.current = g.first

	item := g.first

	for _, cup := range strings.Split(cups, "") {
		val, err := strconv.Atoi(cup)
		if err != nil {
			log.Fatalf("error %s", err)
		}

		item.Value = Cup{value: val}
		item = item.Next()
	}

	return g
}

func (g Game) String() string {
	cupsStr := []string{}
	item := g.first
	for i := 0; i < item.Len(); i++ {
		c := item.Value.(Cup)
		cup := fmt.Sprintf("%d", c.value)
		if item == g.current {
			cup = fmt.Sprintf("(%s)", cup)
		}
		cupsStr = append(cupsStr, cup)
		item = item.Next()
	}
	return fmt.Sprintf("cups: %s", strings.Join(cupsStr, " "))
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
	pickup := g.current
	return pickup.Unlink(3)
}
func (g Game) SelectDestination() *ring.Ring {
	target := g.current.Value.(Cup).value - 1
	if target == 0 {
		target = g.len
	}

	fmt.Printf("current %v, target %d\n", g.current.Value, target)

	search := g.current.Next()
	PrintRing(search)

	max := 0

	for {

		cup := search.Value.(Cup)
		if cup.value > max {
			max = cup.value
		}
		//fmt.Printf("target: %d value: %d\n", target, cup.value)

		if cup.value == target {
			fmt.Printf("destination: %d\n", target)
			return search
		}

		if search == g.current {
			target = (target - 1 + g.len) % g.len
			if target == 0 {
				target = max
			}
		}
		search = search.Next()
	}
}

func (g Game) PlaceCups(pickup *ring.Ring, destination *ring.Ring) {
	destination = destination.Link(pickup)
}

func (g *Game) Play(moves int) int {
	move := 0

	for move < moves {
		move++
		fmt.Printf("-- move %d --\n", move)
		fmt.Printf("%s\n", g.String())

		pickup := g.PickUp()
		fmt.Printf("pickup: ")
		PrintRing(pickup)
		destination := g.SelectDestination()
		g.PlaceCups(pickup, destination)
		g.current = g.current.Next()
	}

	return 0
}

func findAnswer(lines []string) (int, error) {
	answer := 0

	cups := lines[0]
	rounds, err := strconv.Atoi(lines[1])
	if err != nil {
		return 0, err
	}

	g := NewGame(cups)
	answer = g.Play(rounds)

	PrintRing(g.current)

	a := g.current
	for a.Value.(Cup).value != 1 {
		a = a.Next()
	}

	a = a.Next()
	ans := ""
	for a.Value.(Cup).value != 1 {
		ans = fmt.Sprintf("%s%d", ans, a.Value.(Cup).value)
		a = a.Next()
	}
	fmt.Printf("\n")

	answer, _ = strconv.Atoi(ans)

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
