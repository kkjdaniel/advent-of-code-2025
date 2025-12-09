package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	lines := readInput("input.txt")

	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}

func part1(lines []string) int {
	pointer := 50
	zeroCount := 0

	for _, v := range lines {
		movement, _ := strconv.Atoi(v[1:])

		movement = movement % 100

		if string(v[0]) == "L" {
			pointer = pointer - movement
		} else {
			pointer = pointer + movement
		}

		if pointer < 0 {
			pointer = 100 + pointer
		} else if pointer > 99 {
			pointer = pointer - 100
		}

		if pointer == 0 {
			zeroCount++
		}
	}

	return zeroCount
}

func part2(lines []string) int {
	pointer := 50
	zeroCount := 0

	for _, v := range lines {
		rotation, _ := strconv.Atoi(v[1:])

		movement := rotation % 100

		fullRotations := (rotation - movement) / 100
		zeroCount += fullRotations

		startedOnZero := pointer == 0

		if string(v[0]) == "L" {
			pointer = pointer - movement
		} else {
			pointer = pointer + movement
		}

		if pointer <= 0 {
			if !startedOnZero {
				zeroCount++
			}
			if pointer != 0 {
				pointer = 100 + pointer
			}
		} else if pointer > 99 {
			if !startedOnZero {
				zeroCount++
			}
			pointer = pointer - 100
		}
	}

	return zeroCount
}

func readInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
