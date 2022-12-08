package day8

import (
	"testing"

	"github.com/adgai19/aoc2022/utils"
)

func TestPart1(t *testing.T) {
	t.Run("given test", func(t *testing.T) {
		input := utils.ReadInputTest()
		want := 21
		got := Day8a(input)
		utils.Assert(t, got, want)
	})

	t.Run("input", func(t *testing.T) {
		input := utils.ReadInput("input.txt")
		want := 1843
		got := Day8a(input)
		utils.Assert(t, got, want)
	})

}

func TestPart2(t *testing.T) {

	t.Run("given test2", func(t *testing.T) {
		input := utils.ReadInputTest()
		want := 8
		got := Day8b(input)
		utils.Assert(t, got, want)
	})

	t.Run("given test2", func(t *testing.T) {
		input := utils.ReadInput("input.txt")
		want := 180000
		got := Day8b(input)
		utils.Assert(t, got, want)
	})
}
