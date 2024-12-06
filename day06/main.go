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

func tickGuard(input input, g guard) (guard, bool) {
	nextPosition := [2]int{
		g.position[0] + g.direction[0],
		g.position[1] + g.direction[1],
	}
	if nextPosition[0] < 0 ||
		nextPosition[1] < 0 ||
		nextPosition[0] >= len(input[0]) ||
		nextPosition[1] >= len(input) {
		return guard{nextPosition, g.direction}, false
	}
	if input[nextPosition[1]][nextPosition[0]] == '#' {
		nextDirection := [2]int{-g.direction[1], g.direction[0]}
		return guard{g.position, nextDirection}, true
	} else {
		return guard{nextPosition, g.direction}, true
	}
}

func willLoop(input input) bool {
	guardPosition := getGuardPosition(input)
	guardDirection := [2]int{0, -1}

	g := guard{guardPosition, guardDirection}
	ok := true

	visited := make(map[guard]bool)
	for ok {
		g, ok = tickGuard(input, g)
		if visited[g] {
			return true
		}
		visited[g] = true
	}
	return false
}

func getGuardPosition(input input) [2]int {
	for y, line := range input {
		for x, c := range line {
			if c == '^' {
				return [2]int{x, y}
			}
		}
	}
	return [2]int{-1, -1}
}

func part1(input input) int {
	guardPosition := getGuardPosition(input)
	guardDirection := [2]int{0, -1}

	g := guard{guardPosition, guardDirection}
	ok := true

	visited := make(map[[2]int]bool)
	for ok {
		visited[g.position] = true
		g, ok = tickGuard(input, g)
	}
	return len(visited)
}

func part2(input input) int {
	guardPosition := getGuardPosition(input)

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
