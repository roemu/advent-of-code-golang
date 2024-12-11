package main

import (
	"os"

	"aoc/internal/utils"

	"github.com/charmbracelet/log"
)


var memory Memoized

func (m *Memoized) Call(digit int, blinks int) int {
	key := Key{Digit: digit, Blinks: blinks}
	if _, ok := m.Cache[key]; !ok {
		m.Cache[key] = m.Func(digit, blinks)
	}
	return m.Cache[key]
}
type Key struct {
	Digit  int
	Blinks int
}
type Memoized struct {
	Func     func(int, int) int
	Cache map[Key]int
}


func Split(num int) (left, right int) {
	length := len(utils.Itoa(num))
	left = utils.Atoi(utils.Itoa(num)[:length/2])
	right = utils.Atoi(utils.Itoa(num)[length/2:])
	return
}

func ParseStones(input string) []int {
	return utils.MapAtoi(utils.Split(input, " "))
}
func StoneF(digit int, blinks int) (amount int) {
	if blinks == 0 {
		return 1
	}
	if digit == 0 {
		return memory.Call(1, blinks-1)
	}
	if len(utils.Itoa(digit))%2 != 0 {
		return memory.Call(digit*2024, blinks-1)
	} else {
		left, right := Split(digit)
		return memory.Call(left, blinks-1) + StoneF(right, blinks-1)
	}
}

func main() {
	session := os.Getenv("AOC_SESSION")
	input := utils.ReadHTTP(2024, 11, session)

	log.Info("--- Part One ---")
	log.Info("", "Result", part1(input))
	log.Info("--- Part Two ---")
	log.Info("", "Result", part2(input))

	os.Exit(0)
}
// part one
func part1(input string) int {
	stones := ParseStones(input)

	memory = Memoized{
		Func: StoneF,
		Cache: make(map[Key]int),
	}
	return utils.Reduce(stones, func(acc int, stone int) int {
		return acc + memory.Call(stone, 25)
	}, 0)
}

// part two
func part2(input string) int {
	stones := ParseStones(input)
	memory = Memoized{
		Func: StoneF,
		Cache: make(map[Key]int),
	}
	return utils.Reduce(stones, func(acc int, stone int) int {
		return acc + StoneF(stone, 75)
	}, 0)
}
