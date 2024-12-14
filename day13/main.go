package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	a     [2]int
	b     [2]int
	prize [2]int
}

type input []game

func main() {
	input := load()

	answerPart1 := part1(input)
	fmt.Printf("Answer part 1: %d\n", answerPart1)

	answerPart2 := part2(input)
	fmt.Printf("Answer part 2: %d\n", answerPart2)
}

func load() input {
	file, _ := os.ReadFile("inputs/day13.txt")
	var input []game

	r := regexp.MustCompile("\\d+")
	for _, c := range strings.Split(string(file), "\n\n") {
		lines := strings.Split(c, "\n")

		matchA := r.FindAllString(lines[0], 2)
		matchB := r.FindAllString(lines[1], 2)
		prize := r.FindAllString(lines[2], 2)

		aX, _ := strconv.Atoi(matchA[0])
		aY, _ := strconv.Atoi(matchA[1])
		bX, _ := strconv.Atoi(matchB[0])
		bY, _ := strconv.Atoi(matchB[1])
		prizeX, _ := strconv.Atoi(prize[0])
		prizeY, _ := strconv.Atoi(prize[1])

		g := game{
			a:     [2]int{aX, aY},
			b:     [2]int{bX, bY},
			prize: [2]int{prizeX, prizeY},
		}
		input = append(input, g)
	}
	return input
}

func solveGame(game game) (int, int) {
	for a := 0; a < 200; a++ {
		for b := 0; b < 200; b++ {
			if game.a[0]*a+game.b[0]*b == game.prize[0] &&
				game.a[1]*a+game.b[1]*b == game.prize[1] {
				return a, b
			}
		}
	}
	return -1, -1
}

func part1(input input) int {
	sum := 0
	for _, g := range input {
		a, b := solveGame(g)
		if a >= 0 && b >= 0 {
			sum += a*3 + b
		}
	}
	return sum
}

func part2(input input) int {
	sum := 0
	for _, g := range input {
		ax := float64(g.a[0])
		ay := float64(g.a[1])
		bx := float64(g.b[0])
		by := float64(g.b[1])
		ad := ay / ax
		bd := by / bx
		tx := float64(g.prize[0]) + 10000000000000
		ty := float64(g.prize[1]) + 10000000000000

		x := (-bd*tx + ty) / (ad - bd)

		a := int(math.Round(x / ax))
		b := int(math.Round((tx - x) / bx))

		if a < 0 || b < 0 ||
			a*g.a[0]+b*g.b[0] != int(tx) ||
			a*g.a[1]+b*g.b[1] != int(ty) {
			continue
		}

		sum += a*3 + b
	}
	return sum
}
