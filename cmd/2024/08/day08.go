package main

import (
	"os"

	"aoc/internal/utils"

	"github.com/charmbracelet/log"
	"github.com/quartercastle/vector"
)

type CityMap struct {
	Antennas  map[string][]vector.Vector
	Antinodes []vector.Vector
	MapWidth  int
	MapHeight int
}

func NewCityMap(input string) *CityMap {
	mapHeight := len(utils.Split(input, "\n"))
	mapWidth := 0
	antennas := make(map[string][]vector.Vector, mapHeight*mapHeight)
	for y, line := range utils.Split(input, "\n") {
		mapWidth = len(line)
		for x, obj := range utils.Split(line, "") {
			if obj != "." {
				antennas[obj] = append(antennas[obj], vector.Vector{float64(x), float64(y)})
			}
		}
	}
	return &CityMap{
		MapWidth:  mapWidth,
		MapHeight: mapHeight,
		Antennas:  antennas,
	}
}

// Given a list of antennas, of the same type, will return a list of antinodes
// NOTE: Could contain duplicates
func CreateAntinodes(antennas []vector.Vector) (antinode []vector.Vector) {
	if len(antennas) < 2 {
		return []vector.Vector{}
	}
	antinodes := []vector.Vector{}
	for i, antennaRoot := range antennas {
		for _, antennaSecondary := range antennas[i+1:] {
			antinodes = append(antinodes, CreateAntinodePair(antennaRoot, antennaSecondary)...)
		}
	}
	return antinodes
}

// Given a list of antennas, of the same type, will return a list of antinodes
// NOTE: Could contain duplicates
func CreateAntinodesUnrestricted(antennas []vector.Vector, cityMap *CityMap) (antinode []vector.Vector) {
	if len(antennas) < 2 {
		return []vector.Vector{}
	}
	antinodes := []vector.Vector{}
	for _, antennaRoot := range antennas {
		for _, antennaSecondary := range antennas {
			antinodes = append(antinodes, CreateAntinodeLane(antennaRoot, antennaSecondary, cityMap)...)
		}
	}
	return antinodes
}

// Given a root antenna (antennaA) and a secondary antenna (antennaB) will return exactly two antinodes
// The antinodes may be out of bounds and may not be unique
func CreateAntinodePair(antennaA vector.Vector, antennaB vector.Vector) []vector.Vector {
	distance := antennaB.Sub(antennaA)
	antinode1 := antennaA.Add(distance.Scale(2))
	antinode2 := antennaA.Sub(distance)

	return []vector.Vector{antinode1, antinode2}
}

// Given a root antenna (antennaA) and a secondary antenna (antennaB) will return exactly two antinodes
// The antinodes may be out of bounds and may not be unique
func CreateAntinodeLane(antennaA vector.Vector, antennaB vector.Vector, cityMap *CityMap) []vector.Vector {
	antinodes := []vector.Vector{}
	distance := antennaB.Sub(antennaA)
	if distance.Magnitude() == 0 {
		return []vector.Vector{}
	}

	step := float64(1)
	for {
		newNode := antennaA.Add(distance.Scale(step))
		if OutOfBounds(newNode, cityMap) {
			break;
		}
		antinodes = append(antinodes, newNode)
		step++
	}
	step = float64(1)
	for {
		newNode := antennaA.Sub(distance.Scale(step))
		if OutOfBounds(newNode, cityMap) {
			break;
		}
		antinodes = append(antinodes, newNode)
		step++
	}

	return antinodes
}

func OutOfBounds(pos vector.Vector, cityMap *CityMap) bool {
	return pos.X() < 0 ||
		pos.X() >= float64(cityMap.MapWidth) ||
		pos.Y() < 0 ||
		pos.Y() >= float64(cityMap.MapHeight)
}
func FilterOutOfBounds(antinodes []vector.Vector, cityMap *CityMap) []vector.Vector {
	return utils.Filter(antinodes, func(antinode vector.Vector) bool {
		return !OutOfBounds(antinode, cityMap)
	})
}
func FilterDuplicates(antinodes []vector.Vector) []vector.Vector {
	res := []vector.Vector{}
	for _, antinode := range antinodes {
		if !utils.ContainsVector(res, antinode) {
			res = append(res, antinode)
		}
	}
	return res
}

func main() {
	session := os.Getenv("AOC_SESSION")
	input := utils.ReadHTTP(2024, 8, session)

	log.Info("--- Part One ---")
	log.Info("", "Result", part1(input))
	log.Info("--- Part Two ---")
	log.Info("", "Result", part2(input))

	os.Exit(0)
}

// part one
func part1(input string) int {
	cityMap := NewCityMap(input)
	for _, val := range cityMap.Antennas {
		cityMap.Antinodes = append(cityMap.Antinodes, CreateAntinodes(val)...)
	}
	cityMap.Antinodes = FilterOutOfBounds(cityMap.Antinodes, cityMap)
	cityMap.Antinodes = FilterDuplicates(cityMap.Antinodes)
	// DrawCityMapToFile(cityMap, "map-completed.txt")
	return len(cityMap.Antinodes)
}

// part two
func part2(input string) int {
	cityMap := NewCityMap(input)
	for _, val := range cityMap.Antennas {
		cityMap.Antinodes = append(cityMap.Antinodes, CreateAntinodesUnrestricted(val, cityMap)...)
	}
	cityMap.Antinodes = FilterOutOfBounds(cityMap.Antinodes, cityMap)
	cityMap.Antinodes = FilterDuplicates(cityMap.Antinodes)
	// DrawCityMapToFile(cityMap, "map-completed-2.txt")
	return len(cityMap.Antinodes)
}
