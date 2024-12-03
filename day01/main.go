package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type input struct {
	left  []int
	right []int
}

func main() {
	input := load()

	answerPart1 := part1(input)
	fmt.Printf("Answer part 1: %d\n", answerPart1)

	answerPart2 := part2(input)
	fmt.Printf("Answer part 2: %d\n", answerPart2)
}

func load() input {
	file, _ := os.Open("inputs/day01.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r := regexp.MustCompile("\\d+")
	var left []int
	var right []int
	for scanner.Scan() {
		line := scanner.Text()
		match := r.FindAllString(line, 2)
		lhs, _ := strconv.Atoi(match[0])
		rhs, _ := strconv.Atoi(match[1])
		left = append(left, lhs)
		right = append(right, rhs)
	}
	return input{
		left:  left,
		right: right,
	}
}

func part1(input input) int {
	slices.Sort(input.left)
	slices.Sort(input.right)
	sum := 0
	for i := range input.left {
		dist := input.right[i] - input.left[i]
		if dist < 0 {
			dist = -dist
		}
		sum += dist
	}
	return sum
}

func part2(input input) int {
	sum := 0
	for _, n := range input.left {
		appearances := 0
		for _, m := range input.right {
			if m == n {
				appearances++
			}
		}
		sum += n * appearances
	}
	return sum
}
