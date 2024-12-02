package main

import (
	"aoc/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestSafeReport(t *testing.T) {
	assert := assert.New(t)

	report := NewReport("1 1 2 3 4 5")
	assert.True(report.IsSafeIgnoreOne(), report.Raw)

	report = NewReport("7 6 4 2 1")
	assert.True(report.IsSafe(), report.Raw)

	report = NewReport("1 3 6 7 9")
	assert.True(report.IsSafe(), report.Raw)

	report = NewReport("8 6 4 4 1")
	assert.True(report.IsSafeIgnoreOne(), report.Raw)

	report = NewReport("8 6 4 4 1")
	assert.True(report.IsSafeIgnoreOne(), report.Raw)
}
func TestSafeReportOrder(t *testing.T) {
	assert := assert.New(t)

	report := NewReport("1 3 2 4 5")
	assert.False(report.IsSafe(), report.Raw)

	report = NewReport("1 3 2 4 5")
	assert.True(report.IsSafeIgnoreOne(), report.Raw)
}
func TestUnsafeReport(t *testing.T) {
	assert := assert.New(t)

	report := NewReport("1 2 7 8 9")
	assert.False(report.IsSafe(), report.Raw)

	report = NewReport("1 2 7 8 9")
	assert.False(report.IsSafeIgnoreOne(), report.Raw)

	report = NewReport("9 7 6 2 1")
	assert.False(report.IsSafe(), report.Raw)

		report = NewReport("8 6 4 4 1")
	assert.False(report.IsSafe(), report.Raw)
}
func TestReportSize(t *testing.T) {
	assert := assert.New(t)
	report := NewReport("1 2 7 8 9")

	assert.Len(report.Numbers, 5)
}
func TestReportIsConsecutive(t *testing.T) {
	assert := assert.New(t)

	report := NewReport("1 2 7 8 9")
	assert.True(IsConsecutive(report.Numbers), report.Raw)

	report = NewReport("5 2 7 8 9")
	assert.False(IsConsecutive(report.Numbers), report.Raw)

	report = NewReport("-1 2 7 5 9")
	assert.False(IsConsecutive(report.Numbers), report.Raw)
}
func TestSafeDistances(t *testing.T) {
	assert := assert.New(t)

	report := NewReport("1 2 7 8 9")
	assert.False(SafeDistances(report.Numbers), report.Raw)

	report = NewReport("1 2 7 9")
	assert.False(SafeDistances(report.Numbers), report.Raw)

	report = NewReport("3 5 7 9")
	assert.True(SafeDistances(report.Numbers), report.Raw)

	report = NewReport("3 -1 7 9")
	assert.False(SafeDistances(report.Numbers), report.Raw)

}
func TestDataReportLength(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("./day02.txt")
	data := NewData(input)

	assert.Len(data.Reports, 10)
}
func TestDataSafeReports(t *testing.T) {
	// assert := assert.New(t)
	// input := utils.ReadFile("./day02.txt")
	// data := NewData(input)
	//
	// assert.Equal(data.SafeReports(), 2)
}
func TestDataSafeReportsIgnoreOne(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("./day02.txt")
	data := NewData(input)

	assert.Equal(10, data.SafeReportsIgnoreOne())
}

func TestDay02(t *testing.T) {
	assert := assert.New(t)
	input := utils.ReadFile("./day02.txt")

	// t.Run("part 1", func(t *testing.T) {
	// 	expected := 2
	// 	actual := part1(input)
	//
	// 	assert.Equal(expected, actual)
	// })

	t.Run("part 2", func(t *testing.T) {
		expected := 10
		actual := part2(input)

		assert.Equal(expected, actual)
	})
}
