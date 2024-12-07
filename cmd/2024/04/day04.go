package main

import (
	"fmt"
	"os"

	"aoc/internal/utils"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
	NorthEast
	SouthEast
	NorthWest
	SouthWest
)

var gridWidth = 0
var gridHeight = 0

func GetVectorValues(dir Direction) (int, int) {
	switch dir {
	case North:
		return 0, -1
	case East:
		return 1, 0
	case South:
		return 0, 1
	case West:
		return -1, 0
	case NorthEast:
		return 1, -1
	case SouthEast:
		return 1, 1
	case NorthWest:
		return -1, -1
	case SouthWest:
		return -1, 1
	}
	return 0, 0
}

func CheckDirection(x int, y int, grid [][]string, direction Direction) int {
	vx, vy := GetVectorValues(direction)
	xmas := CharAt(x, y, grid) +
		CharAt(x+(vx*1), y+(vy*1), grid) +
		CharAt(x+(vx*2), y+(vy*2), grid) +
		CharAt(x+(vx*3), y+(vy*3), grid)
	if xmas == "XMAS"{
		return 1
	}
	return 0
}
func CheckDirectionMAS(x int, y int, grid [][]string, direction Direction) int {
	vx, vy := GetVectorValues(direction)
	xmas :=
		CharAt(x+(vx*1), y+(vy*1), grid) +
		CharAt(x, y, grid) +
		CharAt(x+(vx*(-1)), y+(vy*(-1)), grid)
	if xmas == "MAS"{
		return 1
	}
	return 0
}
func CharAt(x int, y int, grid [][]string) string {
	if x >= gridHeight || y >= gridWidth || x < 0 || y < 0 {
		return ""
	}
	return grid[x][y]
}
func StartSearch(grid [][]string) int {
	gridWidth = len(grid[0])
	gridHeight = len(grid)
	count := 0
	for x := range grid {
		for y := range grid[x] {
			count = count +
				CheckDirection(x, y, grid, North) +
				CheckDirection(x, y, grid, East) +
				CheckDirection(x, y, grid, South) +
				CheckDirection(x, y, grid, West) +
				CheckDirection(x, y, grid, NorthEast) +
				CheckDirection(x, y, grid, SouthEast) +
				CheckDirection(x, y, grid, NorthWest) +
				CheckDirection(x, y, grid, SouthWest)
		}
	}
	return count
}
func StartSearchMAS(grid [][]string) int {
	gridWidth = len(grid[0])
	gridHeight = len(grid)
	count := 0
	for x := range grid {
		for y := range grid[x] {
			mas := CheckDirectionMAS(x, y, grid, NorthEast) +
				CheckDirectionMAS(x, y, grid, SouthEast) +
				CheckDirectionMAS(x, y, grid, NorthWest) +
				CheckDirectionMAS(x, y, grid, SouthWest)
			if mas == 2 {
				count++;
			}
		}
	}
	return count
}

func NewInput(input string) [][]string {
	return utils.Map(utils.Split(input, "\n"), func(line string) []string {
		return utils.Map([]byte(line), func(letter byte) string {
			return string(letter)
		})
	})
}
func main() {
	session := os.Getenv("AOC_SESSION")
	input := utils.ReadHTTP(2024, 4, session)

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input))
	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input))

	os.Exit(0)
}

// part one
func part1(input string) int {
	return StartSearch(NewInput(input))
}

// part two
func part2(input string) int {
	return StartSearchMAS(NewInput(input))
}
