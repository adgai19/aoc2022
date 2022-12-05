package day2

import (
	"strings"

	"github.com/adgai19/aoc2022/utils"
)

const (
	theyPlayRock     = "A"
	theyPlayPaper    = "B"
	theyPlayScissors = "C"

	youPlayRock     = "X"
	youPlayPaper    = "Y"
	youPlayScissors = "Z"
)

var exactScore = map[string]map[string]int{
	theyPlayRock: {
		youPlayRock:     4,
		youPlayPaper:    8,
		youPlayScissors: 3,
	},
	theyPlayPaper: {
		youPlayRock:     1,
		youPlayPaper:    5,
		youPlayScissors: 9,
	},
	theyPlayScissors: {
		youPlayRock:     7,
		youPlayPaper:    2,
		youPlayScissors: 6,
	},
}

var forDesiredOutcome = map[string]map[string]string{
	theyPlayRock: {
		"X": youPlayScissors,
		"Y": youPlayRock,
		"Z": youPlayPaper,
	},
	theyPlayPaper: {
		"X": youPlayRock,
		"Y": youPlayPaper,
		"Z": youPlayScissors,
	},
	theyPlayScissors: {
		"X": youPlayPaper,
		"Y": youPlayScissors,
		"Z": youPlayRock,
	},
}

func Day2a(input string) int {
	lines := utils.GetLines(input)
	score := 0
	for _, line := range lines {
		sel := strings.Fields(line)
		score += exactScore[sel[0]][sel[1]]
	}
	return score

}
func Day2b(input string) int {
	lines := utils.GetLines(input)
	score := 0
	for _, line := range lines {
		sel := strings.Fields(line)
		score += exactScore[sel[0]][forDesiredOutcome[sel[0]][sel[1]]]
	}
	return score

}
