package day4

import (
	"strconv"
	"strings"
)

type Stack []rune

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str rune) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 'a', false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}
func (s *Stack) Peek() (rune, bool) {
	if s.IsEmpty() {
		return 'a', false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		return element, true
	}

}
func (s *Stack) Reverse() {

	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func Day5a(input string) string {
	instructions := strings.Split(strings.Split(input, "\n\n")[1], "\n")
	crates := strings.Split(strings.Split(input, "\n\n")[0], "\n")

	length := (len(crates[len(crates)-1]) + 1)
	height := len(crates) - 1
	cratesSplit := make([]Stack, length/4)
	for i := 0; i < height; i += 1 {
		line := []rune(crates[i])
		for j := 1; j < length; j += 4 {
			cratesSplit[j/4].Push(line[j])
		}

	}

	for _, e := range cratesSplit {
		e.Reverse()
	}
	for i := range instructions {
		if instructions[i] == "" {
			continue
		}
		splitInstr := strings.Split(instructions[i], " ")
		depth, _ := strconv.Atoi(splitInstr[1])

		from, _ := strconv.Atoi(splitInstr[3])

		to, _ := strconv.Atoi(splitInstr[5])
		for i := 0; i < depth; {
			element, _ := cratesSplit[from-1].Pop()
			if element != 32 {
				cratesSplit[to-1].Push(element)
				i += 1
			}
		}
	}
	answer := make([]rune, 0)

	for _, e := range cratesSplit {

		top, _ := e.Peek()
		answer = append(answer, top)
	}
	return string(answer)
}

func Day5b(input string) string {
	instructions := strings.Split(strings.Split(input, "\n\n")[1], "\n")
	crates := strings.Split(strings.Split(input, "\n\n")[0], "\n")

	length := (len(crates[len(crates)-1]) + 1)
	height := len(crates) - 1
	cratesSplit := make([]Stack, length/4)
	for i := 0; i < height; i += 1 {
		line := []rune(crates[i])
		for j := 1; j < length; j += 4 {
			cratesSplit[j/4].Push(line[j])
		}

	}

	for _, e := range cratesSplit {
		e.Reverse()
	}
	for i := range instructions {
		if instructions[i] == "" {
			continue
		}
		splitInstr := strings.Split(instructions[i], " ")
		depth, _ := strconv.Atoi(splitInstr[1])

		from, _ := strconv.Atoi(splitInstr[3])

		to, _ := strconv.Atoi(splitInstr[5])
		intStack := Stack{}
		for i := 0; i < depth; {
			element, _ := cratesSplit[from-1].Pop()
			if element != 32 {
				intStack.Push(element)
				i += 1
			}
		}
		for i := 0; i < depth; i++ {
			element, _ := intStack.Pop()
			cratesSplit[to-1].Push(element)
		}

	}
	answer := make([]rune, 0)

	for _, e := range cratesSplit {

		top, _ := e.Peek()
		answer = append(answer, top)
	}
	return string(answer)
}
