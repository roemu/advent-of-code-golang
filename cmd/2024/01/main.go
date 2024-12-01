package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"aoc/internal/utils"
)

func main() {
	session := os.Getenv("AOC_SESSION")
	input, err := utils.ReadHTTP(2024, 1, session)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input))
	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input))

	os.Exit(0)
}

// part one
func part1(input string) int {
	split := strings.Split(input, "\n")
	split = utils.Filter(split, func(line string) bool {
		return line != ""
	})
	left := utils.Map(split, func(line string, index int) int {
		res, err := strconv.ParseInt(strings.Split(line, "   ")[0], 10, 32)
		if err != nil {
			println("Line: ", line)
			panic("Couldn't convert string to number")
		}
		return int(res)
	})
	right := utils.Map(split, func(line string, index int) int {
		res, err := strconv.ParseInt(strings.Split(line, "   ")[1], 10, 32)
		if err != nil {
			println("Line: ", line)
			panic("Couldn't convert string to number")
		}
		return int(res)
	})
	slices.Sort(left)
	slices.Sort(right)
	if len(left) != len(right) {
		panic("slices were not same size")
	}
	distances := utils.Map(right, func(val int, index int) int {
		return int(math.Abs(float64(val-left[index])))
	})
	return utils.Reduce(distances, func(acc int, elem int) int {
		return acc + elem
	}, 0)
}

// part two
func part2(input string) int {
	split := strings.Split(input, "\n")
	split = utils.Filter(split, func(line string) bool {
		return line != ""
	})
	left := utils.Map(split, func(line string, index int) int {
		res, err := strconv.ParseInt(strings.Split(line, "   ")[0], 10, 32)
		if err != nil {
			println("Line: ", line)
			panic("Couldn't convert string to number")
		}
		return int(res)
	})
	right := utils.Map(split, func(line string, index int) int {
		res, err := strconv.ParseInt(strings.Split(line, "   ")[1], 10, 32)
		if err != nil {
			println("Line: ", line)
			panic("Couldn't convert string to number")
		}
		return int(res)
	})
	similarity := utils.Map(left, func(val, index int) int {
		occurences := len(utils.Filter(right, func(r int) bool {
			return r == val
		}))
		return val * occurences
	})

	return utils.Reduce(similarity, func(acc int, elem int) int {
		return acc + elem
	}, 0)
}
