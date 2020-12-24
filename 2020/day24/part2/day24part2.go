package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var tileFormat = regexp.MustCompile(`(se|e|sw|nw|ne|w)`)

type Coordinate struct {
	x int
	y int
}

func (c Coordinate) Move(direction string) Coordinate {
	switch direction {
	case "ne":
		return Coordinate{c.x + 1, c.y + 1}
	case "e":
		return Coordinate{c.x + 2, c.y}
	case "se":
		return Coordinate{c.x + 1, c.y - 1}
	case "sw":
		return Coordinate{c.x - 1, c.y - 1}
	case "w":
		return Coordinate{c.x - 2, c.y}
	case "nw":
		return Coordinate{c.x - 1, c.y + 1}
	}

	log.Fatalf("invalid direction %s", direction)
	return Coordinate{}
}
func (c Coordinate) Around() []Coordinate {
	dirs := []string{"ne", "e", "se", "nw", "w", "sw"}
	coords := []Coordinate{}
	for _, dir := range dirs {
		coords = append(coords, c.Move(dir))
	}
	return coords
}

type Floor struct {
	tiles map[Coordinate]bool
}

func NewFloor() *Floor {
	f := &Floor{
		tiles: map[Coordinate]bool{},
	}

	return f
}

func (f *Floor) FlipTile(c Coordinate) {
	log.Printf("flip %v %t", c, f.tiles[c])
	f.tiles[c] = !f.tiles[c]
}

func (f *Floor) DailyFlip() {
	newTiles := map[Coordinate]bool{}

	for c, black := range f.tiles {
		blacks := f.CountBlacksAround(c)
		newTiles[c] = black

		if black {
			if blacks == 0 || blacks > 2 {
				newTiles[c] = false
			}

			for _, n := range c.Around() {
				if !f.tiles[n] {
					blacks := f.CountBlacksAround(n)
					if blacks == 2 {
						newTiles[n] = true
					}
				}
			}
		} else {
			if blacks == 2 {
				newTiles[c] = true
			}
		}
	}

	log.Printf("blacks %d", f.CountBlack())
	f.tiles = newTiles
}

func (f *Floor) CountBlack() int {
	count := 0
	for _, black := range f.tiles {
		if black {
			count++
		}
	}

	return count
}

func (f *Floor) CountBlacksAround(c Coordinate) int {
	count := 0
	for _, dir := range c.Around() {
		if f.tiles[dir] {
			count++
		}
	}

	return count
}

func findAnswer(lines []string) (int, error) {
	answer := 0

	floor := NewFloor()

	for _, line := range lines {
		m := tileFormat.FindAllStringSubmatch(line, -1)

		c := Coordinate{}

		for _, match := range m {
			c = c.Move(match[0])
		}

		floor.FlipTile(c)
		log.Printf("black: %d", floor.CountBlack())
	}
	log.Printf("floor: %v", floor.tiles)

	for i := 0; i < 100; i++ {
		floor.DailyFlip()
		if (i+1)%10 == 0 {
			log.Printf("floor: %v", floor.tiles)
		}
		log.Printf("day: %d: %d", i+1, floor.CountBlack())
	}

	answer = floor.CountBlack()

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
