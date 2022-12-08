package day4

import (
	"testing"

	"github.com/adgai19/aoc2022/utils"
)

func TestPart1(t *testing.T) {
	t.Run("given test", func(t *testing.T) {
		input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
		want := 7
		got := Day6a(input)
		utils.Assert(t, got, want)
	})

	t.Run("given test2", func(t *testing.T) {
		input := "bvwbjplbgvbhsrlpgdmjqwftvncz"
		want := 5
		got := Day6a(input)
		utils.Assert(t, got, want)
	})

	t.Run("given test3", func(t *testing.T) {
		input := "nppdvjthqldpwncqszvftbrmjlhg"
		want := 6
		got := Day6a(input)
		utils.Assert(t, got, want)
	})

	t.Run("final output", func(t *testing.T) {
		input := utils.ReadInput("input.txt")
		want := 1794
		got := Day6a(input)
		utils.Assert(t, want, got)
	})

}

func TestPart2(t *testing.T) {

	t.Run("given test2", func(t *testing.T) {
		input := "bvwbjplbgvbhsrlpgdmjqwftvncz"
		want := 23
		got := Day6b(input)
		utils.Assert(t, got, want)
	})

	t.Run("given test3", func(t *testing.T) {
		input := "nppdvjthqldpwncqszvftbrmjlhg"
		want := 23
		got := Day6b(input)
		utils.Assert(t, got, want)
	})
	t.Run("final output", func(t *testing.T) {
		input := utils.ReadInput("input.txt")
		want := 2851
		got := Day6b(input)
		utils.Assert(t, want, got)
	})
}
