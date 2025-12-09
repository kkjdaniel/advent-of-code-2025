package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ranges, ids := readInput("input.txt")

	fmt.Println("Part 1:", part1(ranges, ids))
	fmt.Println("Part 2:", part2(ranges))
}

func part1(ranges []string, ids []string) int {
	count := 0

	for _, id := range ids {
		idInt, _ := strconv.Atoi(id)

		for _, r := range ranges {
			parts := strings.Split(r, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])

			if idInt >= start && idInt <= end {
				count++
				break
			}
		}
	}

	return count
}

func part2(ranges []string) int {
	type rangeVal struct {
		start int
		end   int
	}

	var parsedRanges []rangeVal
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		end, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		parsedRanges = append(parsedRanges, rangeVal{start, end})
	}

	var dedupedRanges []rangeVal
	for _, r := range parsedRanges {
		start := r.start
		end := r.end

		var newDeduped []rangeVal
		for _, other := range dedupedRanges {
			if start <= other.start && end >= other.end {
				continue
			}
			newDeduped = append(newDeduped, other)
		}
		dedupedRanges = newDeduped

		for _, other := range dedupedRanges {
			if start >= other.start && start <= other.end {
				start = other.end + 1
			}
			if end >= other.start && end <= other.end {
				end = other.start - 1
			}
		}

		if start > end {
			continue
		}

		dedupedRanges = append(dedupedRanges, rangeVal{start, end})
	}

	total := 0
	for _, r := range dedupedRanges {
		total += r.end - r.start + 1
	}

	return total
}

func readInput(filename string) ([]string, []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var ranges []string
	var ids []string
	isRanges := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isRanges = false
			continue
		}
		if isRanges {
			ranges = append(ranges, line)
		} else {
			ids = append(ids, line)
		}
	}

	return ranges, ids
}
