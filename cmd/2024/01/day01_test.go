package main

import (
	"aoc/internal/utils"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	assert := assert.New(t)
	input, err := utils.ReadFile("cmd/2024/01/input_test.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	println(input)

	t.Run("part 1", func(t *testing.T) {
		expected := 11
		actual := part1(input)

		assert.Equal(actual, expected)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 0
		actual := part2(input)

		assert.Equal(actual, expected)
	})
}
