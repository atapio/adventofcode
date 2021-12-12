package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var bigCaveFormat = regexp.MustCompile(`^[A-Z]+$`)
var smallCaveFormat = regexp.MustCompile(`^[a-z]+$`)

type Cave struct {
	name  string
	big   bool
	paths map[string]bool
}

func NewCave() *Cave {
	return &Cave{
		paths: map[string]bool{},
	}
}

func (c *Cave) SetName(name string) {
	c.name = name
	c.big = bigCaveFormat.MatchString(c.name)
}

func (c *Cave) AddPath(name string) {
	c.paths[name] = true
}

type CaveMap struct {
	caves map[string]*Cave
}

func NewCaveMap() *CaveMap {
	return &CaveMap{
		caves: map[string]*Cave{},
	}
}

func (c *CaveMap) AddPath(line string) error {
	path := strings.Split(line, "-")
	if len(path) != 2 {
		return fmt.Errorf("invalid path: %s", line)
	}

	from := c.FindOrAddCave(path[0])
	to := c.FindOrAddCave(path[1])
	from.AddPath(to.name)
	to.AddPath(from.name)

	return nil
}

func (c *CaveMap) FindOrAddCave(name string) *Cave {
	if cave, ok := c.caves[name]; ok {
		return cave
	}

	cave := NewCave()
	cave.name = name

	c.caves[name] = cave
	return cave
}

func (c *CaveMap) FindAllPaths(start string, end string, path []string) [][]string {
	fmt.Printf("FindAllPaths: %s -> %s %v\n", start, end, path)
	path = append(path, start)

	if start == end {
		return [][]string{path}
	}

	paths := [][]string{}

	for cave := range c.caves[start].paths {
		visited := false

		for _, p := range path {
			if cave == p && smallCaveFormat.MatchString(p) {
				visited = true
				break
			}
		}

		if visited {
			continue
		}

		newPaths := c.FindAllPaths(cave, end, path)
		paths = append(paths, newPaths...)
	}

	return paths
}

func (c *CaveMap) CountPaths() int {
	paths := c.FindAllPaths("start", "end", []string{})
	return len(paths)
}

func findAnswer(lines []string) (int, error) {
	answer := 0

	cm := NewCaveMap()
	for _, line := range lines {
		err := cm.AddPath(line)
		if err != nil {
			return 0, err
		}
	}

	answer = cm.CountPaths()

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
