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
func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}
func MapAtoi(ts []string) []int {
	us := make([]int, len(ts))
	for i := range ts {
		num, err := strconv.Atoi(ts[i])
		if err != nil {
			log.Fatal("Couldn't convert string to number", "string", ts[i], "input", ts)
		}
		us[i] = num
	}
	return us
}

func Atoi(input string) int {
	num, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("Couldn't convert string to number", "input", input)
	}
	return num
}

func Reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
    acc := initValue
    for _, v := range s {
        acc = f(acc, v)
    }
    return acc
}

func Split(input string, separator string) []string {
	splits := strings.Split(input, separator)
	return Filter(splits, func(split string) bool {
		return split != ""
	})
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
		return str != "" && str != " "
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
