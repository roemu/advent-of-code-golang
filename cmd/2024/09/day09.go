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
	disk := []string{}
	fileId := 0
	for i, char := range utils.Split(input, "") {
		if i == 0 || i%2 == 0 {
			for j := 0; j < utils.Atoi(char); j++ {
				disk = append(disk, utils.Itoa(fileId))
			}
			fileId++
		} else {
			for j := 0; j < utils.Atoi(char); j++ {
				disk = append(disk, ".")
			}
		}
	}
	handledFileIds := []string{}
	for {
		start, length, char := GetLastFileRange(disk, &handledFileIds)
		if char == "" {
			break
		}
		dotStart, found := FindSuitableDotRange(disk, length)
		if !found || dotStart > start {
			continue
		}
		disk = utils.SwapElements(disk, start, dotStart, length)
	}
	return utils.ReduceI(disk, func(acc int, char string, index int) int {
		if char == "." {
			return acc
		}
		return acc + (index * utils.Atoi(char))
	}, 0)

}

func GetLastFileRange(disk []string, seen *[]string) (start, length int, char string) {
	start = -1
	for i := len(disk) - 1; i > 0; i-- {
		char = disk[i]
		if slices.Contains(*seen, char) || char == "." {
			continue
		}
		*seen = append(*seen, char)
		start = i
		length = 1
		for {
			i--
			if i < 0 || disk[i] != char {
				return
			}
			length++
			start = i
		}
	}
	char = ""
	return
}

func FindSuitableDotRange(disk []string, length int) (start int, found bool) {
	start = -1
	found = false

	count := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] == "." {
			count++
		}
		if i == len(disk)-1 || disk[i] != "." {
			if count >= length {
				start = i - count
				if i == len(disk)-1 && disk[i] == "." {
					start++
				}
				found = true
				return
			}
			count = 0
		}
	}
	return
}
