package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
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
	g := getGuard(input)
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

func getGuard(input input) guard {
	for y, line := range input {
		for x, c := range line {
			if c == '^' {
				return guard{
					[2]int{x, y},
					[2]int{0, -1},
				}
			}
		}
	}
	return guard{
		[2]int{-1, -1},
		[2]int{0, 0},
	}
}

func part1(input input) int {
	g := getGuard(input)
	ok := true

	visited := make(map[[2]int]bool)
	for ok {
		visited[g.position] = true
		g, ok = tickGuard(input, g)
	}
	return len(visited)
}

func addingObstacleWillLoop(input input, position [2]int) bool {
	inputCopy := make([][]rune, len(input))
	copy(inputCopy, input)
	for i := range input {
		inputCopy[i] = make([]rune, len(input[i]))
		copy(inputCopy[i], input[i])
	}
	inputCopy[position[1]][position[0]] = '#'
	return willLoop(inputCopy)
}

func part2(input input) int {
	g := getGuard(input)
	sum := 0

	var wg sync.WaitGroup
	var mu sync.Mutex
	for y, line := range input {
		for x, c := range line {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if [2]int{x, y} == g.position || c == '#' {
					return
				}
				if addingObstacleWillLoop(input, [2]int{x, y}) {
					mu.Lock()
					sum++
					mu.Unlock()
				}
			}()
		}
	}
	wg.Wait()

	return sum
}
