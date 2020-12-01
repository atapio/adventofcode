package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func isValid(password int) bool {
	digitsStr := strings.Split(strconv.Itoa(password), "")
	digits := make([]int, len(digitsStr))
	for i, d := range digitsStr {
		n, err := strconv.Atoi(d)
		if err != nil {
			log.Fatalf("error %s", err)
		}
		digits[i] = n
	}

	if containsDecreasing(digits) {
		return false
	}
	if !hasDouble(digits) {
		return false
	}

	return true
}

func containsDecreasing(digits []int) bool {
	prev := -1
	for _, d := range digits {
		if d < prev {
			return true
		}
		prev = d
	}
	return false
}

func hasDouble(digits []int) bool {
	prev := -1
	prevprev := -1
	for i, d := range digits {
		if prevprev != prev && d == prev && (i == len(digits)-1 || digits[i+1] != d) {
			return true
		}
		prevprev = prev
		prev = d
	}
	return false
}

func main() {
	start := 134792
	end := 675810
	count := 0

	for i := start; i <= end; i++ {
		if isValid(i) {
			count++
		}
	}
	fmt.Printf("answer: %d\n", count)
}
