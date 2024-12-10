package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type input = [][]int

func main() {
	input := load()

	answerPart1 := part1(input)
	fmt.Printf("Answer part 1: %d\n", answerPart1)

	answerPart2 := part2(input)
	fmt.Printf("Answer part 2: %d\n", answerPart2)
}

func load() input {
	file, _ := os.Open("inputs/day10.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input [][]int
	for scanner.Scan() {
		var line []int
		for _, c := range scanner.Text() {
			n, _ := strconv.Atoi(string(c))
			line = append(line, n)
		}
		input = append(input, line)
	}
	return input
}

func getTrailHeads(input input) [][2]int {
	var heads [][2]int
	for y, line := range input {
		for x, d := range line {
			if d == 0 {
				heads = append(heads, [2]int{x, y})
			}
		}
	}
	return heads
}

func getNeighbors(input input, position [2]int) [][2]int {
	deltas := [4][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	var neighbors [][2]int
	for _, delta := range deltas {
		neighbor := [2]int{position[0] + delta[0], position[1] + delta[1]}
		if neighbor[0] >= 0 &&
			neighbor[0] < len(input[0]) &&
			neighbor[1] >= 0 &&
			neighbor[1] < len(input) {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func getReachablePeaks(input input, head [2]int) [][2]int {
	var open [][2]int
	var closed [][2]int
	open = append(open, head)

	var reachablePeaks [][2]int
	for len(open) > 0 {
		current := open[0]
		open = open[1:]

		currentHeight := input[current[1]][current[0]]
		if currentHeight == 9 {
			reachablePeaks = append(reachablePeaks, current)
			continue
		}

		neighbors := getNeighbors(input, current)
		for _, neighbor := range neighbors {
			neighborHeight := input[neighbor[1]][neighbor[0]]
			if neighborHeight == currentHeight+1 && !slices.Contains(closed, neighbor) {
				open = append(open, neighbor)
			}
		}

		closed = append(closed, current)
	}
	return reachablePeaks
}

func part1(input input) int {
	heads := getTrailHeads(input)
	sum := 0
	for _, head := range heads {
		peaks := getReachablePeaks(input, head)
		var uniquePeaks [][2]int
		for _, peak := range peaks {
			if !slices.Contains(uniquePeaks, peak) {
				uniquePeaks = append(uniquePeaks, peak)
			}
		}
		sum += len(uniquePeaks)
	}
	return sum
}

func part2(input input) int {
	heads := getTrailHeads(input)
	sum := 0
	for _, head := range heads {
		peaks := getReachablePeaks(input, head)
		sum += len(peaks)
	}
	return sum
}
