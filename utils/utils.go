package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadInput() string {
	file, _ := ioutil.ReadFile("input.txt")
	input := string(file)
	return input
}
func SplitInt(input string) []int {
	split := strings.Split(input, ",")
	out := make([]int, 0)
	for _, val := range split {
		i, _ := strconv.Atoi(val)
		out = append(out, i)
	}
	return out

}
