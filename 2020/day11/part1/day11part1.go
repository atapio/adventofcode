package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type WaitingArea struct {
	Layout [][]string
}

func NewWaitingArea(lines []string) *WaitingArea {
	wa := &WaitingArea{}

	for _, line := range lines {
		row := strings.Split(line, "")
		wa.Layout = append(wa.Layout, row)
	}

	return wa
}

func (w *WaitingArea) OccupiedSeats() int {
	count := 0

	for _, row := range w.Layout {
		for _, col := range row {
			if col == "#" {
				count++
			}
		}
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (w *WaitingArea) OccupiedSeatsAround(row int, col int) int {
	count := 0

	minRow := max(row-1, 0)
	minCol := max(col-1, 0)
	maxRow := min(row+1, len(w.Layout)-1)
	maxCol := min(col+1, len(w.Layout[0])-1)

	for y := minRow; y <= maxRow; y++ {
		for x := minCol; x <= maxCol; x++ {
			if x == col && y == row {
				continue
			}
			if w.Layout[y][x] == "#" {
				count++
			}
		}
	}

	//fmt.Printf("r: %d c: %d co: %d\n", row, col, count)
	return count
}

func (w *WaitingArea) String() string {
	lines := []string{}
	for _, row := range w.Layout {
		lines = append(lines, strings.Join(row, ""))
	}

	return strings.Join(lines, "\n")
}

func (w *WaitingArea) NextRound() bool {
	next := [][]string{}
	changed := false

	for y, row := range w.Layout {
		nextrow := []string{}
		for x, pos := range row {
			nextrow = append(nextrow, w.Layout[y][x])

			switch pos {
			case "L":
				if w.OccupiedSeatsAround(y, x) == 0 {
					nextrow[x] = "#"
					changed = true
				}
			case "#":
				if w.OccupiedSeatsAround(y, x) >= 4 {
					nextrow[x] = "L"
					changed = true
				}
			}
		}
		next = append(next, nextrow)
	}

	w.Layout = next
	return changed
}

type data struct {
	input *WaitingArea
}

func findAnswer(d data) (int, error) {
	for {
		fmt.Println("next round")
		fmt.Printf("%s\n", d.input.String())
		if !d.input.NextRound() {
			break
		}
	}
	return d.input.OccupiedSeats(), nil
}

func processInput(lines []string) (data, error) {
	d := data{}

	wa := NewWaitingArea(lines)
	d.input = wa

	return d, nil
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

	d, err := processInput(input)
	if err != nil {
		log.Fatal(err)
	}

	answer, err := findAnswer(d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("answer: %d\n", answer)
}
