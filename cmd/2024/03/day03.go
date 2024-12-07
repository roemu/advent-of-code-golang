package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"aoc/internal/utils"
)

type Instruction struct {
	a, b int
}
func (instruction *Instruction) Product() int {
	return instruction.a * instruction.b
}

func NewInstruction(input string) *Instruction {
	val := utils.MapAtoi(strings.Split(input[4:len(input) - 1], ","))
	if len(val) != 2 {
		log.Fatal("Val wasn't of length 2", "val", val, "input", input)
	}
	return &Instruction{
		a: val[0],
		b: val[1],
	}
}

func FindAllMatches(input string) []string {
	reg := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	return reg.FindAllString(input, -1)
}

func FindAllMatchesConditional(input string) []string {
	reg := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\))`)

	enabled := true
	matches := utils.Filter(reg.FindAllString(input, -1), func(match string) bool {
		if match == "don't()" {
			enabled = false
			return false
		} else if match == "do()" {
			enabled = true
			return false
		}
		return enabled
	})

	return matches
}

func main() {
	session := os.Getenv("AOC_SESSION")
	input := utils.ReadHTTP(2024, 3, session)
	reg := regexp.MustCompile(`/mul\(\d{1,3},\d{1,3}\)/gm`)
	fmt.Println(reg.FindAllString(input, -1))

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input))
	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input))

	os.Exit(0)
}

// part one
func part1(input string) int {
	matches := FindAllMatches(input)

	instructions := utils.Map(matches, func(match string) *Instruction {
		return NewInstruction(match)
	})

	return utils.Reduce(instructions, func(acc int, instruction *Instruction) int {
		return acc + instruction.Product()
	}, 0)
}

// part two
func part2(input string) int {
	matches := FindAllMatchesConditional(input)

	instructions := utils.Map(matches, func(match string) *Instruction {
		return NewInstruction(match)
	})

	return utils.Reduce(instructions, func(acc int, instruction *Instruction) int {
		return acc + instruction.Product()
	}, 0)
}

