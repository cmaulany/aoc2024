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
	for scanner.Scan() {
		line := scanner.Text()
		match := r.FindAllString(line, 2)
		lhs, _ := strconv.Atoi(match[0])
		rhs, _ := strconv.Atoi(match[1])
		fmt.Printf("%d - %d = %d\n", lhs, rhs, lhs-rhs)
	}
}
