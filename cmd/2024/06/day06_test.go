package main

import (
	"aoc/internal/utils"
	"math"
	"testing"

	"github.com/quartercastle/vector"
	"github.com/stretchr/testify/assert"
)

func TestVectorRotation(t *testing.T) {
	assert := assert.New(t)
	direction := vector.Vector{0, 1}
	direction = direction.Rotate(90 * (math.Pi / 180))
	direction = vector.Vector{math.Round(direction.X()), math.Round(direction.Y())}
	assert.Equal(vector.Vector{-1, 0}, direction)
}

func TestMapInit(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("./day06_default.txt")
	guardMap, guard := NewGuardMap(input)
	assert.Equal(10, guardMap.MapWidth, "MapWidth")
	assert.Equal(10, guardMap.MapHeight, "MapHeight")
	assert.Len(guardMap.Obstacles, 8, "Obstacles amount")
	assert.Len(guardMap.PatroledPositions, 1, "Newly init, should have 1 positions visited")
	assert.Equal(vector.Vector{4,6}, guard.Position, "Guard position")
	assert.Equal(vector.Vector{0,-1}, guard.Direction, "Guard direction vector")
}

func TestEdgeCaseInputs(t *testing.T) {
	assert := assert.New(t)

	input := utils.ReadFile("./day06_edgecase_02.txt")
	assert.Equal(11, part1(input))

	input = utils.ReadFile("./day06_edgecase_01.txt")
	assert.Equal(4, part1(input))
}
func TestSimpleLoop(t *testing.T) {
	assert := assert.New(t)

	input := utils.ReadFile("./day06_simple.txt")
	assert.Equal(1, part2(input))
}

func TestDay06(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("day06_default.txt")

	t.Run("part 1", func(t *testing.T) {
		expected := 41
		actual := part1(input)

		assert.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 6
		actual := part2(input)

		assert.Equal(expected, actual)
	})
}
