package main

import (
	"aoc/internal/utils"
	"testing"

	"github.com/quartercastle/vector"
	"github.com/stretchr/testify/assert"
)

func TestCreateAntinodePair(t *testing.T) {
	assert := assert.New(t)
	antenna1 := vector.Vector{1, 1}
	antenna2 := vector.Vector{3, 3}
	antinode1 := vector.Vector{5,5}
	antinode2 := vector.Vector{-1, -1}
	antinodes := CreateAntinodePair(antenna1, antenna2)
	assert.Len(antinodes, 2)
	assert.Contains(antinodes, antinode1)
	assert.Contains(antinodes, antinode2)
}
func TestFilterDuplicates(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("day08_default.txt")
	cityMap := NewCityMap(input)
	antenna1 := vector.Vector{1, 0}
	antenna2 := vector.Vector{3, 3}
	antenna3 := vector.Vector{float64(cityMap.MapWidth) + 1, 3}
	antenna4 := vector.Vector{3, 3}
	antinodes := FilterDuplicates([]vector.Vector{antenna1, antenna2, antenna3, antenna4})
	assert.Len(antinodes, 3)
}
func TestFilterOutOfBounds(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("day08_default.txt")
	cityMap := NewCityMap(input)
	antenna1 := vector.Vector{-1, 0}
	antenna2 := vector.Vector{3, 3}
	antenna3 := vector.Vector{float64(cityMap.MapWidth) + 1, 3}
	antinodes := FilterOutOfBounds([]vector.Vector{antenna1, antenna2, antenna3}, cityMap)
	assert.Len(antinodes, 1)
}
func TestCreateAntinodes(t *testing.T) {
	assert := assert.New(t)
	antenna1 := vector.Vector{1, 1}
	antenna2 := vector.Vector{3, 3}
	antinodes := CreateAntinodes([]vector.Vector{antenna1, antenna2})
	assert.Len(antinodes, 2)
}
func TestNewCityMap(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("day08_default.txt")

	cityMap := NewCityMap(input)
	assert.Len(cityMap.Antennas, 2)
	assert.Len(cityMap.Antennas["0"], 4)
	assert.Len(cityMap.Antennas["A"], 3)
}
func TestDay08(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("day08_default.txt")

	t.Run("part 1", func(t *testing.T) {
		expected := 14
		actual := part1(input)

		assert.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 34
		actual := part2(input)

		assert.Equal(expected, actual)
	})
}
