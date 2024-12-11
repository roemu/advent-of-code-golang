package utils

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/quartercastle/vector"
)

func RemoveDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
func Any[T any](array []T, predicate func(item T) bool) bool {
	for _, s := range array {
		if predicate(s) {
			return true
		}
	}
	return false
}
func All[T any](array []T, predicate func(item T) bool) bool {
	for _, s := range array {
		if !predicate(s) {
			return false
		}
	}
	return true
}
func Filter[T any](array []T, predicate func(item T) bool) (ret []T) {
	for _, s := range array {
		if predicate(s) {
			ret = append(ret, s)
		}
	}
	return
}
func Reorder[T any](array []T, a int, b int) []T {
	toInsert := array[a]
	array = slices.Delete(array, a, a+1)
	array = slices.Insert(array, b, toInsert)
	return array
}
func Swap[T any](array []T, a int, b int) []T {
	array[a], array[b] =  array[b], array[a]
	return array

}
func MapI[T, U any](ts []T, f func(T, int) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i], i)
	}
	return us
}
func SwapElements[T any](arr []T, from, to, length int) []T {
	fromElements := make([]T, length)
	copy(fromElements, arr[from:from+length])
	toElements := make([]T, length)
	copy(toElements, arr[to:to+length])
	for i := range length {
		arr[to+i] = fromElements[i]
		arr[from+i] = toElements[i]
	}
	return arr
}

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}
func MapAtoi64(ts []string) []int64 {
	us := make([]int64, len(ts))
	for i := range ts {
		num, err := strconv.ParseInt(ts[i], 10, 64)
		if err != nil {
			log.Fatal("Couldn't convert string to number", "string", ts[i], "input", ts)
		}
		us[i] = num
	}
	return us
}
func MapAtoi(ts []string) []int {
	us := make([]int, len(ts))
	for i := range ts {
		us[i] = Atoi(ts[i])
	}
	return us
}
func MapItoa(ts []int) []string {
	us := make([]string, len(ts))
	for i := range ts {
		num := strconv.Itoa(ts[i])
		us[i] = num
	}
	return us
}

func Atoi64(input string) int64 {
	input = strings.TrimSpace(input)
	num, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		log.Fatal("Couldn't convert string to number", "input", input)
	}
	return num
}
func Atoi(input string) int {
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("Couldn't convert string to number", "input", input)
	}
	return num
}
func Itoa64(input int64) string {
	return strconv.FormatInt(input, 10)
}
func Itoa(input int) string {
	return strconv.Itoa(input)
}

func Reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
    acc := initValue
    for _, v := range s {
        acc = f(acc, v)
    }
    return acc
}
func ReduceI[T, M any](s []T, f func(acc M, item T, index int) M, initValue M) M {
    acc := initValue
    for i, v := range s {
        acc = f(acc, v, i)
    }
    return acc
}

func Split(input string, separator string) []string {
	splits := strings.Split(input, separator)
	return FilterEmpty(splits)
}

func SplitLeftRight(input string, separator string) (string, string) {
	splits := strings.Split(input, separator)
	if len(splits) > 2 {
		log.Warn("SplitLeftRight resulted in more than 2 results", "result", splits)
	}
	return splits[0], splits[1]
}


func FilterEmpty(input []string) []string {
	return Filter(input, func(str string) bool {
		return str != "" && str != " " && len(strings.TrimSpace(str)) > 0
	})
}

func ContainsVector(input []vector.Vector, elem vector.Vector) bool {
	for _, vec := range input {
		if vec.Equal(elem) {
			return true
		}
	}
	return false
}

func RotateVector(vec vector.Vector, degrees int) vector.Vector {
	rotated := vec.Rotate(float64(degrees) * (math.Pi / 180))
	return vector.Vector{math.Round(rotated.X()), math.Round(rotated.Y())}
}
func Sum64(input []int64) int64 {
	return Reduce(input, func(acc, num int64) int64 {
		return acc + num
	},0)
}
func Product64(input []int64) int64 {
	return Reduce(input, func(acc, num int64) int64 {
		return acc * num
	},0)
}
func Sum(input []int) int {
	return Reduce(input, func(acc, num int) int {
		return acc + num
	},0)
}
func Product(input []int) int {
	return Reduce(input, func(acc, num int) int {
		return acc * num
	},0)
}
