package main

import (
	"math"
	"os"
	"slices"
	"sort"

	"aoc/internal/utils"

	"github.com/charmbracelet/log"
)

type Area struct {
	Name  string
	Tiles map[utils.IVec][]utils.IVec // Position | Borders
}

func (area *Area) Price(bulk bool) int {
	if bulk {
		return area.Size() * area.CompactPerimiter()
	}
	return area.Size() * area.Perimiter()
}

func (area *Area) CompactPerimiter() int {
	dirs := []utils.Direction{utils.North, utils.East, utils.South, utils.West}
	sum := 0
	for _, dir := range dirs {
		matchingTiles := []utils.IVec{}
		for tile, borders := range area.Tiles {
			if slices.Contains(borders, utils.ToIVec(dir)) {
				matchingTiles = append(matchingTiles, tile)
			}
		}

		sort.Slice(matchingTiles, func(i, j int) bool {
			yAxis := matchingTiles[i].Y - matchingTiles[j].Y
			xAxis := matchingTiles[i].X - matchingTiles[j].X
			if dir == utils.North || dir == utils.South {
				if yAxis == 0 {
					return xAxis > 0
				}
				return yAxis > 0
			} else {

				if xAxis == 0 {
					return yAxis > 0
				}
				return xAxis > 0
			}
		})
		sum += CountGapped(matchingTiles, dir)
	}
	return sum
}

func CountGapped(tiles []utils.IVec, dir utils.Direction) int {
	lineSeparated := make(map[int][]int)
	for _, tile := range tiles {
		if dir == utils.North || dir == utils.South {
			lineSeparated[tile.Y] = append(lineSeparated[tile.Y], tile.X)
		} else {
			lineSeparated[tile.X] = append(lineSeparated[tile.X], tile.Y)
		}
	}
	if len(lineSeparated) == 0 {
		log.Fatal("I don't know", tiles, dir)
	}
	sum := 0
	for _, val := range lineSeparated {
		uniques := 1
		for i := 1; i < len(val); i++ {
			if math.Abs(float64(val[i - 1] - val[i])) != 1 {
				uniques++
			}
		}
		sum += uniques
	}
	return sum
}

func (area *Area) Perimiter() int {
	sum := 0
	for _, borders := range area.Tiles {
		sum += len(borders)
	}
	return sum
}
func (area *Area) Size() int {
	return len(area.Tiles)
}

func GenerateAreas(input string) []Area {
	farm := map[utils.IVec]string{}
	areas := []Area{}
	for y, line := range utils.Split(input, "\n") {
		for x, char := range utils.Split(line, "") {
			vec := utils.IVec{
				X: x,
				Y: y,
			}
			farm[vec] = char
		}
	}
	seen := map[utils.IVec]bool{}
	for key, val := range farm {
		if _, ok := seen[key]; ok {
			continue
		}
		areas = append(areas, NewArea(key, val, farm, &seen))
	}
	return areas
}

func NewArea(vec utils.IVec, char string, farm map[utils.IVec]string, seen *map[utils.IVec]bool) Area {
	area := Area{
		Name:  char,
		Tiles: map[utils.IVec][]utils.IVec{},
	}
	area.Tiles[vec] = []utils.IVec{}
	(*seen)[vec] = true
	CheckEachDir(vec, &area, farm, seen)
	return area
}

func CheckEachDir(vec utils.IVec, area *Area, farm map[utils.IVec]string, seen *map[utils.IVec]bool) {
	dirs := []utils.Direction{utils.North, utils.East, utils.South, utils.West}
	for _, dir := range dirs {
		next := vec.Add(utils.ToIVec(dir))
		if _, ok := area.Tiles[next]; ok { // Already in area, skip
			continue
		}
		if char, ok := farm[next]; ok && char == area.Name { // Not in any area and same char
			(*seen)[next] = true
			area.Tiles[next] = []utils.IVec{}
			CheckEachDir(next, area, farm, seen)
		} else if !ok || char != area.Name {
			area.Tiles[vec] = append(area.Tiles[vec], utils.ToIVec(dir))
		}
	}
}

func main() {
	session := os.Getenv("AOC_SESSION")
	input := utils.ReadHTTP(2024, 12, session)

	log.Info("--- Part One ---")
	log.Info("", "Result", part1(input))
	log.Info("--- Part Two ---")
	log.Info("", "Result", part2(input))

	os.Exit(0)
}

// part one
func part1(input string) int {
	areas := GenerateAreas(input)
	return utils.Reduce(areas, func(acc int, area Area) int {
		return acc + area.Price(false)
	}, 0)
}

// part two
func part2(input string) int {
	areas := GenerateAreas(input)
	return utils.Reduce(areas, func(acc int, area Area) int {
		return acc + area.Price(true)
	}, 0)
}
