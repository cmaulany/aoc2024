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

	// answerPart2 := part2(input)
	// fmt.Printf("Answer part 2: %d\n", answerPart2)
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

func order(rules []rule, update update) []int {
	openNs := make([]int, len(update))
	copy(openNs, update)
	var openRules []rule
	for _, rule := range rules {
		if slices.Contains(openNs, rule.low) && slices.Contains(openNs, rule.high) {
			openRules = append(openRules, rule)
		}
	}

	var ordered []int

outer:
	for len(openNs) > 0 {
	inner:
		for i := 0; i < len(openNs); i++ {
			n := openNs[i]
			// fmt.Println(n)
			for _, rule := range openRules {
				if rule.high == n {
					continue inner
				}
			}
			for j := 0; j < len(openRules); j++ {
				rule := openRules[j]
				if rule.low == n {
					openRules[j] = openRules[len(openRules)-1]
					openRules = openRules[:len(openRules)-1]
					j--
				}
			}
			ordered = append(ordered, n)

			openNs[i] = openNs[len(openNs)-1]
			openNs = openNs[:len(openNs)-1]
			i--
			continue outer
		}
	}
	return ordered
}

func part1(input input) int {
	sum := 0
loop:
	for _, update := range input.updates {
		var seen []int
		valid := true
		for _, n := range update {
			seen = append(seen, n)
			for _, rule := range input.rules {
				if rule.low == n && slices.Contains(seen, rule.high) {
					fmt.Println("XX")
					valid = false
					ordered := order(input.rules, update)
					middle := len(update) / 2
					fmt.Printf("%d %d\n", sum, ordered[middle])
					sum += ordered[middle]
					continue loop
				}
			}
		}
		if valid {
			middle := len(update) / 2
			fmt.Println(middle)
			// sum += update[middle]
		}
	}
	return sum
}

// func part2(input input) int {
// 	sum := 0
// 	for y, line := range input {
// 		for x := range line {
// 			if checkXMas(input, [2]int{x - 1, y - 1}) {
// 				sum++
// 			}
// 		}
// 	}
// 	return sum
// }
