package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"

	"aoc/internal/utils"

	"github.com/charmbracelet/log"
)

type Report struct {
	Raw     string
	Numbers []int
}

func NewReport(input string) *Report {
	return &Report{
		Raw: input,
		Numbers: utils.MapAtoi(utils.FilterEmpty(strings.Split(input, " "))),
	}
}
func (report *Report) IsSafe() bool {
	if !IsConsecutive(report.Numbers) {
		return false
	}
	return SafeDistances(report.Numbers)
}
func (report *Report) IsSafeIgnoreOne() bool {
	if SafeDistances(report.Numbers) && IsConsecutive(report.Numbers) {
		return true
	}
	for i := 0; i < len(report.Numbers); i++ {
		var numbers []int
		numbers = append(numbers, report.Numbers...)
		numbers = slices.Delete(numbers, i, i + 1)

		if SafeDistances(numbers) && IsConsecutive(numbers) {
			return true
		}
	}
	return false
}


type Data struct {
	Raw     string
	Reports []Report
}

func (data *Data) SafeReports() int {
	return len(utils.Filter(data.Reports, func(report Report) bool {
		return report.IsSafe()
	}))
}
func (data *Data) SafeReportsIgnoreOne() int {
	return len(utils.Filter(data.Reports, func(report Report) bool {
		return report.IsSafeIgnoreOne()
	}))
}
func NewData(input string) *Data {
	lines := utils.Split(input, "\n")
	if len(lines) == 0 {
		log.Fatal("Lines should not be empty, ever", "Input", input)
	}
	return &Data{
		Raw: input,
		Reports: utils.Map(lines, func(line string) Report {
			return *NewReport(line)
		}),
	}
}

func IsConsecutive(numbers []int) bool {
	if len(numbers) < 2 {
		log.Fatal("Numbers were smaller than 2 elements", "numbers", numbers)
	}
	larger := numbers[0] < numbers[1]
	for i := 0; i < len(numbers) - 1; i++ {
		if larger && numbers[i] > numbers[i + 1] {
			return false
		} else if !larger && numbers[i] < numbers[i + 1] {
			return false
		}
	}
	return true
}
func SafeDistances(numbers []int) bool {
	if len(numbers) < 2 {
		return true
	}
	for i := 0; i < len(numbers)-1; i++ {
		dist := int(math.Abs(float64(numbers[i] - numbers[i+1])))
		if dist > 3 || dist < 1 {
			return false
		}
	}

	return true
}

func main() {
	session := os.Getenv("AOC_SESSION")
	input := utils.ReadHTTP(2024, 2, session)

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input))
	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input))

	os.Exit(0)
}

// part one
func part1(input string) int {
	return NewData(input).SafeReports()
}

// part two
func part2(input string) int {
	return NewData(input).SafeReportsIgnoreOne()
}
