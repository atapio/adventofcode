package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var seaMonster = []string{
	"                  # ",
	"#    ##    ##    ###",
	" #  #  #  #  #  #   ",
}

var seaMonsterRegexp = []*regexp.Regexp{
	regexp.MustCompile(`^.{18}#`),
	regexp.MustCompile(`^#.{4}#{2}.{4}#{2}.{4}#{3}`),
	regexp.MustCompile(`^.#..#..#..#..#..#`),
}

type Coordinate struct {
	x int
	y int
}

type Direction Coordinate

var (
	Top    Direction = Direction{0, -1}
	Right            = Direction{1, 0}
	Bottom           = Direction{0, 1}
	Left             = Direction{-1, 0}
)

type Tile struct {
	id          int
	content     []string
	neighbors   map[Direction]*Tile
	orientation int
}

func NewTile(id int, content []string) *Tile {
	t := &Tile{
		id:        id,
		content:   content,
		neighbors: map[Direction]*Tile{},
	}

	log.Printf("tile: %v\n", t)
	return t
}

func (t Tile) String() string {
	str := fmt.Sprintf("tile %d, orientation: %d\n", t.id, t.orientation)
	for edge, tile := range t.neighbors {
		str += fmt.Sprintf("neighbor: %v: %d\n", edge, tile.id)
	}

	for _, l := range t.content {
		str += fmt.Sprintf("%s\n", l)
	}
	return str
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

func (t Tile) OppositeEdge(edge Direction) Direction {
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
	return otherEdge
}

func (t Tile) EdgesMatch(edge Direction, other *Tile) bool {
	otherEdge := t.OppositeEdge(edge)

	return t.Edge(edge) == other.Edge(otherEdge)
}

func (t *Tile) FindConnectingTile(edge Direction, tiles map[int]*Tile) (*Tile, bool) {
	for _, tile := range tiles {
		for r := 0; r < 8; r++ {
			if t.EdgesMatch(edge, tile) {
				log.Printf("edge %v match %d/%d-%d/%d", edge, t.id, t.orientation, tile.id, tile.orientation)
				t.neighbors[edge] = tile
				tile.neighbors[tile.OppositeEdge(edge)] = t
				return tile, true
			}
			tile.NextOrientation()
		}
	}
	return nil, false
}

func (t Tile) CountSeaMonsters() int {
	count := 0

	size := len(t.content)

	for r := 0; r < 8; r++ {
		log.Printf("counting orientation %d", t.orientation)
		for row := 0; row < size-len(seaMonster); row++ {
		nextcol:
			for col := 0; col < size-len(seaMonster[0]); col++ {
				for sm := 0; sm < len(seaMonster); sm++ {
					//log.Printf("matching row %d\n'%s' to \n'%s'", row, t.content[row][col:len(t.content[row])], seaMonster[sm])
					if !seaMonsterRegexp[sm].MatchString(t.content[row+sm][col:len(t.content[row])]) {
						continue nextcol
					}
					//log.Printf("match %d", sm)
				}
				count++
			}
		}
		t.NextOrientation()
	}

	return count
}

type Grid struct {
	tiles  []*Tile
	layout map[Coordinate]*Tile
	size   int
}

func NewGrid() *Grid {
	g := &Grid{layout: map[Coordinate]*Tile{}}
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

func (g *Grid) Assemble() {
	corners := g.FindCorners()

	tilesLeft := map[int]*Tile{}

	for _, tile := range g.tiles {
		tilesLeft[tile.id] = tile
	}

	topLeft := corners[0]
	delete(tilesLeft, topLeft.id)

	// orient top left
nextorientation:
	for i := 0; i < 8; i++ {
		topLeftEdges := []Direction{Right, Bottom}
		for _, edge := range topLeftEdges {
			if _, found := topLeft.FindConnectingTile(edge, tilesLeft); !found {
				topLeft.NextOrientation()
				continue nextorientation
			}
		}
	}

	// build grid
	g.layout = map[Coordinate]*Tile{}
	c := Coordinate{}
	g.layout[c] = topLeft

	tile := topLeft

	for y := 0; y < g.size; y++ {
		rowFirst := tile

		for x := 1; x < g.size; x++ {
			right, found := tile.FindConnectingTile(Right, tilesLeft)
			if !found {
				log.Fatalf("not found right tile %d", tile.id)
			}
			c := Coordinate{x, y}
			g.layout[c] = right
			tile = right
		}

		if y+1 == g.size {
			// all done
			break
		}

		bottom, found := rowFirst.FindConnectingTile(Bottom, tilesLeft)
		if !found {
			log.Fatalf("not found tile %d", rowFirst.id)
		}
		c := Coordinate{0, y + 1}
		g.layout[c] = bottom
		tile = bottom
	}

	log.Printf("layout %v", g.layout)
	for y := 0; y < g.size; y++ {
		for x := 0; x < g.size; x++ {
			log.Printf("%d %d %d", x, y, g.layout[Coordinate{x, y}].id)
		}
	}
}

func (g Grid) ToImage() *Tile {
	tileSize := len(g.layout[Coordinate{}].content[0]) - 2

	content := make([]string, g.size*tileSize)

	for y := 0; y < g.size; y++ {
		for x := 0; x < g.size; x++ {
			tile := g.layout[Coordinate{x, y}]
			log.Printf("%s", tile.String())
			for row := 0; row < tileSize; row++ {
				targetRow := y*tileSize + row
				tileContent := tile.content[row+1][1 : tileSize+1]
				content[targetRow] += tileContent
				log.Printf("row %d id %d tileContent '%s' content '%s'", targetRow, tile.id, tileContent, content[targetRow])
			}
		}
	}
	t := NewTile(1, content)

	return t
}

func (g Grid) String() string {
	str := ""
	for _, tile := range g.tiles {
		str = fmt.Sprintf("%s\nid: %d", str, tile.id)
		for edge, n := range tile.neighbors {
			str = fmt.Sprintf("%s %d: %d", str, edge, n.id)
		}
		str = fmt.Sprintf("%s", str)
	}

	return str
}

func findAnswer(lines []string) (int, error) {
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

	g.Assemble()
	tile = g.ToImage()
	log.Printf("image: %d %d", len(tile.content), len(tile.content[0]))

	seamonsters := tile.CountSeaMonsters()
	seamonsterHashes := countChars(seaMonster, "#") * seamonsters
	totalHashes := countChars(tile.content, "#")
	log.Printf("seamonsters: %d hashes: %d, totalhashes: %d", seamonsters, seamonsterHashes, totalHashes)

	return totalHashes - seamonsterHashes, nil
}

func countChars(content []string, char string) int {
	count := 0
	for _, l := range content {
		chars := strings.Split(l, "")

		for _, c := range chars {
			if c == char {
				count++
			}
		}
	}
	return count
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
