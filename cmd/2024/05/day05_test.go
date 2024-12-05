package main

import (
	"aoc/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUpdate(t *testing.T) {
	assert := assert.New(t)
	update := NewUpdate("25,47,78,26,17,0,82")

	assert.Len(update.Pages, 7)
}
func TestMiddlePages(t *testing.T) {
	assert := assert.New(t)
	update := NewUpdate("25,47,78,26,17,0,82")

	assert.Len(update.Pages, 7)
	assert.Equal(26, update.MiddlePage())
}
func TestNewInput(t *testing.T) {
	assert := assert.New(t)
	input := NewInput(utils.ReadFile("./day05_default.txt"))

	assert.Len(input.Updates, 6)
	totalRules := 0
	for _, val := range input.AfterRules {
		totalRules = len(val) + totalRules
	}
	assert.Equal(21, totalRules)
	assert.Equal(143, input.SumMiddlePages()) // Will fail when extra logic applied
}
func TestTopoSort(t *testing.T){
	assert := assert.New(t)
	input := NewInput(utils.ReadFile("./day05_toposort.txt"))
	input.Updates[0].FixOrder(input)

	assert.Equal([]int{5,4,3,2,1,0},input.Updates[0].Pages)
}
func TestMiddlePagesCorrect(t *testing.T){
	assert := assert.New(t)
	input := NewInput(utils.ReadFile("./day05_default.txt"))
	assert.Equal(123, input.SumdMiddlePagesCorrect())
}

func TestDay05(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("day05_default.txt")

	t.Run("part 1", func(t *testing.T) {
		expected := 143
		actual := part1(input)

		assert.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 123
		actual := part2(input)

		assert.Equal(expected, actual)
	})
}
