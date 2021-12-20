package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	x int
	y int
}

func (c Point) Around(count int) []Point {
	x := c.x
	y := c.y

	n := []Point{}

	// order is important
	for j := -count; j <= count; j++ {
		for i := -count; i <= count; i++ {
			n = append(n, Point{x + i, y + j})
		}
	}

	return n
}

type Image struct {
	pixels      map[Point]int
	infinity    int
	topLeft     Point
	bottomRight Point
}

func NewImage() *Image {
	return &Image{pixels: map[Point]int{}}
}

func (i *Image) SetPixel(p Point) {
	i.pixels[p] = 1
	if p.x < i.topLeft.x {
		i.topLeft.x = p.x
	}
	if p.y < i.topLeft.y {
		i.topLeft.y = p.y
	}
	if p.x > i.bottomRight.x {
		i.bottomRight.x = p.x
	}
	if p.y > i.bottomRight.y {
		i.bottomRight.y = p.y
	}
}

func (i *Image) GetPixel(p Point) int {
	if val, ok := i.pixels[p]; ok {
		return val
	}

	if p.x >= i.topLeft.x && p.y >= i.topLeft.y &&
		p.x <= i.bottomRight.x && p.y <= i.bottomRight.y {
		return 0
	}

	return i.infinity
}

func (i *Image) Enhance(enhancement string) *Image {
	newImage := NewImage()
	// 0 or 511
	if enhancement[i.infinity*(int(i.infinity<<9)-1)] == '#' {
		newImage.infinity = 1
	}
	//fmt.Printf("infinity: %d\n", newImage.infinity)

	around := 1
	minX := i.topLeft.x - around
	minY := i.topLeft.y - around
	maxX := i.bottomRight.x + around
	maxY := i.bottomRight.y + around

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			p := Point{x, y}
			v := i.PixelArea(p)
			e := enhancement[v]
			//fmt.Printf("p: %v value: %d enhancement: %c\n", p, v, e)

			if e == '#' {
				newImage.SetPixel(p)
			}
		}
	}

	return newImage
}

func (i *Image) PixelArea(p Point) int {
	area := 0
	for _, pixel := range p.Around(1) {
		//fmt.Printf("%d", i.pixels[pixel])
		area = area << 1

		v := i.GetPixel(pixel)
		area += v
	}
	//fmt.Printf("\n")
	return area
}

func (i *Image) PixelsLit() int {
	count := 0
	for _, v := range i.pixels {
		if v > 0 {
			count++
		}
	}
	return count
}

func findAnswer(lines []string) (int, error) {
	enhancement := lines[0]

	fmt.Printf("enhancement: %c\n", enhancement[34])

	image := &Image{pixels: map[Point]int{}}

	for y, line := range lines[2:] {
		for x, b := range line {
			if b == '#' {
				image.SetPixel(Point{x, y})
			}
		}
	}

	for i := 1; i <= 50; i++ {
		fmt.Printf("round %d\n", i)
		image = image.Enhance(enhancement)
	}

	answer := image.PixelsLit()

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
