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

func load() []int {
	file, _ := os.ReadFile("inputs/day11.txt")
	r := regexp.MustCompile("\\d+")
	match := r.FindAllString(string(file), -1)
	var input []int
	for _, s := range match {
		n, _ := strconv.Atoi(s)
		input = append(input, n)
	}
	return input
}

var cache = make(map[[2]int]int)

func getStoneCount(n int, blinkCount int) int {
	if blinkCount == 0 {
		return 1
	}

	key := [2]int{n, blinkCount}
	if value, ok := cache[key]; ok {
		return value
	}

	var result int
	if n == 0 {
		result = getStoneCount(1, blinkCount-1)
	} else if asString := strconv.Itoa(n); len(asString)%2 == 0 {
		a, _ := strconv.Atoi(asString[:len(asString)/2])
		b, _ := strconv.Atoi(asString[len(asString)/2:])
		result = getStoneCount(a, blinkCount-1) + getStoneCount(b, blinkCount-1)
	} else {
		result = getStoneCount(n*2024, blinkCount-1)
	}

	cache[key] = result
	return result
}

func part1(input []int) int {
	sum := 0
	for _, n := range input {
		sum += getStoneCount(n, 25)
	}
	return sum
}

func part2(input []int) int {
	sum := 0
	for _, n := range input {
		sum += getStoneCount(n, 75)
	}
	return sum
}
