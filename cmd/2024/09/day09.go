package main

import (
	"os"
	"slices"

	"aoc/internal/utils"

	"github.com/charmbracelet/log"
)

func main() {
	session := os.Getenv("AOC_SESSION")
	input := utils.ReadHTTP(2024, 9, session)

	log.Info("--- Part One ---")
	log.Info("", "Result", part1(input))
	log.Info("--- Part Two ---")
	log.Info("", "Result", part2(input))

	os.Exit(0)
}

// part one
func part1(input string) int {
	output := []string{}
	fileId := 0
	for i, char := range utils.Split(input, "") {
		if i == 0 || i%2 == 0 {
			for j := 0; j < utils.Atoi(char); j++ {
				output = append(output, utils.Itoa(fileId))
			}
			fileId++
		} else {
			for j := 0; j < utils.Atoi(char); j++ {
				output = append(output, ".")
			}
		}
	}

	fileRef := utils.Filter(output, func(char string) bool {
		return char != "."
	})
	fileIndex := len(fileRef) - 1
	output = utils.MapI(output, func(char string, _ int) string {
		if char == "." && fileIndex == 0 {
			return "."
		}
		if char == "." {
			fileIndex--
			return fileRef[fileIndex+1]
		}
		return char
	})

	toDelete := len(fileRef) - fileIndex - 1
	output = output[:len(output)-toDelete]
	for i := 0; i < toDelete; i++ {
		output = append(output, ".")
	}
	return utils.ReduceI(output, func(acc int, char string, index int) int {
		if char == "." {
			return acc
		}
		return acc + (index * utils.Atoi(char))
	}, 0)
}

// part two
func part2(input string) int {
	blocks := Start(input)
	return CalculateChecksum(blocks)
}

type Block struct {
	Id     int
	Length int
}

func (block *Block) String() string {
	out := ""
	for range block.Length {
		if block.Id == -1 {
			out += "."
			continue
		}
		out += utils.Itoa(block.Id)
	}
	return out
}
func ReadableBlocks(blocks []Block) string {
	out := ""
	for _, block := range blocks {
		out += block.String()
	}
	return out
}

func GenerateBlocks(input string) []Block {
	disk := []Block{}
	fileId := 0
	for i, char := range utils.Split(input, "") {
		block := &Block{
			Id:     fileId,
			Length: utils.Atoi(char),
		}
		if i == 0 || i%2 == 0 {
			fileId++
		} else {
			block.Id = -1
		}
		disk = append(disk, *block)
	}
	return utils.Filter(disk, func(block Block) bool {
		return block.Length != 0
	})
}

func CondenseBlocks(blocks []Block) []Block {
	for i := 1; i < len(blocks); i++ {
		if blocks[i-1].Id == -1 && blocks[i].Id == -1 {
			blocks[i].Length += blocks[i-1].Length
			blocks[i-1].Length = 0
		}
	}
	return utils.Filter(blocks, func(block Block) bool {
		return block.Length != 0
	})
}

func Start(input string) []Block {
	blocks := GenerateBlocks(input)
	nextBlock := RevBlockIter()
	for {
		filePos := nextBlock(blocks)
		if filePos < 0 {
			break
		}
		emptyPos := FindFirstEmptyBlock(blocks, blocks[filePos].Length)
		if emptyPos < 0 || emptyPos > filePos {
			continue
		}
		diff := blocks[emptyPos].Length - blocks[filePos].Length
		blocks[emptyPos].Id, blocks[filePos].Id = blocks[filePos].Id, blocks[emptyPos].Id
		blocks[emptyPos].Length = blocks[filePos].Length
		if diff > 0 {
			blocks = slices.Insert(blocks, emptyPos+1, Block{
				Id:     -1,
				Length: diff,
			})
		}

		blocks = CondenseBlocks(blocks)
	}
	return blocks
}

func CalculateChecksum(blocks []Block) int {
	index := 0
	sum := 0
	for _, block := range blocks {
		if block.Id == -1 {
			index += block.Length
			continue
		}
		for range block.Length {
			sum += block.Id * index
			index++
		}
	}
	return sum
}

func FindFirstEmptyBlock(blocks []Block, minLength int) int {
	for i, block := range blocks {
		if block.Id == -1 && block.Length >= minLength {
			return i
		}
	}
	return -1
}

func RevBlockIter() func([]Block) int {
	seenBlockIds := []int{}
	return func(blocks []Block) int {
		for i := len(blocks) - 1; i > 0; i-- {
			if blocks[i].Id == -1 {
				continue
			}
			if !slices.Contains(seenBlockIds, blocks[i].Id) {
				seenBlockIds = append(seenBlockIds, blocks[i].Id)
				return i
			}
		}
		return -1
	}
}
