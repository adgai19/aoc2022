package day1

import (
	"sort"
	"strconv"
	"strings"
)

func common(input string) []int {
	splitStr := strings.Split(input, "\n\n")
	array := make([]int, 0)
	for _, element := range splitStr {
		sum := 0
		for _, e := range strings.Split(element, "\n") {
			int, _ := strconv.Atoi(e)
			sum = sum + int
		}
		array = append(array, sum)
		sort.Ints(array)
		sum = 0
	}
	return array

}

func Day1a(input string) int {
	out := common(input)
	return out[len(out)-1]

}
func Day1b(input string) int {
	out := common(input)
	return out[len(out)-1] + out[len(out)-2] + out[len(out)-3]
}
