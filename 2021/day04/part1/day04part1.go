package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const BoardSize = 5

type Bingo struct {
	boards  []*Board
	numbers []int
	size    int
}

func NewBingo() *Bingo {
	return &Bingo{
		boards: make([]*Board, 0, 10),
	}
}

func (b *Bingo) ParseNumbers(input string) error {
	numbers := strings.Split(input, ",")
	b.numbers = make([]int, len(numbers))
	for i, n := range numbers {
		num, err := strconv.Atoi(n)
		if err != nil {
			return err
		}
		b.numbers[i] = num
	}
	return nil
}

func (b *Bingo) Play() (*Board, int) {
	for _, n := range b.numbers {
		for _, board := range b.boards {
			if board.Mark(n) {
				return board, n
			}
			fmt.Printf("%s\n", board.String())
		}

	}
	return nil, -1
}

func (b *Bingo) AddBoard(input []string) error {
	if len(input) != BoardSize {
		return fmt.Errorf("invalid input length %d", len(input))
	}

	board := Board{
		number: len(b.boards),
		rows:   make([][]int, 2*BoardSize),
	}
	b.boards = append(b.boards, &board)

	for i, l := range input {
		row := strings.Fields(l)
		board.rows[i] = make([]int, BoardSize)
		// map colunms to rows

		fmt.Printf("row: %s\n", l)

		for j, n := range row {
			// Horrible kludge
			if i == 0 {
				board.rows[j+BoardSize] = make([]int, BoardSize)
			}
			num, err := strconv.Atoi(n)
			if err != nil {
				return err
			}

			board.rows[i][j] = num
			board.rows[j+BoardSize][i] = num
		}
	}

	return nil
}

type Board struct {
	number int
	rows   [][]int
}

func (b *Board) Mark(number int) bool {
	bingo := false
	for r, row := range b.rows {
		for i, n := range row {
			if n == number {
				b.rows[r] = append(b.rows[r][:i], b.rows[r][i+1:]...)
				fmt.Printf("b: %d match %d, row: %v\n", b.number, n, b.rows[r])
				if len(b.rows[r]) == 0 {
					bingo = true
				}
			}
		}
	}
	return bingo
}

func (b *Board) Unmarked() []int {
	unmarked := []int{}
	for i := 0; i < BoardSize; i++ {
		for _, n := range b.rows[i] {
			unmarked = append(unmarked, n)
		}
	}
	return unmarked
}

func (b *Board) String() string {
	s := fmt.Sprintf("board: %d", b.number)
	for i := range b.rows {
		s = fmt.Sprintf("%s\n%v", s, b.rows[i])
	}
	return s
}

func findAnswer(lines []string) (int, error) {
	answer := 0

	bingo := Bingo{size: 5}

	numbers, lines := lines[0], lines[1:]
	bingo.ParseNumbers(numbers)

	for len(lines) >= BoardSize {
		if lines[0] == "" {
			lines = lines[1:]
			continue
		}
		var board []string

		board, lines = lines[0:5], lines[5:]
		fmt.Printf("board str %v\n", board)
		err := bingo.AddBoard(board)
		if err != nil {
			return 0, err
		}
	}
	fmt.Printf("boards: %d lines left: %d\n", len(bingo.boards), len(lines))

	board, number := bingo.Play()

	fmt.Printf("board: %d number: %d\n", board.number, number)

	for _, n := range board.Unmarked() {
		fmt.Printf("%d ", n)
		answer += n
	}

	answer *= number

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
