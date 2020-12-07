package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var lineFormat = regexp.MustCompile(`^(?P<Color>\w+ \w+) bags contain (?P<Contents>.*).$`)
var contentFormat = regexp.MustCompile(`^ ?(?P<Count>\d+) (?P<Color>\w+ \w+) bags?$`)

type Bags struct {
	Colors map[string]*Bag
}

func NewBags() *Bags {
	b := &Bags{
		Colors: map[string]*Bag{},
	}

	return b
}

func (b *Bags) AddBag(color string) *Bag {
	if bag, ok := b.Colors[color]; ok {
		return bag
	}
	newBag := NewBag(color)
	b.Colors[color] = newBag
	return newBag
}

func (b *Bags) ParseRule(rule string) {
	m := lineFormat.FindStringSubmatch(rule)
	fmt.Printf("%s\n%v\n", rule, m)

	bag := b.AddBag(m[1])
	b.ParseContent(bag, m[2])
}

func (b *Bags) ParseContent(bag *Bag, content string) {
	// contain 1 bright white bag, 2 muted yellow bags.
	// contain 1 shiny gold bag.
	// no other bags

	contents := strings.Split(content, ",")

	if contents[0] == "no other bags" {
		return
	}

	for _, c := range contents {
		m := contentFormat.FindStringSubmatch(c)
		// count := m[1]
		fmt.Printf("%s -> %v\n", c, m)
		contentBag := b.AddBag(m[2])
		bag.Contents = append(bag.Contents, contentBag)
		contentBag.AddOuter(bag)
	}
}

type Bag struct {
	Color    string
	Contents []*Bag
	Outer    map[string]*Bag
}

func NewBag(color string) *Bag {
	return &Bag{
		Color: color,
		Outer: map[string]*Bag{},
	}
}

func (b *Bag) AddOuter(outerBag *Bag) {
	if _, ok := b.Outer[outerBag.Color]; ok {
		return
	}
	b.Outer[outerBag.Color] = outerBag
}

func (b *Bag) OuterColors() []string {
	colors := map[string]bool{}

	for _, o := range b.Outer {
		colors[o.Color] = true
		oc := o.OuterColors()
		fmt.Printf("oc: %v\n", oc)
		for _, c := range oc {
			colors[c] = true
		}
	}
	unique := []string{}
	for k, _ := range colors {
		fmt.Printf("colors %s\n", k)
		unique = append(unique, k)
	}
	return unique
}

type data struct {
	input *Bags
}

func findAnswer(d data) int {
	target := "shiny gold"

	//fmt.Printf("%v", d.input.Colors)
	root := d.input.Colors[target]

	colors := root.OuterColors()

	return len(colors)
}

func processInput(lines []string) data {
	d := data{
		input: NewBags(),
	}

	for _, line := range lines {
		d.input.ParseRule(line)
	}

	return d
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		input = append(input, l)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	d := processInput(input)

	answer := findAnswer(d)

	fmt.Printf("answer: %d\n", answer)
}
