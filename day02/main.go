package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	file, _ := os.Open("inputs/day02.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r := regexp.MustCompile("\\d+")
	var input [][]int
	for scanner.Scan() {
		line := scanner.Text()
		match := r.FindAllString(line, -1)
		var l []int
		for _, s := range match {
			n, _ := strconv.Atoi(s)
			l = append(l, n)
		}
		input = append(input, l)
	}
	return input
}

func part1(input input) int {
	sum := 0
	for _, line := range input {
		if isSafe((line)) {
			sum++
		}
	}
	return sum
}

func part2(input input) int {
	sum := 0
	for _, line := range input {
	loop:
		for i := 0; i < len(line); i++ {
			shrunkLine := make([]int, 0, len(line)-1)
			shrunkLine = append(shrunkLine, line[:i]...)
			shrunkLine = append(shrunkLine, line[i+1:]...)
			if isSafe(shrunkLine) {
				sum++
				break loop
			}
		}
	}
	return sum
}

func isSafe(level []int) bool {
	dir := 1
	if level[1] < level[0] {
		dir = -1
	}
	for i := 0; i < len(level)-1; i++ {
		a := level[i]
		b := level[i+1]
		if (b-a)*dir < 1 {
			return false
		}
		delta := b - a
		if delta < 0 {
			delta *= -1
		}
		if delta < 1 || delta > 3 {
			return false
		}
	}
	return true
}
