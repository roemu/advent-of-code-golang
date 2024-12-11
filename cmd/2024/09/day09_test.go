package main

import (
	"aoc/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestStart(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("day09_default.txt")
	blocks := Start(input)

	assert.Equal(2858, CalculateChecksum(blocks))

}
func TestDay09(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("day09_default.txt")

	t.Run("part 1", func(t *testing.T) {
		expected := 1928
		actual := part1(input)

		assert.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 2858
		actual := part2(input)

		assert.Equal(expected, actual)
	})
}
func TestUtilsSwapElements(t *testing.T) {
	assert := assert.New(t)
	arr := []int{10, 11, 12, 10, 11, 12}
	arr = utils.SwapElements(arr, 2, 0, 2)
	assert.Len(arr, 6)
	assert.Equal([]int{12, 10, 10, 11, 11, 12}, arr)

	arr = []int{10, 11, 12, 10, 11, 12}
	arr = utils.SwapElements(arr, 0, 0, 2)
	assert.Len(arr, 6)
	assert.Equal([]int{10, 11, 12, 10, 11, 12}, arr)
}
func TestUtilsAll(t *testing.T) {
	assert := assert.New(t)

	arr := []int{10, 11, 12}

	res := utils.All(arr, func(num int) bool {
		return num > 9
	})
	assert.True(res)
}
func TestGetSuitableDotRange(t *testing.T) {
	assert := assert.New(t)
	start, found := FindSuitableDotRange([]string{"a", "a", ".", ".", "b", "b", "."}, 2)
	assert.Equal(2, start)
	assert.True(found)
	start, found = FindSuitableDotRange([]string{"a", "a", ".", ".", "b", "b", "."}, 3)
	assert.Equal(-1, start)
	assert.False(found)
	start, found = FindSuitableDotRange([]string{"a", "a", ".", ".", "b", "b", "."}, 1)
	assert.Equal(2, start)
	assert.True(found)
	start, found = FindSuitableDotRange([]string{".", "a", ".", ".", "b", "b", "."}, 1)
	assert.Equal(0, start)
	assert.True(found)
	start, found = FindSuitableDotRange([]string{"a", "a", "a", "a", "b", "b", "."}, 1)
	assert.Equal(6, start)
	assert.True(found)
	start, found = FindSuitableDotRange([]string{"a", "a", "a", "a", "b", ".", "x"}, 1)
	assert.Equal(5, start)
	assert.True(found)
	start, found = FindSuitableDotRange([]string{ "0", "0", "9", "9", ".", "1", "1", "1", "7", "7", "7", "2", "4", "4", ".", "3", "3", "3", ".", ".", ".", ".", "5", "5", "5", "5", ".", "6", "6", "6", "6", ".", ".", ".", ".", ".", "8", "8", "8", "8", ".", "."}, 1)
	assert.Equal(4, start)
	assert.True(found)
}
func TestGetRange(t *testing.T) {
	assert := assert.New(t)
	start, length, char := GetLastFileRange([]string{"a", "a", ".", ".", "b", "b", "."}, &[]string{})
	assert.Equal(4, start, "startIndex")
	assert.Equal(2, length, "length")
	assert.Equal("b", char)
	start, length, char = GetLastFileRange([]string{"a", "a", ".", ".", "b", "b", "."}, &[]string{"b"})
	assert.Equal(0, start, "startIndex")
	assert.Equal(2, length, "length")
	assert.Equal("a", char)
	start, length, char = GetLastFileRange([]string{"a", "a", ".", ".", "b", "b", "."}, &[]string{"b", "a"})
	assert.Equal(-1, start, "startIndex")
	assert.Equal(0, length, "length")
	assert.Equal("", char)

}
