package main

import (
	"aoc/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatMap(t *testing.T) {
	assert := assert.New(t)
	input := [][]int{[]int{1,2,3,4}, []int{5,6,7,8}}

	assert.Len(utils.FlatMap(input, func(arr []int) []int {
		return arr
	}), 8)
}
func TestDay11(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("day11_default.txt")

	t.Run("part 1", func(t *testing.T) {
		expected := 55312
		actual := part1(input)

		assert.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 65601038650482
		actual := part2(input)
		assert.Equal(expected, actual)
	})

}
