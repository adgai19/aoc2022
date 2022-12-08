package day4

import (
	"fmt"

	"github.com/adgai19/aoc2022/utils"
)

func Day6a(input string) int {
	splitS := []rune(input)
	charq := make(utils.Queue[rune], 0)
	for i, x := range splitS {
		if i == 4 {
			break
		}
		charq.Enque(x)
	}
	for i, x := range splitS {
		if i < 4 {
			continue
		}
		charq.Enque(x)
		if uniquea(charq.List()) {
			return i
		}

		charq.Deque()
	}

	return 0
}

func Day6b(input string) int {
	splitS := []rune(input)
	charq := make(utils.Queue[rune], 0)
	for i, x := range splitS {
		if i == 14 {
			break
		}
		charq.Enque(x)
	}
	for i, x := range splitS {
		if i < 14 {
			continue
		}
		charq.Enque(x)
		if uniquea(charq.List()) {
			return i
		}

		charq.Deque()
	}

	return 0
}
func contains(s []rune, e rune) bool {
	fmt.Println("comparing", s, e)
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func uniquea(s []rune) bool {

	var unique []rune
	m := map[rune]bool{}
	for _, v := range s {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return len(unique) == len(s)

}
