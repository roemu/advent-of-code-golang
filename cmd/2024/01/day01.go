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
	input := utils.ReadHTTP(2024, 1, session)

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input))
	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input))

	os.Exit(0)
}

type LeftRight struct {
	Left []int
	Right []int
}

// part one
func part1(input string) int {
	split := strings.Split(input, "\n")
	split = utils.Filter(split, func(line string) bool {
		return line != ""
	})
	leftright := utils.Reduce(split, func(acc LeftRight, line string) LeftRight {
		left, _ := strconv.ParseInt(strings.Split(line, "   ")[0], 10, 32)
		right, _ := strconv.ParseInt(strings.Split(line, "   ")[1], 10, 32)
		acc.Left = append(acc.Left, int(left))
		acc.Right = append(acc.Right, int(right))
		return acc
	}, LeftRight{})
	slices.Sort(leftright.Left)
	slices.Sort(leftright.Right)
	distances := utils.MapI(leftright.Left, func(val int, index int) int {
		return int(math.Abs(float64(val-leftright.Right[index])))
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
	leftright := utils.Reduce(split, func(acc LeftRight, line string) LeftRight {
		left, _ := strconv.ParseInt(strings.Split(line, "   ")[0], 10, 32)
		right, _ := strconv.ParseInt(strings.Split(line, "   ")[1], 10, 32)
		acc.Left = append(acc.Left, int(left))
		acc.Right = append(acc.Right, int(right))
		return acc
	}, LeftRight{})
	similarity := utils.Map(leftright.Left, func(val int) int {
		occurences := len(utils.Filter(leftright.Right, func(r int) bool {
			return r == val
		}))
		return val * occurences
	})

	return utils.Reduce(similarity, func(acc int, elem int) int {
		return acc + elem
	}, 0)
}
