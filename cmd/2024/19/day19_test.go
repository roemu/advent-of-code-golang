package main

import (
	"aoc/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay19(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("day19_default.txt")

	t.Run("part 1", func(t *testing.T) {
		expected := 6
		actual := part1(input)

		assert.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 16
		actual := part2(input)

		assert.Equal(expected, actual)
	})
}
