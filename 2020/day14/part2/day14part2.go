package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

const BitmaskLength = 36

var maskFormat = regexp.MustCompile(`^mask = (?P<Value>[X01]+)$`)
var instructionFormat = regexp.MustCompile(`^mem\[(?P<Address>\d+)\] = (?P<Value>\d+)$`)

type Computer struct {
	Memory       map[uint64]uint64
	Bitmask      uint64
	FloatingMask []bool
	mutex        *sync.Mutex
}

func MakeComputer() *Computer {
	c := &Computer{
		Memory:       map[uint64]uint64{},
		FloatingMask: make([]bool, BitmaskLength),
		mutex:        &sync.Mutex{},
	}

	return c
}
func (c *Computer) SetMask(bitmaskStr string) {
	bitmask := strings.Split(bitmaskStr, "")
	c.FloatingMask = make([]bool, BitmaskLength)
	c.Bitmask = 0

	for i, op := range bitmask {
		c.FloatingMask[BitmaskLength-i-1] = false

		//fmt.Printf("%d: %s\n", BitmaskLength-i-1, op)

		switch op {
		case "X":
			c.FloatingMask[BitmaskLength-i-1] = true
		case "1":
			c.Bitmask = setBit(c.Bitmask, BitmaskLength-i-1)
		}
	}

	fmt.Printf("bitmask: %v\n", c.Bitmask)
	fmt.Printf("floating mask: %v\n", c.FloatingMask)
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
	modifiedAddr := address
	modifiedAddr |= c.Bitmask

	fmt.Printf("addr %d, modified: %d\n", value, modifiedAddr)

	c.writeFloating(modifiedAddr, value, c.FloatingMask, 0)

}

func (c *Computer) writeFloating(address uint64, value uint64, floatingMask []bool, level int) {
	write := true

	for b, float := range floatingMask {
		if float {
			level++
			write = false
			modifiedMask := make([]bool, BitmaskLength)
			copy(modifiedMask, floatingMask)
			modifiedMask[b] = false

			one := setBit(address, b)
			zero := clearBit(address, b)

			if level < -1 {
				go c.writeFloating(one, value, modifiedMask, level)
				go c.writeFloating(zero, value, modifiedMask, level)
			} else {
				c.writeFloating(one, value, modifiedMask, level)
				c.writeFloating(zero, value, modifiedMask, level)
			}
		}
	}

	if write {
		//		c.mutex.Lock()
		c.Memory[address] = value
		//		c.mutex.Unlock()
	}
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
