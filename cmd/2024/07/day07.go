package main

import (
	"fmt"
	"os"

	"aoc/internal/utils"
)

type Equation struct {
	Result int64
	Inputs []int64
}

func Calc(a, b int64, operator string) int64 {
	if operator == "*" {
		return a * b
	} else {
		return a + b
	}
}
func GenerateCombinations(elements []string, length int) [][]string {
	if length == 1 {
		return utils.Map(elements, func(elem string) []string {
			return []string{elem}
		})
	}
	combinations := [][]string{}
	for _, elem := range elements {
		combination := []string{elem}
		followingCombs := GenerateCombinations(elements, length-1)
		for _, next := range followingCombs {
			combinations = append(combinations, append(combination, next...))
		}
	}
	return combinations
}
func (equation *Equation) Valid() bool {
	combinations := GenerateCombinations([]string{"*", "+"}, len(equation.Inputs) - 1)
	for _, combination := range combinations {
		sum := int64(equation.Inputs[0])
		for i := 0; i < len(combination); i++ {
			nextNum := equation.Inputs[i + 1]
			sum = Calc(sum, nextNum, combination[i])
		}
		if sum == equation.Result {
			return true
		}
	}
	return false
}
func SumValidEquations(equations []*Equation) int64 {
	var resultChannels = make([]chan int64, len(equations))
	for i, eq := range equations {
		resultChannels[i] = make(chan int64)
		go func(eq *Equation) {
			if eq.Valid() {
				println(eq.Result, ": valid")
				resultChannels[i] <- eq.Result
			} else {
				resultChannels[i] <- 0
			}
		}(eq)
	}
	var results = make([]int64, len(equations))
	for i := range len(equations) {
		results[i] = <-resultChannels[i]
	}
	return utils.Sum64(results)
}
func NewEquation(input string) *Equation {
	result, nums := utils.SplitLeftRight(input, ": ")
	return &Equation{
		Result: utils.Atoi64(result),
		Inputs: utils.MapAtoi64(utils.Split(nums, " ")),
	}
}

func main() {
	session := os.Getenv("AOC_SESSION")
	input := utils.ReadHTTP(2024, 7, session)

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input))
	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input))

	os.Exit(0)
}

// part one
func part1(input string) int64 {
	equations := utils.Map(utils.Split(input, "\n"), func(line string) *Equation {
		return NewEquation(line)
	})
	return SumValidEquations(equations)
}

// part two
func part2(input string) int64 {
	return 0
}
