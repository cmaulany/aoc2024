package main

import (
	"fmt"
	"os"
	"strconv"
)

type input = []int

type block = [2]int

func main() {
	input := load()

	answerPart1 := part1(input)
	fmt.Printf("Answer part 1: %d\n", answerPart1)

	answerPart2 := part2(input)
	fmt.Printf("Answer part 2: %d\n", answerPart2)
}

func load() input {
	file, _ := os.ReadFile("inputs/day09.txt")
	var input []int
	for _, c := range string(file) {
		n, _ := strconv.Atoi(string(c))
		input = append(input, n)
	}
	return input
}

func part1(input input) int {
	var disk []int
	for i := 0; i < len(input); i++ {
		size := input[i]
		for j := 0; j < size; j++ {
			var symbol int
			if i%2 == 0 {
				symbol = i / 2
			} else {
				symbol = -1
			}
			disk = append(disk, symbol)
		}
	}

	cursor := 0
	for i := len(disk) - 1; i > cursor; i-- {
		if disk[i] == -1 {
			continue
		}
		for disk[cursor] != -1 {
			cursor++
		}
		disk[cursor] = disk[i]
		disk[i] = -1
	}

	checksum := 0
	for i, n := range disk {
		if n != -1 {
			checksum += i * n
		}
	}
	return checksum
}

func part2(input input) int {
	var disk []block
	for i := 0; i < len(input); i++ {
		size := input[i]
		var symbol int
		if i%2 == 0 {
			symbol = i / 2
		} else {
			symbol = -1
		}
		disk = append(disk, [2]int{symbol, size})
	}
	for i := len(disk) - 1; i >= 0; i-- {
		symbol := disk[i][0]
		if symbol == -1 {
			continue
		}
		fileSize := disk[i][1]
		for j := 0; j < i; j++ {
			availableSize := disk[j][1]
			if disk[j][0] == -1 && availableSize >= fileSize {
				nextDisk := make([]block, 0, len(disk))
				nextDisk = append(nextDisk, disk[:j]...)
				nextDisk = append(nextDisk, disk[i])
				if availableSize > fileSize {
					nextDisk = append(nextDisk, [2]int{-1, availableSize - fileSize})
				}
				nextDisk = append(nextDisk, disk[j+1:i]...)
				nextDisk = append(nextDisk, [2]int{-1, fileSize})
				nextDisk = append(nextDisk, disk[i+1:]...)
				disk = nextDisk
				break
			}
		}
	}

	checksum := 0
	index := 0
	for _, block := range disk {
		symbol := block[0]
		size := block[1]
		if symbol != -1 {
			for i := 0; i < size; i++ {
				checksum += symbol * (i + index)
			}
		}
		index += size
	}
	return checksum
}
