package main

import (
	"os"

	"aoc/internal/utils"

	"github.com/charmbracelet/log"
)

func main() {
	session := os.Getenv("AOC_SESSION")
	input := utils.ReadHTTP(2024, XX, session)

	log.Info("--- Part One ---")
	log.Infof("Result %d", part1(input))
	log.Info("--- Part Two ---")
	log.Infof("Result %d", part2(input))

	os.Exit(0)
}

// part one
func part1(input string) int {
	return 0
}

// part two
func part2(input string) int {
	return 0
}

