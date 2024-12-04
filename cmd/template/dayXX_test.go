package main

import (
	"aoc/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayXX(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("dayXX_default.txt")

	t.Run("part 1", func(t *testing.T) {
		expected := 0
		actual := part1(input)

		assert.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 0
		actual := part2(input)

		assert.Equal(expected, actual)
	})
}
