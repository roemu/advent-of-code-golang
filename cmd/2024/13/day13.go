package main

import (
	"os"

	"aoc/internal/utils"

	"github.com/charmbracelet/log"
)

type ClawMachine struct {
	ButtonA utils.IVec
	ButtonB utils.IVec
	Prize utils.IVec
}

func (clawMachine *ClawMachine) FindSolution() (a, b int) {
	det := clawMachine.ButtonB.Det(clawMachine.ButtonA)

	areaX := clawMachine.ButtonB.Det(clawMachine.Prize)
	areaY := clawMachine.Prize.Det(clawMachine.ButtonA)

	b = areaY / det
	a = areaX / det
	return
}
func (clawMachine *ClawMachine) Tokens(limited bool) int {
	a,b := clawMachine.FindSolution()
	if a < 0 || b < 0 || (limited && (a > 100 || b > 100)) {
		return 0
	}
	if a * clawMachine.ButtonA.X + b * clawMachine.ButtonB.X == clawMachine.Prize.X &&
		a * clawMachine.ButtonA.Y + b * clawMachine.ButtonB.Y == clawMachine.Prize.Y {
		return a*3 + b
	} else {
		return 0
	}
}

func ParseClawMachine(input string, prizeOffset int) *ClawMachine {
	lines := utils.Split(input, "\n")
	lineA := utils.Split(utils.Split(lines[0], ": ")[1], ", ")
	vecA := utils.IVec{
		X: utils.Atoi(lineA[0][1:]),
		Y: utils.Atoi(lineA[1][1:]),
	}
	lineB := utils.Split(utils.Split(lines[1], ": ")[1], ", ")
	vecB := utils.IVec{
		X: utils.Atoi(lineB[0][1:]),
		Y: utils.Atoi(lineB[1][1:]),
	}
	linePrize := utils.Split(utils.Split(lines[2], ": ")[1], ", ")
	vecPrize := utils.IVec{
		X: utils.Atoi(linePrize[0][2:]) + prizeOffset,
		Y: utils.Atoi(linePrize[1][2:]) + prizeOffset,
	}
	return &ClawMachine{
		ButtonA: vecA,
		ButtonB: vecB,
		Prize: vecPrize,
	}
}

func ParseInput(input string, prizeOffset int) []*ClawMachine {
	return utils.Map(utils.Split(input, "\n\n"), func(claw string) *ClawMachine {
		return ParseClawMachine(claw, prizeOffset)
	})
}

func main() {
	session := os.Getenv("AOC_SESSION")
	input := utils.ReadHTTP(2024, 13, session)

	log.Info("--- Part One ---")
	log.Infof("Result %d", part1(input))
	log.Info("--- Part Two ---")
	log.Infof("Result %d", part2(input))

	os.Exit(0)
}

// part one
func part1(input string) int {
	clawMachines := ParseInput(input, 0)
	return utils.Reduce(clawMachines, func(acc int, clawMachine *ClawMachine) int {
		return acc + clawMachine.Tokens(true)
	}, 0)
}

// part two
func part2(input string) int {
	clawMachines := ParseInput(input, 10000000000000)
	return utils.Reduce(clawMachines, func(acc int, clawMachine *ClawMachine) int {
		return acc + clawMachine.Tokens(false)
	}, 0)
}

