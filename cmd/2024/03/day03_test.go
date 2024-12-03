package main

import (
	"aoc/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAllMatches(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("day03_default.txt")

	output := FindAllMatches(input)
	assert.Len(output, 4, input)
}

func TestDay03(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("day03_default.txt")

	t.Run("part 1", func(t *testing.T) {
		expected := 161
		actual := part1(input)

		assert.Equal(expected, actual)
	})

	input = utils.ReadFile("day03_default_part2.txt")
	t.Run("part 2", func(t *testing.T) {
		expected := 48
		actual := part2(input)

		assert.Equal(expected, actual)
	})
}
