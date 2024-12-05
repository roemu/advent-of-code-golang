package main

import (
	"aoc/internal/utils"
	"slices"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

type Input struct {
	Raw        string
	AfterRules map[int][]int
	Updates    []Update
}

func (input *Input) SumMiddlePages() int {
	return utils.Reduce(input.Updates, func(acc int, update Update) int {
		if !update.IsValidOrder(input) {
			return acc
		}
		return acc + update.MiddlePage()
	}, 0)
}
func (input *Input) SumdMiddlePagesCorrect() int {
	return utils.Reduce(input.Updates, func(acc int, update Update) int {
		if update.IsValidOrder(input) {
			return acc
		}
		update.FixOrder(input)
		return acc + update.MiddlePage()
	}, 0)
}
func NewInput(input string) *Input {
	split := strings.Split(input, "\n\n")
	if len(split) != 2 {
		log.Fatal("Split was not of length 2", "split len", len(split), "input", input)
	}
	orderRules := utils.Split(split[0], "\n")
	afterRules := make(map[int][]int, len(orderRules))
	for _, orderRule := range orderRules {
		left, right := utils.SplitLeftRight(orderRule, "|")
		prev, present := afterRules[utils.Atoi(left)]
		if present {
			prev = append(prev, utils.Atoi(right))
			afterRules[utils.Atoi(left)] = prev
			continue
		}
		afterRules[utils.Atoi(left)] = []int{utils.Atoi(right)}
	}
	updates := utils.Split(split[1], "\n")
	return &Input{
		Raw: input,
		Updates: utils.Map(updates, func(line string, _ int) Update {
			return *NewUpdate(line)
		}),
		AfterRules: afterRules,
	}
}

type Update struct {
	Raw             string
	Pages           []int
	AlreadyAppeared map[int]int
}

func (update *Update) IsValidOrder(input *Input) bool {
	for i, page := range update.Pages {
		if _, ok := update.AlreadyAppeared[page]; ok { // Page already appeared once
			return false
		}
		update.AlreadyAppeared[page] = i
		for key, value := range input.AfterRules {
			if key != page {
				continue
			}
			for _, mustAfter := range value {
				if _, before := update.AlreadyAppeared[mustAfter]; before {
					return false
				}
			}
		}
	}
	return true
}
func (update *Update) MiddlePage() int {
	numPages := len(update.Pages)
	if numPages%2 == 0 {
		log.Fatal("Number of pages was not odd, can't extract middle page", "numPages", numPages, "update.Pages", update.Pages)
	}
	return update.Pages[numPages/2]
}
func (update *Update) FixOrder(input *Input) {
	result := []int{}
	visitedMap := make(map[int]bool, len(update.Pages))

	for _, page := range update.Pages {
		if _, visited := visitedMap[page]; !visited {
			TopoSort(page, input.AfterRules, &visitedMap, &result, update.Pages)
		}
	}

	slices.Reverse(result)
	update.Pages = result
}

func TopoSort(node int, afterMap map[int][]int, visitedMap *map[int]bool, result *[]int, allNodes []int) {
	(*visitedMap)[node] = true
	afters, someAfter := afterMap[node]
	if someAfter {
		for _, after := range afters {
			if !slices.Contains(allNodes, after) {
				continue
			}
			if _, visited := (*visitedMap)[after]; !visited {
				TopoSort(after, afterMap, visitedMap, result, allNodes)
			}
		}
	}
	*result = append(*result, node)

}
func NewUpdate(input string) *Update {
	pages := strings.Split(input, ",")
	return &Update{
		Raw: input,
		Pages: utils.Map(pages, func(str string, _ int) int {
			num, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal("Couldn't convert string to number", "string", str, "input", input)
			}
			return num
		}),
		AlreadyAppeared: make(map[int]int, len(pages)),
	}
}
