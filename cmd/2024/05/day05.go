package main

import (
	"fmt"
	"os"

	"aoc/internal/utils"
)

func main() {
	session := os.Getenv("AOC_SESSION")
	input := utils.ReadHTTP(2024, 5, session)

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input))
	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input))

	os.Exit(0)
}

// part one
func part1(input string) int {
	return NewInput(input).SumMiddlePages()
}

// part two
func part2(input string) int {
	return NewInput(input).SumdMiddlePagesCorrect()
}

