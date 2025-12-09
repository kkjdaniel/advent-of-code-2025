package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ranges := readInput("input.txt")

	fmt.Println("Part 1:", part1(ranges))
	fmt.Println("Part 2:", part2(ranges))
}

func part1(ranges []string) int {
	var invalidIds []int

	for _, v := range ranges {
		currValueInt, _ := strconv.Atoi(strings.Split(v, "-")[0])
		rangeEnd, _ := strconv.Atoi(strings.Split(v, "-")[1])

		for currValueInt <= rangeEnd {
			currValueStr := strconv.Itoa(currValueInt)
			if currValueInt > 9 && len(currValueStr)%2 == 0 {
				seq := currValueStr[0:(len(currValueStr) / 2)]
				if strings.Count(currValueStr, seq) == 2 {
					invalidIds = append(invalidIds, currValueInt)
				}
			}
			currValueInt++
		}
	}

	return sum(invalidIds)
}

func part2(ranges []string) int {
	var invalidIds []int

	for _, v := range ranges {
		currValueInt, _ := strconv.Atoi(strings.Split(v, "-")[0])
		rangeEnd, _ := strconv.Atoi(strings.Split(v, "-")[1])

		for currValueInt <= rangeEnd {
			currValueStr := strconv.Itoa(currValueInt)

			if currValueInt > 9 {
				for seqLen := 1; seqLen <= len(currValueStr)/2; seqLen++ {
					if len(currValueStr)%seqLen == 0 {
						seq := currValueStr[0:seqLen]
						if strings.Count(currValueStr, seq) == len(currValueStr)/seqLen {
							invalidIds = append(invalidIds, currValueInt)
							break
						}
					}
				}
			}

			currValueInt++
		}
	}

	return sum(invalidIds)
}

func sum(nums []int) int {
	total := 0
	for _, x := range nums {
		total += x
	}
	return total
}

func readInput(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), ",")
}
