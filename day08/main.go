package main

import (
	"bufio"
	"fmt"
	"os"
)

type antenna = struct {
	x      int
	y      int
	symbol rune
}

type input struct {
	width    int
	height   int
	antennas map[rune][]antenna
}

func main() {
	input := load()

	answerPart1 := part1(input)
	fmt.Printf("Answer part 1: %d\n", answerPart1)

	answerPart2 := part2(input)
	fmt.Printf("Answer part 2: %d\n", answerPart2)
}

func load() input {
	file, _ := os.Open("inputs/day08.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]rune
	for scanner.Scan() {
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}

	var antennas []antenna
	for y, line := range grid {
		for x, c := range line {
			if c == '.' {
				continue
			}
			antennas = append(antennas, antenna{x, y, c})
		}
	}

	bySymbol := make(map[rune][]antenna)
	for _, a := range antennas {
		bySymbol[a.symbol] = append(bySymbol[a.symbol], a)
	}

	return input{
		width:    len(grid[0]),
		height:   len(grid),
		antennas: bySymbol,
	}
}

func isValid(input input, position [2]int) bool {
	return position[0] >= 0 &&
		position[1] >= 0 &&
		position[0] < input.width &&
		position[1] < input.height
}

func part1(input input) int {
	sum := 0
	seen := make(map[[2]int]bool)
	for _, group := range input.antennas {
		for i := 0; i < len(group); i++ {
			for j := i + 1; j < len(group); j++ {
				a := group[i]
				b := group[j]
				dx := a.x - b.x
				dy := a.y - b.y

				position := [2]int{a.x + dx, a.y + dy}
				if isValid(input, position) && !seen[position] {
					sum++
					seen[position] = true
				}

				position = [2]int{b.x - dx, b.y - dy}
				if isValid(input, position) && !seen[position] {
					sum++
					seen[position] = true
				}
			}
		}
	}
	return sum
}

func part2(input input) int {
	sum := 0
	seen := make(map[[2]int]bool)
	for _, group := range input.antennas {
		for i := 0; i < len(group); i++ {
			for j := i + 1; j < len(group); j++ {
				a := group[i]
				b := group[j]
				dx := a.x - b.x
				dy := a.y - b.y

				position := [2]int{a.x, a.y}
				for isValid(input, position) {
					if !seen[position] {
						sum++
						seen[position] = true
					}
					position[0] += dx
					position[1] += dy
				}

				position = [2]int{b.x, b.y}
				for isValid(input, position) {
					if !seen[position] {
						sum++
						seen[position] = true
					}
					position[0] -= dx
					position[1] -= dy
				}
			}
		}
	}
	return sum
}
