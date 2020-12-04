package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func init() {
	validate.RegisterValidation("height", ValidateHeight)
}

func ValidateHeight(fl validator.FieldLevel) bool {
	f := fl.Field().String()
	split := len(f) - 2

	h, err := strconv.Atoi(f[:split])
	if err != nil {
		return false
	}

	unit := f[split:]
	switch unit {
	case "cm":
		return h >= 150 && h <= 193
	case "in":
		return h >= 59 && h <= 76
	}
	return false
}

type Passport struct {
	BirthYear      int    `validate:"required,gte=1920,lte=2002"`
	IssueYear      int    `validate:"required,gte=2010,lte=2020"`
	ExpirationYear int    `validate:"required,gte=2020,lte=2030"`
	Height         string `validate:"required,height"`
	HairColor      string `validate:"required,hexcolor"`
	EyeColor       string `validate:"required,oneof=amb blu brn gry grn hzl oth"`
	PassportID     string `validate:"required,numeric,len=9"`
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
			v, _ := strconv.Atoi(kv[1])
			p.BirthYear = v
		case "iyr":
			v, _ := strconv.Atoi(kv[1])
			p.IssueYear = v
		case "eyr":
			v, _ := strconv.Atoi(kv[1])
			p.ExpirationYear = v
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
