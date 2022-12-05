package day3

import (
	"strings"

	"github.com/adgai19/aoc2022/utils"
)

func Day3a(input string) int {
	lines := utils.GetLines(input)
	common := make(map[string]string)

	for _, line := range lines {
		halfs := splitByString(line, len(line)/2)
		common[line] = findCommonString(halfs[0], halfs[1])
	}

	priority := 0
	for _, item := range common {
		priority += calculatePriority(item)
	}
	return priority
}

func Day3b(input string) int {
	lines := utils.GetLines(input)
	var badges []string
	splitLines := splitByArray(lines, 3)

	for _, sl := range splitLines {
		badges = append(badges, findCommonString(sl[0], findCommonString(sl[1], sl[2])))
	}

	priority := 0
	for _, badge := range badges {
		priority += calculatePriority(badge)
	}
	return priority
}
func splitByArray(array []string, n int) [][]string {

	var result [][]string
	var temp []string

	for x := 0; x < len(array); x = x + 1 {
		if x%3 == 0 && x != 0 {
			result = append(result, temp)
			temp = nil
		}
		temp = append(temp, array[x])
	}
	result = append(result, temp)

	return result

}

func splitByString(s string, n int) []string {
	var ss []string

	for i := 1; i < len(s); i++ {
		if i%n == 0 {
			ss = append(ss, s[:i])
			s = s[i:]
			i = 1
		}
	}

	ss = append(ss, s)
	return ss
}

func calculatePriority(item string) int {
	priority := 0

	for _, char := range item {
		p := int(char)
		if p > 97 {
			priority += p - 96
		} else {
			priority += p - 64 + 26
		}
	}

	return priority

}

func findCommonString(a, b string) string {
	if a == "" || b == "" {
		return ""
	}

	var result string
	for _, char := range a {
		if strings.Contains(b, string(char)) {
			if !strings.Contains(result, string(char)) {
				result = result + string(char)
			}
		}
	}
	return result
}
