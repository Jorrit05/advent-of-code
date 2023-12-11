package lib

import (
	"fmt"
	"strings"
)

func PrettyPrint(s [][]string) {
	for _, line := range s {
		fmt.Println(strings.Join(line, ""))
	}
}
func PrettyPrintInt(s [][]int) {
	for _, line := range s {
		fmt.Println(line)
	}
}

func SliceIncrementRange(start, increment int) []int {
	var rangeSlice []int
	for i := start; i < start+increment; i++ {
		rangeSlice = append(rangeSlice, i)
	}
	return rangeSlice
}

func InitStringSlice(length int, char string) []string {
	var rangeSlice []string
	for i := 0; i < length; i++ {
		rangeSlice = append(rangeSlice, char)
	}
	return rangeSlice
}

func ReverseSlice(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func InsertIntoSliceAtIndex[T any](destination []T, element T, index int) []T {
	if len(destination) == index {
		return append(destination, element)
	}

	destination = append(destination[:index+1], destination[index:]...) // index < len(a)
	destination[index] = element

	return destination
}
