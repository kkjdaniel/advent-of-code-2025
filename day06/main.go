package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readInput("input.txt")

	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}

func part1(lines []string) int {
	numColumns := len(removeEmpty(strings.Split(lines[0], " ")))
	columns := make([][]string, numColumns)

	for _, line := range lines {
		values := removeEmpty(strings.Split(line, " "))
		for i, value := range values {
			columns[i] = append(columns[i], value)
		}
	}

	var totals []int
	for _, column := range columns {
		var acc = 0
		var operator = column[len(column)-1]
		for i := range len(column) - 1 {
			var value, _ = strconv.Atoi(column[i])
			if acc == 0 {
				acc = value
				continue
			}
			switch operator {
			case "*":
				acc = acc * value
			case "/":
				acc = acc / value
			case "-":
				acc = acc - value
			case "+":
				acc = acc + value
			}
		}
		totals = append(totals, acc)
	}

	return sum(totals)
}

func part2(lines []string) int {
	operatorLine := lines[len(lines)-1]
	var columnStarts []int
	for i, ch := range operatorLine {
		if ch == '*' || ch == '+' || ch == '-' || ch == '/' {
			columnStarts = append(columnStarts, i)
		}
	}

	numColumns := len(columnStarts)
	columnEnds := make([]int, numColumns)

	for _, line := range lines[:len(lines)-1] {
		for colIdx, start := range columnStarts {
			end := start
			for end < len(line) && line[end] != ' ' {
				end++
			}
			if end > columnEnds[colIdx] {
				columnEnds[colIdx] = end
			}
		}
	}

	columns := make([][]string, numColumns)

	for _, line := range lines[:len(lines)-1] {
		for colIdx, start := range columnStarts {
			end := columnEnds[colIdx]
			segment := ""
			for i := start; i < end; i++ {
				if line[i] == ' ' {
					segment += "X"
				} else {
					segment += string(line[i])
				}
			}
			columns[colIdx] = append(columns[colIdx], segment)
		}
	}

	for colIdx, start := range columnStarts {
		columns[colIdx] = append(columns[colIdx], string(operatorLine[start]))
	}

	var newColumns [][]string
	for _, column := range columns {
		largestLen := 0
		for i := range len(column) - 1 {
			if len(column[i]) > largestLen {
				largestLen = len(column[i])
			}
		}

		var newColumn []string
		for digitIdx := 0; digitIdx < largestLen; digitIdx++ {
			number := ""
			for rowIdx := range len(column) - 1 {
				number += string(column[rowIdx][digitIdx])
			}
			number = strings.ReplaceAll(number, "X", "")
			if number == "" {
				number = "0"
			}
			newColumn = append(newColumn, number)
		}
		newColumn = append(newColumn, column[len(column)-1])
		newColumns = append(newColumns, newColumn)
	}

	var totals []int
	for _, column := range newColumns {
		var acc = 0
		var operator = column[len(column)-1]
		for i := range len(column) - 1 {
			var value, _ = strconv.Atoi(column[i])
			if acc == 0 {
				acc = value
				continue
			}
			switch operator {
			case "*":
				acc = acc * value
			case "/":
				acc = acc / value
			case "-":
				acc = acc - value
			case "+":
				acc = acc + value
			}
		}
		totals = append(totals, acc)
	}

	return sum(totals)
}

func removeEmpty(values []string) []string {
	var result []string
	for _, value := range values {
		if value != "" {
			result = append(result, value)
		}
	}
	return result
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
