package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Passport struct {
	BirthYear      string `validate:"required"`
	IssueYear      string `validate:"required"`
	ExpirationYear string `validate:"required"`
	Height         string `validate:"required"`
	HairColor      string `validate:"required"`
	EyeColor       string `validate:"required"`
	PassportID     string `validate:"required"`
	CountryID      string
}

func (p *Passport) IsValid() bool {
	err := validate.Struct(p)
	if err != nil {
		//fmt.Println(err)
		return false
	}

	return true
}

func NewPassport(data string) *Passport {
	p := &Passport{}

	fields := strings.Split(data, " ")

	for _, field := range fields {
		kv := strings.Split(field, ":")

		switch kv[0] {
		case "byr":
			p.BirthYear = kv[1]
		case "iyr":
			p.IssueYear = kv[1]
		case "eyr":
			p.ExpirationYear = kv[1]
		case "hgt":
			p.Height = kv[1]
		case "hcl":
			p.HairColor = kv[1]
		case "ecl":
			p.EyeColor = kv[1]
		case "pid":
			p.PassportID = kv[1]
		case "cid":
			p.CountryID = kv[1]
		}
	}

	return p
}

type data struct {
	input []*Passport
}

func findAnswer(d data) int {
	valid := 0

	for _, p := range d.input {
		if p.IsValid() {
			valid++
		}
	}

	return valid
}

func processInput(lines []string) data {
	d := data{}

	record := ""

	for _, line := range lines {
		switch line {
		case "":
			d.input = append(d.input, NewPassport(record))
			record = ""
		default:
			record = record + " " + line
		}
	}
	d.input = append(d.input, NewPassport(record))

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
