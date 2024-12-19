package main

import (
	"os"

	"aoc/internal/utils"

	"github.com/charmbracelet/log"
)

var memory Memoized

func (m *Memoized) Call(design string, patterns []string) int {
	if _, ok := m.Cache[design]; !ok {
		m.Cache[design] = m.Func(design, patterns)
	}
	return m.Cache[design]
}

type Memoized struct {
	Func  func(string, []string) int
	Cache map[string]int
}

func main() {
	session := os.Getenv("AOC_SESSION")
	input := utils.ReadHTTP(2024, 19, session)

	log.Info("--- Part One ---")
	log.Infof("Result %d", part1(input))
	log.Info("--- Part Two ---")
	log.Infof("Result %d", part2(input))

	os.Exit(0)
}

func MatchDesignAny(design string, patterns []string) int {
	possible := 0
	for _, pattern := range patterns {
		pLen := len(pattern)
		if len(design) < pLen {
			continue
		}
		if design[:pLen] == pattern {
			if len(design[pLen:]) == 0 {
				possible++
			} else {
				possible += memory.Call(design[pLen:], patterns)
			}
		}
	}

	return possible
}

func MatchDesign(design string, patterns []string) bool {
	for _, pattern := range patterns {
		pLen := len(pattern)
		if len(design) < pLen {
			continue
		}
		if design[:pLen] == pattern {
			if len(design[pLen:]) == 0 {
				return true
			}
			if MatchDesign(design[pLen:], patterns) {
				return true
			}
		}
	}

	return false
}

func ParseInput(input string) (patterns []string, designs []string) {
	p, d := utils.SplitLeftRight(input, "\n\n")
	patterns = utils.Split(p, ", ")
	designs = utils.Split(d, "\n")

	return
}

// part one
func part1(input string) int {
	patterns, designs := ParseInput(input)
	return utils.Reduce(designs, func(acc int, design string) int {
		if MatchDesign(design, patterns) {
			return acc + 1
		}

		return acc
	}, 0)
}

// part two
func part2(input string) int {
	memory = Memoized{
		Func:  MatchDesignAny,
		Cache: map[string]int{},
	}
	patterns, designs := ParseInput(input)
	return utils.Reduce(designs, func(acc int, design string) int {
		return acc + memory.Call(design, patterns)
	}, 0)
}
