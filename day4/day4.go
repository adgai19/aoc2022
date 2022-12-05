package day4

import (
	"strconv"
	"strings"

	"github.com/adgai19/aoc2022/utils"
)

func Day4a(input string) int {
	lines := utils.GetLines(input)
	ctr := 0

	for _, line := range lines {
		if line == "" {
			continue
		}
		splits := strings.Split(line, ",")
		x, y := strings.Split(splits[0], "-"), strings.Split(splits[1], "-")
		xleft, _ := strconv.Atoi(x[0])
		xright, _ := strconv.Atoi(x[1])
		yleft, _ := strconv.Atoi(y[0])
		yright, _ := strconv.Atoi(y[1])
		if xleft >= yleft && xright <= yright {
			ctr += 1
		} else if xleft <= yleft && xright >= yright {
			ctr += 1
		}
	}
	return ctr
}
func Day4b(input string) int {
	ctr := 0

	lines := utils.GetLines(input)
	for _, line := range lines {
		if line == "" {
			continue
		}
		splits := strings.Split(line, ",")
		x, y := strings.Split(splits[0], "-"), strings.Split(splits[1], "-")
		xleft, _ := strconv.Atoi(x[0])
		xright, _ := strconv.Atoi(x[1])
		yleft, _ := strconv.Atoi(y[0])
		yright, _ := strconv.Atoi(y[1])
		if xright < yleft && xright < yright {
		} else if yright < xleft && yright < xright {
		} else {
			ctr += 1
		}

	}
	return ctr
}
