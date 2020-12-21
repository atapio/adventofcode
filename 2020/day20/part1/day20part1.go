package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	Top Direction = iota
	Right
	Bottom
	Left
)

type Tile struct {
	id          int
	content     []string
	neighbors   map[Direction]Tile
	orientation int
}

func NewTile(id int, content []string) *Tile {
	t := &Tile{
		id:        id,
		content:   content,
		neighbors: map[Direction]Tile{},
	}

	log.Printf("tile: %v\n", t)
	return t
}

func (t Tile) Edge(dir Direction) string {
	switch dir {
	case Top:
		return t.content[0]
	case Bottom:
		return t.content[len(t.content)-1]
	case Left:
		left := ""
		for _, line := range t.content {
			left += line[0:1]
		}
		return left
	case Right:
		right := ""
		for _, line := range t.content {
			l := len(line)
			right += line[l-1 : l]
		}
		return right
	}

	log.Fatalf("unknown dir %v", dir)
	return ""
}

func (t *Tile) Flip() {
	flip := []string{}
	for i := len(t.content) - 1; i >= 0; i-- {
		flip = append(flip, t.content[i])
	}
	t.content = flip
}

func (t *Tile) Rotate() {
	rotate := []string{}
	l := len(t.content)

	for y := 0; y < l; y++ {
		rotate = append(rotate, "")
		for x := 0; x < l; x++ {
			//log.Printf("x %d y %d l %d", x, y, l)
			//log.Printf("%s", t.content[l-x-1])

			rotate[y] += t.content[l-x-1][y : y+1]
		}
	}
	t.content = rotate

}

func (t *Tile) ResetOrientation() {
	for t.orientation != 0 {
		t.NextOrientation()
	}
}

func (t *Tile) NextOrientation() int {
	if t.orientation == 3 || t.orientation == 7 {
		t.Flip()
	}

	if t.orientation >= 0 && t.orientation < 3 {
		t.Rotate()
	}

	if t.orientation >= 4 && t.orientation < 7 {
		t.Rotate()
	}

	t.orientation = (t.orientation + 1) % 8

	return t.orientation
}

func (t Tile) EdgesMatch(edge Direction, other *Tile) bool {
	otherEdge := Top
	switch edge {
	case Top:
		otherEdge = Bottom
	case Bottom:
		otherEdge = Top
	case Left:
		otherEdge = Right
	case Right:
		otherEdge = Left
	}

	return t.Edge(edge) == other.Edge(otherEdge)
}

type Grid struct {
	tiles []*Tile
	size  int
}

func NewGrid() *Grid {
	g := &Grid{}
	return g
}

func (g *Grid) AddTile(tile *Tile) {
	g.tiles = append(g.tiles, tile)
	g.size = int(math.Sqrt(float64(len(g.tiles))))
}

func (g Grid) FindCorners() []*Tile {
	corners := []*Tile{}
	edges := []Direction{Top, Bottom, Left, Right}

	for _, t := range g.tiles {
		count := 0
	nexttile:
		for _, other := range g.tiles {
			if t == other {
				continue
			}

			for _, edge := range edges {
				for r := 0; r < 8; r++ {
					if t.EdgesMatch(edge, other) {
						count++
						//log.Printf("match #%d %d-%d %d", count, t.id, other.id, edge)
						continue nexttile
					}
					other.NextOrientation()
				}
			}
		}
		if count == 2 {
			corners = append(corners, t)
		}
	}

	return corners
}

func findAnswer(lines []string) (int, error) {
	answer := 0
	g := NewGrid()

	content := []string{}
	tileID := 0
	var tile *Tile

	for _, line := range lines {
		if line == "" {
			if tileID != 0 {
				tile = NewTile(tileID, content)
				g.AddTile(tile)
			}
			tileID = 0
			content = []string{}
			continue
		}
		if strings.HasPrefix(line, "Tile") {
			header := strings.Split(line, " ")
			idStr := strings.TrimRight(header[1], ":")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				return 0, err
			}
			tileID = id
			continue
		}

		content = append(content, line)
	}

	fmt.Printf("tiles: %d, size %d\n", len(g.tiles), g.size)

	corners := g.FindCorners()
	answer = 1
	for _, c := range corners {
		log.Printf("corner: %d", c.id)
		answer *= c.id
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

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
