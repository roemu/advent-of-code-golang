package main

import (
	"fmt"
	"os"

	"aoc/internal/utils"
)

type Equation struct {
	Result int64
	Inputs []int64
	IsPartTwo bool
}

func Calc(a, b int64, operator string) int64 {
	if operator == "*" {
		return a * b
	} else if operator == "|" {
		return int64(utils.Atoi(fmt.Sprintf("%d%d", a, b)))
	} else {
		return a + b
	}
}
// Also known as the cartesion product: https://en.wikipedia.org/wiki/Cartesian_product
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
	operators := []string{"*", "+"}
	if equation.IsPartTwo {
		operators = append(operators, "|")
	}
	combinations := GenerateCombinations(operators, len(equation.Inputs) - 1)
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
		IsPartTwo: false,
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
		eq := NewEquation(line)
		eq.IsPartTwo = false
		return eq
	})
	return SumValidEquations(equations)
}

// part two
func part2(input string) int64 {
	equations := utils.Map(utils.Split(input, "\n"), func(line string) *Equation {
		eq := NewEquation(line)
		eq.IsPartTwo = true
		return eq
	})
	return SumValidEquations(equations)
}
