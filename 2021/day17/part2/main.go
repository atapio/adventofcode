package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var lineFormat = regexp.MustCompile(`^target area: x=(-?\d+)\.\.(-?\d+), y=(-?\d+)\.\.(-?\d+)$`)

type Point struct {
	x int
	y int
}

type Velocity struct {
	x int
	y int
}

type Projectile struct {
	top int
	pos Point
	v   Velocity
}

func (p *Projectile) Step() Point {
	p.pos.x += p.v.x
	p.pos.y += p.v.y

	p.v.y--
	if p.v.x > 0 {
		p.v.x--
	}

	if p.pos.y > p.top {
		p.top = p.pos.y
	}

	return p.pos
}

type Calculator struct {
	TopLeft     Point
	BottomRight Point
}

func (c *Calculator) Hit(p Projectile) bool {
	return c.hitX(p) && c.hitY(p)
}

func (c *Calculator) hitX(p Projectile) bool {
	return p.pos.x >= c.TopLeft.x &&
		p.pos.x <= c.BottomRight.x
}

func (c *Calculator) hitY(p Projectile) bool {
	return p.pos.y >= c.BottomRight.y &&
		p.pos.y <= c.TopLeft.y
}

func (c *Calculator) MaximizeHeight() int {
	fmt.Printf("target: %v %v\n", c.TopLeft, c.BottomRight)
	minX, _ := c.MinimizeX()
	maxX := Abs(c.BottomRight.x)
	minY := c.BottomRight.y
	maxY := Abs(c.BottomRight.y) // why?

	hits := 0

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			v := Velocity{x: x, y: y}
			p := Projectile{v: v}
			for p.pos.y >= c.BottomRight.y {
				p.Step()
				if c.Hit(p) {
					hits++
					fmt.Printf("hit: %v\n", v)
					break
				}
			}
		}
	}

	return hits
}

func (c *Calculator) MinimizeX() (int, int) {
	v := Velocity{}
	steps := 0
	for x := 1; x < c.TopLeft.x; x++ {
		steps = 0
		p := Projectile{v: Velocity{x: x}}

		for p.v.x > 0 {
			steps++
			p.Step()
		}

		if c.hitX(p) {
			v.x = x
			break
		}
	}
	fmt.Printf("MinimizeX: %d steps: %d\n", v.x, steps)
	return v.x, steps
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findAnswer(lines []string) (int, error) {
	var err error

	answer := 0

	m := lineFormat.FindStringSubmatch(lines[0])
	if len(m) != 5 {
		return 0, fmt.Errorf("failed to parse input '%s'", lines[0])
	}

	c := Calculator{}
	c.TopLeft.x, err = strconv.Atoi(m[1])
	if err != nil {
		return 0, err
	}
	c.BottomRight.x, err = strconv.Atoi(m[2])
	if err != nil {
		return 0, err
	}
	c.BottomRight.y, err = strconv.Atoi(m[3])
	if err != nil {
		return 0, err
	}
	c.TopLeft.y, err = strconv.Atoi(m[4])
	if err != nil {
		return 0, err
	}

	answer = c.MaximizeHeight()

	if answer <= 2926 {
		fmt.Println("too low!")
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
