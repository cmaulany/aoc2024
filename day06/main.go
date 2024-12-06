package main

import (
	"bufio"
	"fmt"
	"os"
)

type input = [][]rune

type guard struct {
	position  [2]int
	direction [2]int
}

func main() {
	input := load()

	answerPart1 := part1(input)
	fmt.Printf("Answer part 1: %d\n", answerPart1)

	answerPart2 := part2(input)
	fmt.Printf("Answer part 2: %d\n", answerPart2)
}

func load() input {
	file, _ := os.Open("inputs/day06.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, []rune(line))
	}
	return input
}

func willLoop(input input) bool {
	var guardPosition [2]int
	guardDirection := [2]int{0, -1}

loop:
	for y, line := range input {
		for x, c := range line {
			if c == '^' {
				guardPosition = [2]int{x, y}
				break loop
			}
		}
	}

	visited := make(map[[2][2]int]bool)
	// i := 0
	for true {
		// i++
		state := [2][2]int{guardPosition, guardDirection}
		seen := visited[state]
		if seen {
			return true
		}
		visited[state] = true
		nextPosition := [2]int{
			guardPosition[0] + guardDirection[0],
			guardPosition[1] + guardDirection[1],
		}
		if nextPosition[0] < 0 ||
			nextPosition[1] < 0 ||
			nextPosition[0] >= len(input[0]) ||
			nextPosition[1] >= len(input) {
			return false
		}
		if input[nextPosition[1]][nextPosition[0]] == '#' {
			guardDirection = [2]int{-guardDirection[1], guardDirection[0]}
		} else {
			guardPosition = nextPosition
		}
	}
	return false
}

func part1(input input) int {
	var guardPosition [2]int
	guardDirection := [2]int{0, -1}

loop:
	for y, line := range input {
		for x, c := range line {
			if c == '^' {
				guardPosition = [2]int{x, y}
				break loop
			}
		}
	}

	visited := make(map[[2]int]bool)
	for true {
		visited[guardPosition] = true
		nextPosition := [2]int{
			guardPosition[0] + guardDirection[0],
			guardPosition[1] + guardDirection[1],
		}
		if nextPosition[0] < 0 ||
			nextPosition[1] < 0 ||
			nextPosition[0] >= len(input[0]) ||
			nextPosition[1] >= len(input) {
			break
		}
		if input[nextPosition[1]][nextPosition[0]] == '#' {
			guardDirection = [2]int{-guardDirection[1], guardDirection[0]}
		} else {
			guardPosition = nextPosition
		}
	}
	return len(visited)
}

func part2(input input) int {
	var guardPosition [2]int

loop:
	for y, line := range input {
		for x, c := range line {
			if c == '^' {
				guardPosition = [2]int{x, y}
				break loop
			}
		}
	}

	sum := 0
	for y, line := range input {
		for x, c := range line {
			if [2]int{x, y} == guardPosition || c == '#' {
				continue
			}
			inputCopy := make([][]rune, len(input))
			copy(inputCopy, input)
			for i := range input {
				inputCopy[i] = make([]rune, len(input[i]))
				copy(inputCopy[i], input[i])
			}
			inputCopy[y][x] = '#'
			if willLoop(inputCopy) {
				sum++
			}
		}
	}
	return sum
}
