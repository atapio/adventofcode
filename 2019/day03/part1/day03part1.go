package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func (p *Point) distance(p2 Point) int {
	return Abs(p.x-p2.x) + Abs(p.y-p2.y)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Line struct {
	start Point
	end   Point
}

func (l *Line) minX() int {
	return min(l.start.x, l.end.x)
}
func (l *Line) maxX() int {
	return max(l.start.x, l.end.x)
}

func (l *Line) minY() int {
	return min(l.start.y, l.end.y)
}

func (l *Line) maxY() int {
	return max(l.start.y, l.end.y)
}

func (l *Line) horizontal() bool {
	return l.start.y == l.end.y
}

func (l *Line) vertical() bool {
	return l.start.x == l.end.x
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func (l *Line) crossing(l2 Line) (Point, bool) {
	if l.vertical() {
		return l2.crossing(*l)
	}

	if l.start.y < l2.minY() {
		return Point{}, false
	}
	if l.start.y > l2.maxY() {
		return Point{}, false
	}
	if l2.start.x < l.minX() {
		return Point{}, false
	}
	if l2.start.x > l.maxX() {
		return Point{}, false
	}

	return Point{x: l2.start.x, y: l.end.y}, true
}

type Wire struct {
	horizontal []Line
	vertical   []Line
	current    Point
}

func NewWire(input string) *Wire {
	w := &Wire{}
	path := strings.Split(input, ",")

	for _, p := range path {
		w.AddLine(p)
	}

	return w
}

func (w *Wire) AddLine(path string) {
	line := Line{start: w.current}

	op := path[0]
	l, err := strconv.Atoi(path[1:])
	if err != nil {
		log.Fatal(err)
	}

	switch op {
	case 'U':
		line.end.x = line.start.x
		line.end.y = line.start.y + l
		w.vertical = append(w.vertical, line)
	case 'D':
		line.end.x = line.start.x
		line.end.y = line.start.y - l
		w.vertical = append(w.vertical, line)
	case 'R':
		line.end.x = line.start.x + l
		line.end.y = line.start.y
		w.horizontal = append(w.horizontal, line)
	case 'L':
		line.end.x = line.start.x - l
		line.end.y = line.start.y
		w.horizontal = append(w.horizontal, line)
	}

	w.current = line.end
	//fmt.Printf("%v\n", *w)
}

func (w *Wire) findClosest(w2 *Wire) int {
	mindist := 0
	zero := Point{}

	for _, l := range w.horizontal {
		for _, l2 := range w2.vertical {
			p, found := l.crossing(l2)
			if !found {
				continue
			}
			d := zero.distance(p)
			fmt.Printf("Found crossing min: %d d: %d p: %v\n", mindist, d, p)
			if d > 0 && (mindist == 0 || d < mindist) {
				mindist = d
			}
		}
	}

	for _, l := range w.vertical {
		for _, l2 := range w2.horizontal {
			p, found := l.crossing(l2)
			if !found {
				continue
			}
			d := zero.distance(p)
			fmt.Printf("Found crossing min: %d d: %d p: %v\n", mindist, d, p)
			if d > 0 && (mindist == 0 || d < mindist) {
				mindist = d
			}
		}
	}
	return mindist
}

func findClosest(input []string) int {
	w1 := NewWire(input[0])
	w2 := NewWire(input[1])

	return w1.findClosest(w2)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	answer := findClosest(lines)
	fmt.Printf("answer: %d", answer)

}
