package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type rule struct {
	low  int
	high int
}

type update = []int

type input struct {
	rules   []rule
	updates []update
}

func main() {
	input := load()

	answerPart1 := part1(input)
	fmt.Printf("Answer part 1: %d\n", answerPart1)

	answerPart2 := part2(input)
	fmt.Printf("Answer part 2: %d\n", answerPart2)
}

func load() input {
	file, _ := os.Open("inputs/day05.txt")
	defer file.Close()

	r := regexp.MustCompile("\\d+")

	scanner := bufio.NewScanner(file)
	var rules []rule
	for scanner.Scan() {
		line := scanner.Text()
		if string(line) == "" {
			break
		}
		match := r.FindAllString(line, 2)
		low, _ := strconv.Atoi(match[0])
		high, _ := strconv.Atoi(match[1])
		rules = append(rules, rule{
			low:  low,
			high: high,
		})
	}

	var updates []update
	for scanner.Scan() {
		line := scanner.Text()
		match := r.FindAllString(line, -1)
		var update update
		for _, asString := range match {
			n, _ := strconv.Atoi(asString)
			update = append(update, n)
		}
		updates = append(updates, update)
	}
	return input{
		rules:   rules,
		updates: updates,
	}
}

func sort(rules []rule, update update) update {
	updateCopy := make([]int, len(update))
	copy(updateCopy, update)
	slices.SortFunc(updateCopy, func(i, j int) int {
		if slices.ContainsFunc(rules, func(rule rule) bool {
			return rule.low == i && rule.high == j
		}) {
			return -1
		} else {
			return 1
		}
	})
	return updateCopy
}

func part1(input input) int {
	sum := 0
	for _, update := range input.updates {
		sorted := sort(input.rules, update)
		if slices.Equal(update, sorted) {
			middle := len(sorted) / 2
			sum += sorted[middle]
		}
	}
	return sum
}

func part2(input input) int {
	sum := 0
	for _, update := range input.updates {
		sorted := sort(input.rules, update)
		if !slices.Equal(update, sorted) {
			middle := len(sorted) / 2
			sum += sorted[middle]
		}
	}
	return sum
}
