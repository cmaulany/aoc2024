package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, _ := os.Open("inputs/day1.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r, _ := regexp.Compile("\\d+")
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
	// sort.Ints(left)
	// sort.Ints(right)

	// fmt.Println(left)
	// fmt.Println(right)
	sum := 0
	for _, n := range left {
		count := 0
		for _, m := range right {
			if m == n {
				count++
			}
		}
		sum += count * n
	}
	// for index, _ := range left {
	// 	dist := right[index] - left[index]
	// 	if dist < 0 {
	// 		dist = -dist
	// 	}
	// 	sum += dist

	// }
	fmt.Println((sum))
}
