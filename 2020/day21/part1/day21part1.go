package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var foodFormat = regexp.MustCompile(`^([a-z ]+) \(contains ([a-z, ]+)\)$`)

type Allergen struct {
	name                 string
	potentialIngredients map[string]bool
}

func NewAllergen(name string, ingredients []string) *Allergen {
	a := &Allergen{
		potentialIngredients: map[string]bool{},
		name:                 name,
	}

	for _, i := range ingredients {
		a.potentialIngredients[i] = true
	}

	return a
}

func (a *Allergen) AddFood(ingredients []string) {
	for i := range a.potentialIngredients {
		a.potentialIngredients[i] = false
	}
	for _, i := range ingredients {
		if _, ok := a.potentialIngredients[i]; ok {
			a.potentialIngredients[i] = true
		}
	}

	for i, potential := range a.potentialIngredients {
		if !potential {
			delete(a.potentialIngredients, i)
		}
	}

	log.Printf("potential: %s %v", a.name, a.potentialIngredients)
}

func findAnswer(lines []string) (int, error) {
	answer := 0
	allergens := map[string]*Allergen{}
	ingredients := map[string]bool{}
	ingredientCounts := map[string]int{}

	for _, line := range lines {
		m := foodFormat.FindStringSubmatch(line)
		foodIngredients := strings.Split(m[1], " ")
		foodAllergens := strings.Split(m[2], ", ")

		log.Printf("ingredients: %v", foodIngredients)
		log.Printf("allergens: %v", foodAllergens)

		for _, allergen := range foodAllergens {
			if _, ok := allergens[allergen]; !ok {
				allergens[allergen] = NewAllergen(allergen, foodIngredients)
			} else {
				allergens[allergen].AddFood(foodIngredients)
			}
		}

		for _, ingredient := range foodIngredients {
			ingredients[ingredient] = true
			ingredientCounts[ingredient]++
		}
	}

	for _, allergen := range allergens {
		for ingredient, potential := range allergen.potentialIngredients {
			if potential {
				ingredients[ingredient] = false
			}
		}
	}

	for i, safe := range ingredients {
		if safe {
			log.Printf("safe: %s %d", i, ingredientCounts[i])
			answer += ingredientCounts[i]
		}
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
