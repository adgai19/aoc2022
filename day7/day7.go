package day4

import (
	"fmt"
	"path"
	"strings"
	"unicode"

	"github.com/adgai19/aoc2022/utils"
)

func Day7a(input string) int {
	cd, fs := "", map[string]int{}
	for _, line := range utils.GetLines(input) {
		if strings.HasPrefix(line, "$ cd") {
			cd = path.Join(cd, strings.Fields(line)[2])
		} else if unicode.IsDigit([]rune(line)[0]) {
			var size int
			var name string
			fmt.Sscanf(line, "%d %s", &size, &name)
			fs[path.Join(cd, name)] = size
		}
	}
	ds := map[string]int{}
	for f, s := range fs {
		for d := path.Dir(f); d != "/"; d = path.Dir(d) {
			ds[d] += s
		}
		ds["/"] += s
	}
	part1 := 0
	for _, s := range ds {
		if s < 100000 {
			part1 += s
		}
	}

	return part1
}

func Day7b(input string) int {
	cd, fs := "", map[string]int{}
	for _, line := range utils.GetLines(input) {
		if strings.HasPrefix(line, "$ cd") {
			cd = path.Join(cd, strings.Fields(line)[2])
		} else if unicode.IsDigit([]rune(line)[0]) {
			var size int
			var name string
			fmt.Sscanf(line, "%d %s", &size, &name)
			fs[path.Join(cd, name)] = size
		}
	}
	ds := map[string]int{}
	for f, s := range fs {
		for d := path.Dir(f); d != "/"; d = path.Dir(d) {
			ds[d] += s
		}
		ds["/"] += s
	}
	part2 := ds["/"]
	for _, s := range ds {
		if s+70000000-ds["/"] >= 30000000 && s < part2 {
			part2 = s
		}
	}
	fmt.Println(part2)
	return part2
}
