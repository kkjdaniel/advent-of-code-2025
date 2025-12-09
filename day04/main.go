package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := readInput("input.txt")

	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}

type GridState int

const (
	PaperRoll GridState = iota
	Empty     GridState = iota
)

func part1(lines []string) int {
	grid := buildGrid(lines, nil)

	directions := []struct{ row, col int }{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}

	var accessibleCount = 0
	for rowIdx, row := range grid {
		for itemIdx, item := range row {
			if item == Empty {
				continue
			}
			var itemsSurrounding = 0

			for _, dir := range directions {
				newRow := rowIdx + dir.row
				newCol := itemIdx + dir.col
				if newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(row) {
					if grid[newRow][newCol] != Empty {
						itemsSurrounding++
					}
				}
			}

			if itemsSurrounding < 4 {
				accessibleCount++
			}
		}
	}

	return accessibleCount
}

func part2(lines []string) int {
	var removedIndexes []struct{ row, col int }

	directions := []struct{ row, col int }{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}

	for {
		grid := buildGrid(lines, removedIndexes)
		var rollsToRemove []struct{ row, col int }

		for rowIdx, row := range grid {
			for itemIdx, item := range row {
				if item == Empty {
					continue
				}
				var itemsSurrounding = 0

				for _, dir := range directions {
					newRow := rowIdx + dir.row
					newCol := itemIdx + dir.col
					if newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(row) {
						if grid[newRow][newCol] != Empty {
							itemsSurrounding++
						}
					}
				}

				if itemsSurrounding < 4 {
					rollsToRemove = append(rollsToRemove, struct{ row, col int }{rowIdx, itemIdx})
				}
			}
		}

		if len(rollsToRemove) == 0 {
			break
		}

		removedIndexes = append(removedIndexes, rollsToRemove...)
	}

	return len(removedIndexes)
}

func buildGrid(lines []string, removedIndexes []struct{ row, col int }) [][]GridState {
	var grid [][]GridState

	for rowIdx, line := range lines {
		var items []GridState
		for colIdx, item := range strings.Split(line, "") {
			isRemoved := false
			for _, removed := range removedIndexes {
				if removed.row == rowIdx && removed.col == colIdx {
					isRemoved = true
					break
				}
			}

			if isRemoved {
				items = append(items, Empty)
			} else if item == "@" {
				items = append(items, PaperRoll)
			} else {
				items = append(items, Empty)
			}
		}
		grid = append(grid, items)
	}

	return grid
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
