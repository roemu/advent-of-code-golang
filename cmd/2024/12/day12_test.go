package main

import (
	"aoc/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAreas(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("day12_default.txt")

	areas := GenerateAreas(input)
	assert.Len(areas, 11)
}
func TestCountGapped(t *testing.T) {
	assert := assert.New(t)
	input := []utils.IVec{utils.IVec{X: 0}, utils.IVec{X: 1}, utils.IVec{X: 2}}
	uniques := CountGapped(input, utils.North)
	assert.Equal(1, uniques)

	input = []utils.IVec{utils.IVec{X: 1}, utils.IVec{X: 4}, utils.IVec{X: 5}}
	uniques = CountGapped(input, utils.North)
	assert.Equal(2, uniques)

	input = GenerateVecList([][]int{[]int{1, 0}, []int{2, 0}, []int{4, 0}, []int{5, 0}, []int{7, 0}, []int{5, 2}, []int{3, 4}, []int{4, 4}})
	uniques = CountGapped(input, utils.North)
	assert.Equal(5, uniques)

}
func GenerateVecList(input [][]int) []utils.IVec {
	return utils.Map(input, func(t []int) utils.IVec {
		return utils.IVec{X: t[0], Y: t[1]}
	})
}
func TestEdgeCase1(t *testing.T) {
	assert := assert.New(t)

	input := utils.ReadFile("./day12_edge_1.txt")
	assert.Equal(236, part2(input))
}
func TestEdgeCase2(t *testing.T) {
	assert := assert.New(t)

	input := utils.ReadFile("./day12_edge_2.txt")
	assert.Equal(196, part2(input))
}
func TestEdgeCase3(t *testing.T) {
	assert := assert.New(t)

	input := utils.ReadFile("./day12_edge_3.txt")
	areas := GenerateAreas(input)
	assert.Len(areas, 2)
}
func TestDay12(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("day12_default.txt")

	t.Run("part 1", func(t *testing.T) {
		expected := 1930
		actual := part1(input)

		assert.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 1206
		actual := part2(input)

		assert.Equal(expected, actual)
	})
}
