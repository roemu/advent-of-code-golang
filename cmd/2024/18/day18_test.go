package main

import (
	"aoc/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay18(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("day18_default.txt")

	t.Run("part 1", func(t *testing.T) {
		expected := 22
		actual := part1(input, 12)

		assert.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := "6,1"
		actual := part2(input, 12)

		assert.Equal(expected, actual)
	})
}
