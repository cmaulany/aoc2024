package main

import (
	"bufio"
	"fmt"
	"os"
)

type input = [][]rune

func main() {
	input := load()

	answerPart1 := part1(input)
	fmt.Printf("Answer part 1: %d\n", answerPart1)

	answerPart2 := part2(input)
	fmt.Printf("Answer part 2: %d\n", answerPart2)
}

func load() input {
	file, _ := os.Open("inputs/day04.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input [][]rune
	for scanner.Scan() {
		line := []rune(scanner.Text())
		input = append(input, line)
	}
	return input
}

func checkWord(input input, word string, pos [2]int, dir [2]int) bool {
	for i, c := range word {
		x := pos[0] + dir[0]*i
		y := pos[1] + dir[1]*i
		if x < 0 || x >= len(input[0]) || y < 0 || y >= len(input) {
			return false
		}
		if input[y][x] != c {
			return false
		}
	}
	return true
}

func checkXMas(input input, pos [2]int) bool {
	x := pos[0]
	y := pos[1]
	if !checkWord(input, "MAS", [2]int{x - 1, y - 1}, [2]int{1, 1}) &&
		!checkWord(input, "MAS", [2]int{x + 1, y + 1}, [2]int{-1, -1}) {
		return false
	}
	if !checkWord(input, "MAS", [2]int{x - 1, y + 1}, [2]int{1, -1}) &&
		!checkWord(input, "MAS", [2]int{x + 1, y - 1}, [2]int{-1, 1}) {
		return false
	}
	return true
}

func part1(input input) int {
	dirs := [8][2]int{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}
	sum := 0
	for y, line := range input {
		for x := range line {
			for _, dir := range dirs {
				if checkWord(input, "XMAS", [2]int{x, y}, dir) {
					sum++
				}
			}
		}
	}
	return sum
}

func part2(input input) int {
	sum := 0
	for y, line := range input {
		for x := range line {
			if checkXMas(input, [2]int{x - 1, y - 1}) {
				sum++
			}
		}
	}
	return sum
}
