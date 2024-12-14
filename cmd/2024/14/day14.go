package main

import (
	"fmt"
	"os"
	"strings"

	"aoc/internal/utils"

	"github.com/charmbracelet/log"
)

type Robot struct {
	Position utils.IVec
	Velocity utils.IVec
}

func ParseRobot(input string) *Robot {
	p, v := utils.SplitLeftRight(input, " ")
	pos := utils.IVec{
		X: utils.Atoi(utils.Split(p[2:], ",")[0]),
		Y: utils.Atoi(utils.Split(p[2:], ",")[1]),
	}
	vel := utils.IVec{
		X: utils.Atoi(utils.Split(v[2:], ",")[0]),
		Y: utils.Atoi(utils.Split(v[2:], ",")[1]),
	}
	return &Robot{
		Position: pos,
		Velocity: vel,
	}
}

func ParseInput(input string) []*Robot {
	return utils.Map(utils.Split(input, "\n"), func(line string) *Robot {
		return ParseRobot(line)
	})
}

func (robot *Robot) PositionAfterSeconds(seconds, width, height int) {
	actual := robot.Position.Add(robot.Velocity.Scale(seconds))
	robot.Position.X = utils.Wrap(actual.X, width)
	robot.Position.Y = utils.Wrap(actual.Y, height)
}

func (robot *Robot) Quadrant(width, height int) int {
	if robot.Position.X < width/2 && robot.Position.Y < height/2 {
		return 0
	} else if robot.Position.X > width/2 && robot.Position.Y < height/2 {
		return 1
	} else if robot.Position.X < width/2 && robot.Position.Y > height/2 {
		return 2
	} else if robot.Position.X > width/2 && robot.Position.Y > height/2 {
		return 3
	} else {
		return -1
	}
}

func main() {
	session := os.Getenv("AOC_SESSION")
	input := utils.ReadHTTP(2024, 14, session)

	log.Info("--- Part One ---")
	log.Infof("Result %d", part1(input, 101, 103))
	log.Info("--- Part Two ---")
	log.Infof("Result %d", part2(input, 101, 103))

	os.Exit(0)
}

func PrintRobotsAtSecond(input string, w, h, seconds int) string {
	robots := ParseInput(input)
	robots = utils.Map(robots, func(robot *Robot) *Robot {
		robot.PositionAfterSeconds(seconds, w, h)
		return robot
	})

	out := ""
	for i := range h {
		line := ""
		for j := range w {
			gridPos := utils.IVec{X: j, Y: i}
			val := utils.Itoa(len(utils.Filter(robots, func(rob *Robot) bool {
				return rob.Position == gridPos
			})))
			if val == "0" {
				line = line + "."
				continue
			}
			line = line + val
		}
		out = out + fmt.Sprintf("%s\n", line)
	}

	return out
}

// part one
func part1(input string, w, h int) int {
	robots := ParseInput(input)
	quadrants := make(map[int]int)
	for _, robot := range robots {
		robot.PositionAfterSeconds(100, w, h)
		quadrants[robot.Quadrant(w, h)]++
	}
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

// part two
func part2(input string, w, h int) int {
	robots := ParseInput(input)
	seconds := 0
	const window int = 100
	unqiuePositions := make(map[utils.IVec]bool)
	for {

		for _, robot := range robots {
			robot.PositionAfterSeconds(1, w, h)
			unqiuePositions[robot.Position] = true
		}

		res := PrintRobotsAtSecond(input, w, h, seconds)
		if strings.Contains(res, "11111111111") {
			return seconds
		}

		seconds++
		if seconds > 10000 {
			break
		}
	}
	return 0
}
