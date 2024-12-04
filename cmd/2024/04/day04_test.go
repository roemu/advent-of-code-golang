package main

import (
	"aoc/internal/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestNewInput(t *testing.T) {
	assert := assert.New(t)
	array := NewInput(fmt.Sprintf("ab\ncd"))
	assert.Len(array, 2, array)
}
func TestStartSearch(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("./day04_default.txt")
	array := NewInput(input)
	res := StartSearch(array)
	assert.Equal(18, res)
}

func TestDay04(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("day04_default.txt")

	t.Run("part 1", func(t *testing.T) {
		expected := 18
		actual := part1(input)

		assert.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 9
		actual := part2(input)

		assert.Equal(expected, actual)
	})
}
