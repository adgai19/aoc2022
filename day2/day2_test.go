package day2

import (
	"testing"

	"github.com/adgai19/aoc2022/utils"
)

func getDummyInput() string {
	return `A Y
B X
C Z`
}

func TestPart1(t *testing.T) {
	t.Run("given test", func(t *testing.T) {
		input := getDummyInput()
		want := 15
		got := Day2a(input)
		utils.Assert(t, want, got)
	})
	t.Run("final output", func(t *testing.T) {
		input := utils.ReadInput()
		want := 10404
		got := Day2a(input)
		utils.Assert(t, want, got)
	})

}

func TestPart2(t *testing.T) {
	t.Run("given test", func(t *testing.T) {
		input := getDummyInput()
		want := 12
		got := Day2b(input)
		utils.Assert(t, want, got)
	})
	t.Run("final output", func(t *testing.T) {
		input := utils.ReadInput()
		want := 10334
		got := Day2b(input)
		utils.Assert(t, want, got)
	})
}
