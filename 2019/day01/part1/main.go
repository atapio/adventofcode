package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func fuelNeeded(mass int) int {
	return mass/3 - 2
}

func processLine(line string, current int) (int, error) {
	mass, err := strconv.Atoi(line)
	if err != nil {
		return 0, err
	}

	return current + fuelNeeded(mass), nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fuel := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fuel, err = processLine(scanner.Text(), fuel)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result: %d\n", fuel)
}
