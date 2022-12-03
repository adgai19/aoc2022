package day3

import (
	"testing"

	"github.com/adgai19/aoc2022/utils"
)

func getDummyInput() string {
	return `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
}

func TestPart1(t *testing.T) {
	t.Run("given test", func(t *testing.T) {
		input := getDummyInput()
		want := 157
		got := Day3a(input)
		utils.Assert(t, want, got)
	})
	t.Run("final output", func(t *testing.T) {
		input := utils.ReadInput()
		want := 8240
		got := Day3a(input)
		utils.Assert(t, want, got)
	})

}

func TestPart2(t *testing.T) {
	t.Run("given test", func(t *testing.T) {
		input := getDummyInput()
		want := 70
		got := Day3b(input)
		utils.Assert(t, want, got)
	})
	t.Run("final output", func(t *testing.T) {
		input := utils.ReadInput()
		want := 2587
		got := Day3b(input)
		utils.Assert(t, want, got)
	})
}
