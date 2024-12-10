package main

import (
	"fmt"
	"os"
	"sync"

	"aoc/internal/utils"

	"github.com/charmbracelet/log"
	"github.com/quartercastle/vector"
)

type Direction vector.Vector

type GuardMap struct {
	Raw               string
	Obstacles         []vector.Vector
	LoopObstacles     []vector.Vector
	PatroledPositions []PatroledPosition
	MapWidth          int
	MapHeight         int
}
type Guard struct {
	Position  vector.Vector
	Direction vector.Vector

	// for part 2
	SeenPositions []PatroledPosition
}
type PatroledPosition struct {
	Position  vector.Vector
	Direction vector.Vector
}

func PositionDirectionPatroled(patroled []PatroledPosition, position PatroledPosition) bool {
	for _, patrol := range patroled {
		if patrol.Position.Equal(position.Position) && patrol.Direction.Equal(position.Direction) {
			return true
		}
	}
	return false
}
func PositionPatroled(patroled []PatroledPosition, position PatroledPosition) bool {
	for _, patrol := range patroled {
		if patrol.Position.Equal(position.Position) {
			return true
		}
	}
	return false
}
func (guard *Guard) StartPatrol(guardMap *GuardMap, startPosition vector.Vector) (uniquePositions []PatroledPosition) {
	newlyPatroledPositions := []PatroledPosition{}
	for {
		patroled, reachedEdge := guard.Move(guardMap)
		if reachedEdge {
			break
		}
		newlyPatroledPositions = append(newlyPatroledPositions, patroled)
	}
	return newlyPatroledPositions
}
func (guard *Guard) Move(guardMap *GuardMap) (patroledPosition PatroledPosition, reachedEdge bool) {
	nextPos := guard.Position.Add(guard.Direction)
	if PositionOutsideOfMap(nextPos, guardMap.MapWidth, guardMap.MapHeight) {
		return PatroledPosition{}, true
	}
	if utils.ContainsVector(guardMap.Obstacles, nextPos) {
		rotated := utils.RotateVector(guard.Direction, 90)
		guard.Direction = rotated
		return PatroledPosition{Position: guard.Position, Direction: rotated}, false
	}
	guard.Position = guard.Position.Add(guard.Direction)
	return PatroledPosition{Position: guard.Position, Direction: guard.Direction}, false
}

func (guard *Guard) HasSeenPosition(position vector.Vector, direction vector.Vector) bool {
	for _, patroled := range guard.SeenPositions {
		if patroled.Position.Equal(position) && patroled.Direction.Equal(direction) {
			return true
		}
	}
	return false
}

func (guard *Guard) StartLoopSearch(guardMap *GuardMap) (foundLoop bool) {
	for {
		patroled, reachedEdge := guard.Move(guardMap)
		if reachedEdge {
			return false
		}

		if guard.HasSeenPosition(patroled.Position, patroled.Direction) {
			return true
		}
		guard.SeenPositions = append(guard.SeenPositions, patroled)
	}
}

func PositionOutsideOfMap(nextPos vector.Vector, mapWidth, mapHeight int) bool {
	return nextPos.X() >= float64(mapWidth) ||
		nextPos.X() < 0 ||
		nextPos.Y() >= float64(mapHeight) ||
		nextPos.Y() < 0
}

func NewGuardMap(input string) (*GuardMap, *Guard) {
	lines := utils.Split(input, "\n")
	guardMap := &GuardMap{
		MapWidth:          0,
		MapHeight:         len(lines),
		Obstacles:         []vector.Vector{},
		PatroledPositions: []PatroledPosition{},
		LoopObstacles:     []vector.Vector{},
	}
	guard := &Guard{}
	for y, line := range lines {
		guardMap.MapWidth = len(line)
		for x, object := range utils.Split(line, "") {
			pos := vector.Vector{float64(x), float64(y)}
			if object == "." {
				continue
			} else if object == "#" {
				guardMap.Obstacles = append(guardMap.Obstacles, pos)
			} else {
				guard.Position = pos
				guard.Direction = DetermineGuardDirection(object)
				guardMap.PatroledPositions = append(guardMap.PatroledPositions, PatroledPosition{Position: pos, Direction: DetermineGuardDirection(object)}) // Set starting position as patroled, since it checks only future positions
			}
		}
	}
	return guardMap, guard
}

func DetermineGuardDirection(char string) vector.Vector {
	switch char {
	case ">":
		return vector.Vector{1, 0}
	case "<":
		return vector.Vector{-1, 0}
	case "^":
		return vector.Vector{0, -1}
	case "v":
		return vector.Vector{0, 1}
	default:
		log.Fatal("Unparsable direction given: ", "char", char)
	}
	return nil
}

func main() {
	session := os.Getenv("AOC_SESSION")
	input := utils.ReadHTTP(2024, 6, session)

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input))
	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input))

	os.Exit(0)
}

// part one
func part1(input string) int {
	guardMap, guard := NewGuardMap(input)
	guardMap.PatroledPositions = append(guardMap.PatroledPositions,
		guard.StartPatrol(guardMap, guard.Position)...)
	appeared := []vector.Vector{}
	guardMap.PatroledPositions = utils.Filter(guardMap.PatroledPositions, func(pos PatroledPosition) bool {
		if utils.ContainsVector(appeared, pos.Position) {
			return false
		}
		appeared = append(appeared, pos.Position)
		return true
	})
	return len(guardMap.PatroledPositions)
}

// part two
func part2(input string) int {
	guardMap, guard := NewGuardMap(input)
	guardMap.PatroledPositions = append(guardMap.PatroledPositions,
		guard.StartPatrol(guardMap, guard.Position)...)

	var wg sync.WaitGroup
	for _, patroled := range guardMap.PatroledPositions {
		wg.Add(1)
		rotated := utils.RotateVector(patroled.Direction, 90)
		funnyGuard := &Guard{
			Position: patroled.Position,
			// We turn him to simulate a obstacle here
			Direction:     rotated,
			SeenPositions: []PatroledPosition{
				PatroledPosition{Position: patroled.Position, Direction: patroled.Direction},
			},
		}
		customObstaclePos := patroled.Position.Clone().Add(patroled.Direction)
		go func() {
			defer wg.Done()
			if funnyGuard.StartLoopSearch(guardMap) {
				guardMap.LoopObstacles = append(guardMap.LoopObstacles, customObstaclePos)
			}
		}()

	}
	wg.Wait()
	appeared := []vector.Vector{}
	guardMap.LoopObstacles = utils.Filter(guardMap.LoopObstacles, func(pos vector.Vector) bool {
		if utils.ContainsVector(appeared, pos) && !PositionOutsideOfMap(pos, guardMap.MapWidth, guardMap.MapHeight) {
			return false
		}
		appeared = append(appeared, pos)
		return true
	})
	return len(guardMap.LoopObstacles)
}
