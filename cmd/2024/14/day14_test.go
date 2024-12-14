package main

import (
	"aoc/internal/utils"
	"testing"

	"github.com/charmbracelet/log"
	"github.com/stretchr/testify/assert"
)

func TestParseRobot(t *testing.T) {
	assert := assert.New(t)
	input := "p=2,4 v=-1,3"
	robot := ParseRobot(input)

	assert.Equal(utils.IVec{X:2,Y:4}, robot.Position)
	assert.Equal(utils.IVec{X:-1,Y:3}, robot.Velocity)
}

func TestRobotPositionAfterSeconds(t *testing.T) {
	assert := assert.New(t)
	input := "p=2,4 v=-1,3"
	robot := ParseRobot(input)

	assert.Equal(utils.IVec{X:2,Y:4}, robot.Position)
	assert.Equal(utils.IVec{X:-1,Y:3}, robot.Velocity)

	// robot.PositionAfterSeconds(1, 11, 7)
	//
	// assert.Equal(utils.IVec{X:1,Y:0}, robot.Position)
	// assert.Equal(utils.IVec{X:-1,Y:3}, robot.Velocity)

	robot.PositionAfterSeconds(8, 11, 7)
	log.Info("After seconds: %v", robot)
}
func TestParseInput(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("day14_default.txt")
	robots := ParseInput(input)

	assert.Len(robots, 12)
}

func TestDay14(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("day14_default.txt")

	t.Run("part 1", func(t *testing.T) {
		expected := 12
		actual := part1(input, 11, 7)

		assert.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 0
		actual := part2(input, 11, 7)

		assert.Equal(expected, actual)
	})
}
