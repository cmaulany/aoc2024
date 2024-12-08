package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type equation struct {
	result int
	ns     []int
}

type input = []equation

type operation func(int, int) int

func main() {
	input := load()

	answerPart1 := part1(input)
	fmt.Printf("Answer part 1: %d\n", answerPart1)

	answerPart2 := part2(input)
	fmt.Printf("Answer part 2: %d\n", answerPart2)
}

func load() input {
	file, _ := os.Open("inputs/day07.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r := regexp.MustCompile("\\d+")
	var input input
	for scanner.Scan() {
		line := scanner.Text()
		match := r.FindAllString(line, -1)

		var eq equation
		n, _ := strconv.Atoi(match[0])
		eq.result = n
		for _, val := range match[1:] {
			n, _ := strconv.Atoi(val)
			eq.ns = append(eq.ns, n)
		}
		input = append(input, eq)
	}
	return input
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func concat(a, b int) int {
	n, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	return n
}

func isSolvable(eq equation, operations []operation) bool {
	if len(eq.ns) == 1 {
		return eq.result == eq.ns[0]
	}
	return slices.ContainsFunc(operations, func(operation operation) bool {
		return isSolvable(
			equation{
				result: eq.result,
				ns:     append([]int{operation(eq.ns[0], eq.ns[1])}, eq.ns[2:]...),
			},
			operations,
		)
	})
}

func part1(input input) int {
	operations := []operation{
		add,
		multiply,
	}

	sum := 0
	for _, eq := range input {
		if isSolvable(eq, operations) {
			sum += eq.result
		}
	}
	return sum
}

func part2(input input) int {
	operations := []operation{
		add,
		multiply,
		concat,
	}

	sum := 0
	for _, eq := range input {
		if isSolvable(eq, operations) {
			sum += eq.result
		}
	}
	return sum
}
