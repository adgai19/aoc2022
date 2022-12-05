package day4

import (
	"strconv"
	"strings"

	"github.com/adgai19/aoc2022/utils"
)

func createCratePile(input string) []utils.Stack[rune] {
	crates := strings.Split(strings.Split(input, "\n\n")[0], "\n")

	length := (len(crates[len(crates)-1]) + 1)
	height := len(crates) - 1
	cratesSplit := make([]utils.Stack[rune], length/4)

	for i := 0; i < height; i += 1 {
		line := []rune(crates[i])
		for j := 1; j < length; j += 4 {
			cratesSplit[j/4].Push(line[j])
		}
	}

	for _, e := range cratesSplit {
		e.Reverse()
	}

	return cratesSplit
}

func getInstr(line string) (int, int, int) {
	splitInstr := strings.Split(line, " ")
	depth, _ := strconv.Atoi(splitInstr[1])
	from, _ := strconv.Atoi(splitInstr[3])
	to, _ := strconv.Atoi(splitInstr[5])
	return depth, from, to
}
func getTopLine(crate []utils.Stack[rune]) string {
	answer := make([]rune, 0)

	for _, e := range crate {
		top, _ := e.Peek()
		answer = append(answer, top)
	}
	return string(answer)
}

func Day5a(input string) string {
	instructions := strings.Split(strings.Split(input, "\n\n")[1], "\n")
	cratePile := createCratePile(input)

	for i := range instructions {

		depth, from, to := getInstr(instructions[i])
		for i := 0; i < depth; {
			element, _ := cratePile[from-1].Pop()
			if element != 32 {
				cratePile[to-1].Push(element)
				i += 1
			}
		}
	}

	return getTopLine(cratePile)

}

func Day5b(input string) string {
	instructions := strings.Split(strings.Split(input, "\n\n")[1], "\n")
	cratePile := createCratePile(input)

	for i := range instructions {

		depth, from, to := getInstr(instructions[i])
		intStack := utils.Stack[rune]{}

		for i := 0; i < depth; {
			element, _ := cratePile[from-1].Pop()
			if element != 32 {
				intStack.Push(element)
				i += 1
			}
		}

		for i := 0; i < depth; i++ {
			element, _ := intStack.Pop()
			cratePile[to-1].Push(element)
		}

	}
	return getTopLine(cratePile)
}
