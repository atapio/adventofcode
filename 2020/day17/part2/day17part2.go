package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Coordinate struct {
	v int
	x int
	y int
	z int
}

func (c Coordinate) Neighbors() []Coordinate {
	n := []Coordinate{}

	for v := c.v - 1; v <= c.v+1; v++ {
		for x := c.x - 1; x <= c.x+1; x++ {
			for y := c.y - 1; y <= c.y+1; y++ {
				for z := c.z - 1; z <= c.z+1; z++ {
					coord := Coordinate{v, x, y, z}
					if c == coord {
						continue
					}
					n = append(n, coord)

				}
			}
		}
	}

	return n
}

type Grid struct {
	Cubes map[Coordinate]bool
}

func MakeGrid(initial []string) *Grid {
	g := &Grid{
		Cubes: map[Coordinate]bool{},
	}

	for y, line := range initial {
		for x, v := range strings.Split(line, "") {
			if v == "#" {
				g.Cubes[Coordinate{0, x, y, 0}] = true
			}
		}
	}

	return g
}

func (g *Grid) BoundingBox() (Coordinate, Coordinate) {
	min := Coordinate{}
	max := Coordinate{}

	for coord, active := range g.Cubes {
		if !active {
			continue
		}
		if coord.v < min.v {
			min.v = coord.v
		}
		if coord.x < min.x {
			min.x = coord.x
		}
		if coord.y < min.y {
			min.y = coord.y
		}
		if coord.z < min.z {
			min.z = coord.z
		}
		if coord.v > max.v {
			max.v = coord.v
		}
		if coord.x > max.x {
			max.x = coord.x
		}
		if coord.y > max.y {
			max.y = coord.y
		}
		if coord.z > max.z {
			max.z = coord.z
		}
	}

	min.v--
	min.x--
	min.y--
	min.z--
	max.v++
	max.x++
	max.y++
	max.z++

	return min, max
}

func (g *Grid) SimulateCycle() {
	next := map[Coordinate]bool{}
	min, max := g.BoundingBox()

	fmt.Printf("min: %v max: %v\n", min, max)

	for v := min.v; v <= max.v; v++ {
		for x := min.x; x <= max.x; x++ {
			for y := min.y; y <= max.y; y++ {
				for z := min.z; z <= max.z; z++ {
					coord := Coordinate{v, x, y, z}
					count := g.CountNeigbours(coord)
					fmt.Printf("coord: %v/%t count: %d\n", coord, g.Cubes[coord], count)
					switch g.Cubes[coord] {
					case true:
						if count == 2 || count == 3 {
							next[coord] = true
						}
					case false:
						if count == 3 {
							next[coord] = true
						}
					}
				}
			}
		}
	}
	/*
		for coord, isActive := range g.Cubes {
			if isActive {
				count := g.CountNeigbours(coord)
				//fmt.Printf("active count %v %d\n", coord, count)
				next[coord] = (count == 2 || count == 3)
				for _, neigh := range coord.Neighbors() {
					if !g.Cubes[neigh] {
						count := g.CountNeigbours(neigh)
						//fmt.Printf("n count %v %d\n", neigh, count)
						next[neigh] = count == 3
					}
				}
			} else {
				next[coord] = g.CountNeigbours(coord) == 3
			}

		}
	*/
	fmt.Printf("cubesb: %v\n", g.Cubes)
	fmt.Printf("cubesa: %v\n", next)
	g.Cubes = next
}

func (g Grid) CountNeigbours(c Coordinate) int {
	count := 0
	for _, coord := range c.Neighbors() {
		if g.Cubes[coord] {
			count++
		}
	}

	return count
}

func (g Grid) CountActive() int {
	count := 0
	for _, active := range g.Cubes {
		if active {
			count++
		}
	}

	return count
}

func findAnswer(lines []string) (int, error) {
	// fields
	g := MakeGrid(lines)
	for i := 0; i < 6; i++ {
		g.SimulateCycle()
		fmt.Printf("active: %d\n", g.CountActive())
	}

	answer := g.CountActive()

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
	// remove last empty line if it exists
	/*
		if input[len(input)-1] == "" {
			input = input[:len(input)-2]
		}
	*/
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
