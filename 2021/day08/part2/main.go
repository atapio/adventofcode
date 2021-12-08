package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const MaxCycle = 8
const CycleStart = 6

/*
  Number, segments
  1: 2
  7: 3
  4: 4
  2: 5
  3: 5
  5: 5
  0: 6
  6: 6
  9: 6
  8: 7

  aaaa
 b    c
 b    c
  dddd
 e    f
 e    f
  gggg

  Segments in numbers without 1478:
  a: 6 // all
  b: 4
  c: 4
  d: 5 // 0 is only without
  e: 3
  f: 5 // 2 is only without
  g: 6 // all

  Segments in numbers 3569 / without 012478:
  b: 3 // 3 is only without
  c: 2 // 3,9
  e: 1 // 6 is only with

  Segments in numbers 59
  c: 1 // 9 is only
*/

type Decoder struct {
	digits map[string]int
}

func NewDecoder() *Decoder {
	return &Decoder{
		digits: make(map[string]int),
	}
}

func (d *Decoder) Init(input string) {
	toDecode := []string{}
	nextDecode := []string{}

	for _, segments := range strings.Fields(input) {
		segments := SortString(segments)
		switch len(segments) {
		case 2:
			d.digits[segments] = 1
		case 3:
			d.digits[segments] = 7
		case 4:
			d.digits[segments] = 4
		case 7:
			d.digits[segments] = 8
		default:
			toDecode = append(toDecode, segments)
		}
	}

	segmentCounts := d.countSegments(toDecode)
	fmt.Printf("step1: %v\ncounts: %v\n", d.digits, segmentCounts)

	// we can now find 0 and 2
	zero := ""
	two := ""
	for _, segments := range toDecode {
		switch len(segments) {
		case 5:
			for segment, count := range segmentCounts {
				switch count {
				case 5:
					if !strings.Contains(segments, segment) {
						d.digits[segments] = 2
						two = segments
						break
					}
				}
				nextDecode = append(nextDecode, segments)
			}
		case 6:
			for segment, count := range segmentCounts {
				switch count {
				case 5:
					if !strings.Contains(segments, segment) {
						d.digits[segments] = 0
						zero = segments
						break
					}
				}
				nextDecode = append(nextDecode, segments)
			}
		default:
			nextDecode = append(nextDecode, segments)
		}
	}

	nextDecode = d.unique(nextDecode)
	nextDecode = d.filter(nextDecode, zero)
	nextDecode = d.filter(nextDecode, two)
	toDecode = nextDecode
	nextDecode = []string{}
	segmentCounts = d.countSegments(toDecode)

	fmt.Printf("step2: %v\ncounts: %v\nto decode: %v\n", d.digits, segmentCounts, toDecode)

	// we can now find 3 and 6
	three := ""
	six := ""
	for _, segments := range toDecode {
		switch len(segments) {
		case 1:
			d.digits[segments] = 6
			six = segments
		case 3:
			for segment, count := range segmentCounts {
				switch count {
				case 3:
					if !strings.Contains(segments, segment) {
						d.digits[segments] = 3
						three = segments
						break
					}
				}
				nextDecode = append(nextDecode, segments)
			}
		case 5:
			for segment, count := range segmentCounts {
				switch count {
				case 3:
					if !strings.Contains(segments, segment) {
						d.digits[segments] = 3
						three = segments
						break
					}
				}
				nextDecode = append(nextDecode, segments)
			}
		case 6:
			for segment, count := range segmentCounts {
				switch count {
				case 1:
					if strings.Contains(segments, segment) {
						d.digits[segments] = 6
						six = segments
						break
					}
				}
				nextDecode = append(nextDecode, segments)
			}
		default:
			nextDecode = append(nextDecode, segments)
		}
	}

	nextDecode = d.unique(nextDecode)
	nextDecode = d.filter(nextDecode, three)
	nextDecode = d.filter(nextDecode, six)
	toDecode = nextDecode

	fmt.Printf("step3: %v\nto decode: %v\n", d.digits, toDecode)

	// 5 and 9 left
	for _, segments := range toDecode {
		switch len(segments) {
		case 5:
			d.digits[segments] = 5
		case 6:
			d.digits[segments] = 9
		}
	}
}

func (d *Decoder) countSegments(input []string) map[string]int {
	counts := map[string]int{}
	for _, segments := range input {
		for _, segment := range strings.Split(segments, "") {
			counts[segment]++
		}
	}
	return counts
}

func (d *Decoder) filter(input []string, unwanted string) []string {
	filtered := []string{}
	for _, segments := range input {
		if segments == unwanted {
			continue
		}
		filtered = append(filtered, segments)
	}
	return filtered
}
func (d *Decoder) unique(input []string) []string {
	u := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

func (d *Decoder) Decode(segments string) (int, error) {
	segments = SortString(segments)

	if val, ok := d.digits[segments]; ok {
		return val, nil
	}

	return 0, fmt.Errorf("invalid segments %s", segments)
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func findAnswer(lines []string) (int, error) {
	answer := 0

	for _, line := range lines {
		input := strings.Split(line, "|")
		tests := input[0]
		data := input[1]

		decoder := NewDecoder()
		decoder.Init(tests)
		fmt.Printf("digits: %v\n", decoder.digits)

		value := 0
		for _, segments := range strings.Fields(data) {
			digit, err := decoder.Decode(segments)
			if err != nil {
				return 0, err
			}
			value = value*10 + digit
		}
		answer += value
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
