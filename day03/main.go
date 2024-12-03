package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input := load()

	answerPart1 := part1(input)
	fmt.Printf("Answer part 1: %d\n", answerPart1)

	answerPart2 := part2(input)
	fmt.Printf("Answer part 2: %d\n", answerPart2)
}

func load() string {
	file, _ := os.ReadFile("inputs/day03.txt")
	return string(file)
}

func part1(input string) int {
	r := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
	match := r.FindAllStringSubmatch(input, -1)

	sum := 0
	for _, m := range match {
		a, _ := strconv.Atoi(m[1])
		b, _ := strconv.Atoi(m[2])
		sum += a * b
	}
	return sum
}

func part2(input string) int {
	r := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)|do\\(\\)|don't\\(\\)")
	match := r.FindAllStringSubmatch(input, -1)

	sum := 0
	do := true
	for _, m := range match {
		if m[0] == "do()" {
			do = true
		} else if m[0] == "don't()" {
			do = false
		} else if do {
			a, _ := strconv.Atoi(m[1])
			b, _ := strconv.Atoi(m[2])
			sum += a * b
		}
	}
	return sum
}
