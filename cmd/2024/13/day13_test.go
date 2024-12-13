package main

import (
	"aoc/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVecDet(t *testing.T) {
	assert := assert.New(t)
	vecA := utils.IVec{
		X: 4,
		Y: 2,
	}
	vecB := utils.IVec{
		X: -1,
		Y: 1,
	}
	assert.Equal(6, vecA.Det(vecB))
}
func TestParseClawMachine(t *testing.T) {
	assert := assert.New(t)
	input := `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400`
	clawMachine := ParseClawMachine(input, 0)
	assert.Equal(utils.IVec{X:94, Y:34}, clawMachine.ButtonA)
	assert.Equal(utils.IVec{X:22, Y:67}, clawMachine.ButtonB)
	assert.Equal(utils.IVec{X:8400, Y:5400}, clawMachine.Prize)

	input = `Button A: X+94, Y+34
Button B: X+22, Y-67
Prize: X=8400, Y=5400`
	clawMachine = ParseClawMachine(input, 0)
	assert.Equal(utils.IVec{X:94, Y:34}, clawMachine.ButtonA)
	assert.Equal(utils.IVec{X:22, Y:-67}, clawMachine.ButtonB)
	assert.Equal(utils.IVec{X:8400, Y:5400}, clawMachine.Prize)
}
func TestParseInput(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("day13_default.txt")
	clawMachines := ParseInput(input, 0)
	assert.Len(clawMachines, 4)
}

func TestDay13(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("day13_default.txt")

	t.Run("part 1", func(t *testing.T) {
		expected := 480
		actual := part1(input)

		assert.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 875318608908
		actual := part2(input)

		assert.Equal(expected, actual)
	})
}
