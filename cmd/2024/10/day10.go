package main

import (
	"os"
	"strings"
	"sync"

	"aoc/internal/utils"

	"github.com/charmbracelet/log"
)

type HikeMap struct {
	Grid      [][]int
	MapHeight int
	MapWidth  int
}

func NewHikeMap(input string) *HikeMap {
	positions := [][]int{}
	mapHeight := len(utils.Split(input, "\n"))
	mapWidth := 0
	for _, line := range utils.Split(input, "\n") {
		mapWidth = len(line)
		row := []int{}
		for _, height := range utils.Split(line, "") {
			if height == "." {
				row = append(row, 0)
				continue
			}
			row = append(row, utils.Atoi(height))
		}
		positions = append(positions, row)
	}
	return &HikeMap{
		Grid:      positions,
		MapHeight: mapHeight,
		MapWidth:  mapWidth,
	}
}

func (hikeMap *HikeMap) TrailHeads() []utils.IVec2 {
	trailHeads := []utils.IVec2{}
	for y, row := range hikeMap.Grid {
		for x, h := range row {
			if h == 0 {
				trailHeads = append(trailHeads, utils.IVec2{int64(x), int64(y)})
			}
		}
	}
	return trailHeads
}

func (hikeMap *HikeMap) PrintGrid() {
	for _, row := range hikeMap.Grid {
		log.Info(strings.Join(utils.MapItoa(row), ""))
	}
}
func Start(trailHead utils.IVec2, hikeMap HikeMap, removeDuplicate bool) int {
	dirs := []utils.Direction{utils.North, utils.East, utils.South, utils.West}
	sum := []utils.IVec2{}
	for _, dir := range dirs {
		sum = append(sum, CheckPath(trailHead, dir, hikeMap)...)
	}
	if removeDuplicate {
		sum = utils.RemoveDuplicate(sum)
	}
	return len(sum)
}
func HeightOfPos(pos utils.IVec2, hikeMap HikeMap) int {
	return hikeMap.Grid[pos[1]][pos[0]]
}
func OutOfBounds(pos utils.IVec2, hikeMap HikeMap) bool {
	if pos[0] >= int64(hikeMap.MapWidth) || pos[0] < 0 {
		return true
	}
	if pos[1] >= int64(hikeMap.MapHeight) || pos[1] < 0 {
		return true
	}
	return false
}
func AddVector(vec1 utils.IVec2, vec2 utils.IVec2) utils.IVec2 {
	return utils.IVec2{vec1[0] + vec2[0], vec1[1] + vec2[1]}
}
func CheckPath(prevPos utils.IVec2, dir utils.Direction, hikeMap HikeMap) []utils.IVec2 {
	nextPos := AddVector(prevPos, utils.ToIVec2(dir))
	if OutOfBounds(nextPos, hikeMap) {
		return []utils.IVec2{}
	}
	prevHeight := HeightOfPos(prevPos, hikeMap)
	nextHeight := HeightOfPos(nextPos, hikeMap)
	if prevHeight+1 != nextHeight {
		return []utils.IVec2{}
	}
	if nextHeight == 9 {
		return []utils.IVec2{nextPos}
	}
	dirs := []utils.Direction{utils.North, utils.East, utils.South, utils.West}
	sum := []utils.IVec2{}
	for _, d := range dirs {
		sum = append(sum, CheckPath(nextPos, d, hikeMap)...)
	}
	return sum
}

func main() {
	session := os.Getenv("AOC_SESSION")
	input := utils.ReadHTTP(2024, 10, session)

	log.Info("--- Part One ---")
	log.Info("", "Result", part1(input))
	log.Info("--- Part Two ---")
	log.Info("", "Result", part2(input))

	os.Exit(0)
}

// part one
func part1(input string) int {
	hikeMap := NewHikeMap(input)
	trailHeads := hikeMap.TrailHeads()
	sum := 0
	var wg sync.WaitGroup
	for _, trailHead := range trailHeads {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sum += Start(trailHead, *hikeMap, true)
		}()
	}
	wg.Wait()
	return sum
}

// part two
func part2(input string) int {
	hikeMap := NewHikeMap(input)
	trailHeads := hikeMap.TrailHeads()
	sum := 0
	var wg sync.WaitGroup
	for _, trailHead := range trailHeads {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sum += Start(trailHead, *hikeMap, false)
		}()
	}
	wg.Wait()
	return sum
}
