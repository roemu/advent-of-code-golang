package main

import (
	"os"

	"aoc/internal/utils"

	"github.com/beefsack/go-astar"
	"github.com/charmbracelet/log"
)

var GRID_SIZE int = 7

type Tile struct {
	Pos utils.IVec
}

func (tile Tile) PathNeighbors() []astar.Pather {
	potentials := []utils.IVec{
		tile.Up(),
		tile.Down(),
		tile.Left(),
		tile.Right(),
	}
	potentials = utils.Filter(potentials, func(pot utils.IVec) bool {
		return !OutOfBounds(pot) && !WALLS[pot]
	})
	return utils.Map(potentials, func(pot utils.IVec) astar.Pather {
		return Tile{Pos: pot}
	})
}
func (tile *Tile) Up() utils.IVec {
	return tile.Pos.Add(utils.ToIVec(utils.North))
}
func (tile *Tile) Down() utils.IVec {
	return tile.Pos.Add(utils.ToIVec(utils.South))
}
func (tile *Tile) Left() utils.IVec {
	return tile.Pos.Add(utils.ToIVec(utils.West))
}
func (tile *Tile) Right() utils.IVec {
	return tile.Pos.Add(utils.ToIVec(utils.East))
}
func OutOfBounds(vec utils.IVec) bool {
	if vec.X < 0 || vec.X >= GRID_SIZE {
		return true
	}
	if vec.Y < 0 || vec.Y >= GRID_SIZE {
		return true
	}
	return false
}
func (tile Tile) PathNeighborCost(to astar.Pather) float64 {
	return 1
}
func (tile Tile) PathEstimatedCost(to astar.Pather) float64 {
	return 0
}

func main() {
	session := os.Getenv("AOC_SESSION")
	input := utils.ReadHTTP(2024, 18, session)
	GRID_SIZE = 71

	log.Info("--- Part One ---")
	log.Infof("Result %d", part1(input, 1024))
	log.Info("--- Part Two ---")
	log.Infof("Result %s", part2(input, 1024))

	os.Exit(0)
}

func ParseInput(input string) []utils.IVec {
	return utils.Map(utils.Split(input, "\n"), func(line string) utils.IVec {
		a, b := utils.SplitLeftRight(line, ",")
		return utils.IVec{
			X: utils.Atoi(a),
			Y: utils.Atoi(b),
		}
	})
}

var WALLS = make(map[utils.IVec]bool)

// part one
func part1(input string, steps int) int {
	values := ParseInput(input)[:steps]
	for _, vec := range values {
		WALLS[vec] = true
	}
	start := Tile{Pos: utils.IVec{X: 0, Y: 0}}
	end := Tile{Pos: utils.IVec{X: GRID_SIZE - 1, Y: GRID_SIZE - 1}}

	_, distance, found := astar.Path(start, end)
	if !found {
		log.Fatal("Shouldn't happen")
	}
	return int(distance)
}

// part two
func part2(input string, steps int) string {
	values := ParseInput(input)
	for i, vec := range values {
		if i >= steps {
			break
		}
		WALLS[vec] = true
	}
	start := Tile{Pos: utils.IVec{X: 0, Y: 0}}
	end := Tile{Pos: utils.IVec{X: GRID_SIZE - 1, Y: GRID_SIZE - 1}}

	for {
		WALLS[values[steps]] = true
		_, _, found := astar.Path(start, end)
		if found {
			steps++
		}
		if !found {
			return utils.Itoa(values[steps].X) + "," + utils.Itoa(values[steps].Y)
		}
	}
}
