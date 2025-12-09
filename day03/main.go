package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	banks := readInput("input.txt")

	fmt.Println("Part 1:", part1(banks))
	fmt.Println("Part 2:", part2(banks))
}

func part1(banks []string) int {
	var joltage []int

	for _, bank := range banks {
		best := 0
		for i := 0; i < len(bank); i++ {
			for j := i + 1; j < len(bank); j++ {
				combined, _ := strconv.Atoi(string(bank[i]) + string(bank[j]))
				if combined > best {
					best = combined
				}
			}
		}

		joltage = append(joltage, best)
	}

	return sum(joltage)
}

func part2(banks []string) int {
	var joltage []int

	for _, bank := range banks {
		var digits []string

		firstDigit, cursor := findHighestNumber(bank[:len(bank)-11])
		digits = append(digits, firstDigit)

		for len(digits) < 12 {
			remaining := bank[cursor+1:]
			digitsNeeded := 12 - len(digits)

			if len(remaining) == digitsNeeded {
				for _, c := range remaining {
					digits = append(digits, string(c))
				}
				break
			}

			candidates := bank[cursor+1 : len(bank)-digitsNeeded+1]
			nextDigit, idx := findHighestNumber(candidates)
			digits = append(digits, nextDigit)
			cursor += idx + 1
		}

		result, _ := strconv.Atoi(strings.Join(digits, ""))
		joltage = append(joltage, result)
	}

	return sum(joltage)
}

func findHighestNumber(s string) (string, int) {
	best := 0
	bestStr := ""
	bestIdx := 0
	for i := range len(s) {
		value, _ := strconv.Atoi(string(s[i]))
		if value > best {
			best = value
			bestStr = string(s[i])
			bestIdx = i
		}
	}
	return bestStr, bestIdx
}

func sum(nums []int) int {
	total := 0
	for _, x := range nums {
		total += x
	}
	return total
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
