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

const Mask uint = 65535

var connectionFormat = regexp.MustCompile(`^(.*) -> ([a-z]+)$`)

type Gate interface {
	Signal(left, right uint) (uint, error)
}

type InputGate struct {
	value uint
}

func (i InputGate) Signal(left, right uint) (uint, error) {
	if left != 0 || right != 0 {
		return 0, fmt.Errorf("input gate, no inputs allowed")
	}
	return i.value, nil
}

type XORGate struct {
}

func (_ XORGate) Signal(left, right uint) (uint, error) {
	if left != 0 {
		return 0, fmt.Errorf("unary gate, got %d", left)
	}

	return Mask ^ right, nil
}

type ANDGate struct {
}

func (_ ANDGate) Signal(left, right uint) (uint, error) {
	return left & right, nil
}

type ORGate struct {
}

func (_ ORGate) Signal(left, right uint) (uint, error) {
	return left | right&Mask, nil
}

type LSHIFTGate struct {
}

func (_ LSHIFTGate) Signal(left, right uint) (uint, error) {
	return left << right & Mask, nil
}

type RSHIFTGate struct {
}

func (_ RSHIFTGate) Signal(left, right uint) (uint, error) {
	return left >> right & Mask, nil
}

type Circuit struct {
	wires map[string]uint
}

func NewCircuit() *Circuit {
	return &Circuit{
		wires: map[string]uint{},
	}
}

func (c *Circuit) Process(inputs []string) ([]string, error) {
	remaining := []string{}

	for _, input := range inputs {
		fmt.Printf("wires: %v\n", c.wires)
		fmt.Printf("%s: ", input)
		m := connectionFormat.FindStringSubmatch(input)

		cmdStr, to := m[1], m[2]

		cmd := strings.Fields(cmdStr)
		fmt.Printf("cmd: %v\n", cmd)

		if len(cmd) == 1 {
			l, err := strconv.Atoi(cmd[0])
			left := uint(l)
			if err != nil {
				l, ok := c.wires[cmd[0]]
				if !ok {
					remaining = append(remaining, input)
					continue
				}
				left = l
			}

			o, err := InputGate{value: left}.Signal(0, 0)
			if err != nil {
				return remaining, err
			}
			c.wires[to] = o
			if to == "b" {
				c.wires[to] = 3176
			}
			continue
		}

		switch cmd[0] {
		case "NOT":
			if val, ok := c.wires[cmd[1]]; ok {
				o, err := XORGate{}.Signal(0, val)
				if err != nil {
					return remaining, err
				}
				c.wires[to] = o
			} else {
				remaining = append(remaining, input)
				continue
			}
		default:
			{
				l, err := strconv.Atoi(cmd[0])
				left := uint(l)
				if err != nil {
					l, ok := c.wires[cmd[0]]
					if !ok {
						remaining = append(remaining, input)
						continue
					}
					left = l
				}

				r, err := strconv.Atoi(cmd[2])
				right := uint(r)
				if err != nil {
					r, ok := c.wires[cmd[2]]
					if !ok {
						remaining = append(remaining, input)
						continue
					}
					right = r
				}

				switch cmd[1] {
				case "AND":
					o, err := ANDGate{}.Signal(left, right)
					if err != nil {
						return remaining, err
					}
					c.wires[to] = o
				case "OR":
					o, err := ORGate{}.Signal(left, right)
					if err != nil {
						return remaining, err
					}
					c.wires[to] = o
				case "LSHIFT":
					o, err := LSHIFTGate{}.Signal(left, right)
					if err != nil {
						return remaining, err
					}
					c.wires[to] = o
				case "RSHIFT":
					o, err := RSHIFTGate{}.Signal(left, right)
					if err != nil {
						return remaining, err
					}
					c.wires[to] = o
				default:
					return remaining, fmt.Errorf("unknown command %s", cmdStr)
				}
			}
		}
	}

	return remaining, nil
}

func findAnswer(lines []string) (map[string]uint, error) {
	circuit := NewCircuit()

	prevCount := 0
	loopCount := 0

	remaining, err := circuit.Process(lines)
	if err != nil {
		return map[string]uint{}, err
	}

	for len(remaining) > 0 {
		loopCount++
		fmt.Printf("loop: %d\n", loopCount)
		if len(remaining) == prevCount {
			//fmt.Printf("remaining: %v", remaining)
			return map[string]uint{}, fmt.Errorf("no new lines processed")
		}
		prevCount = len(remaining)
		remaining, err = circuit.Process(remaining)
		if err != nil {
			return map[string]uint{}, err
		}
	}

	return circuit.wires, nil
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

	fmt.Printf("answer: %v\n", answer)
}
