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

func (l *Line) points() []Point {
	points := []Point{}
	for x := l.minX(); x <= l.maxX(); x++ {
		for y := l.minY(); y <= l.maxY(); y++ {
			points = append(points, Point{x: x, y: y})
		}
	}

	return points
}

func (l *Line) crossing(l2 Line) (Point, bool) {
	for _, p1 := range l.points() {
		for _, p2 := range l2.points() {
			if p1 == p2 {
				return p1, true
			}
		}
	}

	return Point{}, false
}

func (l *Line) Length() int {
	return l.start.distance(l.end)
}

func (l *Line) DistanceTo(p Point) (int, bool) {
	for _, lp := range l.points() {
		if lp == p {
			return l.start.distance(p), true
		}
	}

	return 0, false
}

type Wire struct {
	lines   []Line
	current Point
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
	case 'D':
		line.end.x = line.start.x
		line.end.y = line.start.y - l
	case 'R':
		line.end.x = line.start.x + l
		line.end.y = line.start.y
	case 'L':
		line.end.x = line.start.x - l
		line.end.y = line.start.y
	}

	w.lines = append(w.lines, line)

	w.current = line.end
	//fmt.Printf("%v\n", *w)
}

func (w *Wire) DistanceTo(p Point) (int, bool) {
	distance := 0
	for _, l := range w.lines {
		if d, ok := l.DistanceTo(p); ok {
			distance = distance + d
			return distance, true
		}
		distance = distance + l.Length()
	}

	return 0, false
}

func (w *Wire) FindShortest(w2 *Wire) int {
	mindist := 0

	for _, l := range w.lines {
		for _, l2 := range w2.lines {
			p, found := l.crossing(l2)
			if !found {
				continue
			}
			d1, _ := w.DistanceTo(p)
			d2, _ := w2.DistanceTo(p)
			d := d1 + d2
			if d > 0 && (mindist == 0 || d < mindist) {
				mindist = d
			}
		}
	}

	return mindist
}

func FindAnswer(input []string) int {
	w1 := NewWire(input[0])
	w2 := NewWire(input[1])

	return w1.FindShortest(w2)
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

	answer := FindAnswer(lines)
	fmt.Printf("answer: %d", answer)

}
