package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var maskFormat = regexp.MustCompile(`^mask = (?P<Value>[X01]+)$`)
var instructionFormat = regexp.MustCompile(`^mem\[(?P<Address>\d+)\] = (?P<Value>\d+)$`)

type Computer struct {
	Memory  map[uint64]uint64
	Bitmask map[int]bool
}

func MakeComputer() *Computer {
	c := &Computer{
		Memory:  map[uint64]uint64{},
		Bitmask: map[int]bool{},
	}

	return c
}
func (c *Computer) SetMask(bitmaskStr string) {
	c.Bitmask = map[int]bool{}
	bitmask := strings.Split(bitmaskStr, "")

	for i, op := range bitmask {
		//fmt.Printf("bit: %d %s\n", 35-i, op)
		switch op {
		case "1":
			c.Bitmask[35-i] = true
		case "0":
			c.Bitmask[35-i] = false
		}
	}

	fmt.Printf("bitmask: %v\n", c.Bitmask)

}

// Sets the bit at pos in the integer n.
func setBit(n uint64, pos int) uint64 {
	n |= (1 << pos)
	return n
}

// Clears the bit at pos in n.
func clearBit(n uint64, pos int) uint64 {
	mask := ^(uint64(1 << pos))
	n &= mask
	return n
}

func (c *Computer) Write(address uint64, value uint64) {
	modifiedVal := value

	for b, op := range c.Bitmask {
		switch op {
		case true:
			modifiedVal = setBit(modifiedVal, b)
		case false:
			modifiedVal = clearBit(modifiedVal, b)
		}
	}
	fmt.Printf("val %d, written: %d\n", value, modifiedVal)

	c.Memory[address] = modifiedVal
}

func (c *Computer) CountMemory() uint64 {
	count := uint64(0)

	for _, m := range c.Memory {
		count += m
	}
	return count
}

func findAnswer(lines []string) (int, error) {
	c := MakeComputer()

	for _, i := range lines {
		switch i[0:4] {
		case "mask":
			m := maskFormat.FindStringSubmatch(i)
			c.SetMask(m[1])
		case "mem[":
			m := instructionFormat.FindStringSubmatch(i)
			addr, err := strconv.ParseUint(m[1], 10, 64)
			if err != nil {
				return 0, err
			}
			val, err := strconv.ParseUint(m[2], 10, 64)
			if err != nil {
				return 0, err
			}

			c.Write(addr, val)
		}
	}
	return int(c.CountMemory()), nil
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

	answer, err := findAnswer(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("answer: %d\n", answer)
}
