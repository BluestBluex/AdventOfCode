package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bluestbluex/MyAoCCalendar/utils"
)

func main() {
	input, err := utils.GetInput("2015/Day2/input.txt")
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Total wrapping paper: %d\n", solutionPart1(input))
	fmt.Printf("Total ribbon length: %d\n", solutionPart2(input))
}

func solutionPart1(input string) int {
	return utils.PerLine(input, func(i int, s string) int {
		dim := getDim(s)
		sideA := dim[0] * dim[1]
		sideB := dim[1] * dim[2]
		sideC := dim[2] * dim[0]

		return (2 * sideA) + (2 * sideB) + (2 * sideC) + utils.FindMinN([]int{sideA, sideB, sideC}, 1)[0]
	})
}

func solutionPart2(input string) int {
	return utils.PerLine(input, func(i int, s string) int {
		dim := getDim(s)
		min := utils.FindMinN(dim, 2)
		return (2 * min[0]) + (2 * min[1]) + (dim[0] * dim[1] * dim[2])
	})
}

func getDim(input string) []int {
	dimensions := strings.Split(input, "x")

	length, err := strconv.Atoi(dimensions[0])
	if err != nil {
		log.Panicf("failed to read input: %v", err)
	}

	width, err := strconv.Atoi(dimensions[1])
	if err != nil {
		log.Panicf("failed to read input: %v", err)
	}

	dimensions[2] = strings.Trim(dimensions[2], "\r")
	height, err := strconv.Atoi(dimensions[2])
	if err != nil {
		log.Panicf("failed to read input: %v", err)
	}

	return []int{length, width, height}
}
