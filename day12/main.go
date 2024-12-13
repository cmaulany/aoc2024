package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type input [][]rune

func main() {
	input := load()

	answerPart1 := part1(input)
	fmt.Printf("Answer part 1: %d\n", answerPart1)

	answerPart2 := part2(input)
	fmt.Printf("Answer part 2: %d\n", answerPart2)
}

func load() input {
	file, _ := os.Open("inputs/day12.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input [][]rune
	for scanner.Scan() {
		line := []rune(scanner.Text())
		input = append(input, line)
	}
	return input
}

func calculatePerimeter(group [][2]int) int {
	sum := 0
	for _, pos := range group {
		x := pos[0]
		y := pos[1]
		if !slices.Contains(group, [2]int{x - 1, y}) {
			sum++
		}
		if !slices.Contains(group, [2]int{x, y - 1}) {
			sum++
		}
		if !slices.Contains(group, [2]int{x + 1, y}) {
			sum++
		}
		if !slices.Contains(group, [2]int{x, y + 1}) {
			sum++
		}
	}
	return sum
}

func calculateSides(group [][2]int) int {
	var xs []int
	var ys []int
	for _, pos := range group {
		xs = append(xs, pos[0])
		ys = append(ys, pos[1])
	}
	minX := slices.Min(xs)
	maxX := slices.Max(xs) + 1
	minY := slices.Min(ys)
	maxY := slices.Max(ys) + 1

	sum := 0
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			tl := slices.Contains(group, [2]int{x - 1, y - 1})
			tr := slices.Contains(group, [2]int{x, y - 1})
			bl := slices.Contains(group, [2]int{x - 1, y})
			br := slices.Contains(group, [2]int{x, y})

			if !tl && !tr && !bl && br {
				sum += 2
			} else if !tl && !tr && bl && !br {
				sum += 1
			} else if !tl && tr && !bl && !br {
				sum += 1
			} else if tl && !tr && !bl && !br {
				sum += 0
			} else if tl && tr && bl && !br {
				sum += 2
			} else if tl && tr && !bl && br {
				sum += 1
			} else if tl && !tr && bl && br {
				sum += 1
			} else if !tl && tr && bl && br {
				sum += 0
			} else if tl && !tr && !bl && br {
				sum += 2
			} else if !tl && tr && bl && !br {
				sum += 2
			}
		}
	}
	return sum
}

func getGroups(input input) map[int][][2]int {
	nextGroupId := 0
	posToGroup := make(map[[2]int]int)
	groups := make(map[int][][2]int)

	for y, line := range input {
		for x, c := range line {
			pos := [2]int{x, y}
			if y > 0 && x > 0 && posToGroup[[2]int{x - 1, y}] != posToGroup[[2]int{x, y - 1}] && input[y-1][x] == c && input[y][x-1] == c {
				topGroupId := posToGroup[[2]int{x, y - 1}]
				leftGroupId := posToGroup[[2]int{x - 1, y}]

				for _, pos := range groups[leftGroupId] {
					posToGroup[pos] = topGroupId
				}
				groups[topGroupId] = append(groups[topGroupId], groups[leftGroupId]...)
				delete(groups, leftGroupId)

				posToGroup[pos] = topGroupId
				groups[topGroupId] = append(groups[topGroupId], pos)
			} else if y > 0 && input[y-1][x] == c {
				groupId := posToGroup[[2]int{x, y - 1}]
				posToGroup[pos] = groupId
				groups[groupId] = append(groups[groupId], pos)
			} else if x > 0 && input[y][x-1] == c {
				groupId := posToGroup[[2]int{x - 1, y}]
				posToGroup[pos] = groupId
				groups[groupId] = append(groups[groupId], pos)
			} else {
				posToGroup[pos] = nextGroupId
				groups[nextGroupId] = append(groups[nextGroupId], pos)
				nextGroupId++
			}
		}
	}
	return groups
}

func part1(input input) int {
	groups := getGroups(input)

	sum := 0
	for _, group := range groups {
		sum += calculatePerimeter(group) * len(group)
	}

	return sum
}

func part2(input input) int {
	groups := getGroups(input)

	sum := 0
	for _, group := range groups {
		sum += calculateSides(group) * len(group)
	}

	return sum
}
