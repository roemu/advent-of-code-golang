package main

import (
	"aoc/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannel(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("day07_default.txt")

	equations := utils.Map(utils.Split(input, "\n"), func(line string) *Equation {
		return NewEquation(line)
	})
	sum := SumValidEquations(equations)
	assert.Equal(int64(3749), sum)

}

func TestEdgeCases(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("./day07_edge_1.txt")

	expected := int64(42)
	actual := part1(input)
	assert.Equal(expected, actual)
}
func TestGenerateCombinations(t *testing.T) {
	assert := assert.New(t)

	combinations := GenerateCombinations([]string{"*", "+"}, 2)
	assert.Len(combinations, 4)

	combinations = GenerateCombinations([]string{"*", "+"}, 5)
	assert.Len(combinations, 32)
}

func TestDay07(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("day07_default.txt")

	t.Run("part 1", func(t *testing.T) {
		expected := int64(3749)
		actual := part1(input)

		assert.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := int64(0)
		actual := part2(input)

		assert.Equal(expected, actual)
	})
}
