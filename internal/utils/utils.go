package utils

import "strings"

func removeDuplicate[T comparable](sliceList []T) []T {
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

func Map[T, U any](ts []T, f func(T, int) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i], i)
	}
	return us
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

func FilterEmpty(input []string) []string {
	return Filter(input, func(str string) bool {
		return str != "" && str != " "
	})
}
