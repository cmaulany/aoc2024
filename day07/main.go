package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type row struct {
	result int
	values []int
}

type input = []row

func main() {
	input := load()

	answerPart1 := part1(input)
	fmt.Printf("Answer part 1: %d\n", answerPart1)

	// answerPart2 := part2(input)
	// fmt.Printf("Answer part 2: %d\n", answerPart2)
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

		var row row
		n, _ := strconv.Atoi(match[0])
		row.result = n
		for _, val := range match[1:] {
			n, _ := strconv.Atoi(val)
			row.values = append(row.values, n)
		}
		input = append(input, row)
	}
	return input
}

var cache = make(map[string]bool)

func concat(a, b int) int {
	n, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	return n
}

func isSolvable(r row) bool {
	key := []int{r.result}
	for _, n := range r.values {
		key = append(key, n)
	}
	k := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(key)), ","), "[]")
	if v, ok := cache[k]; ok {
		return v
	}
	if len(r.values) == 2 {
		return r.values[0]*r.values[1] == r.result ||
			r.values[0]+r.values[1] == r.result ||
			concat(r.values[0], r.values[1]) == r.result
	}
	return isSolvable(row{
		result: r.result,
		values: append([]int{r.values[0] + r.values[1]}, r.values[2:]...),
	}) || isSolvable(row{
		result: r.result,
		values: append([]int{r.values[0] * r.values[1]}, r.values[2:]...),
	}) || isSolvable(row{
		result: r.result,
		values: append([]int{concat(r.values[0], r.values[1])}, r.values[2:]...),
	})
}

func part1(input input) int {
	sum := 0
	for _, row := range input {
		if isSolvable(row) {
			sum += row.result
		}
	}
	return sum
}
