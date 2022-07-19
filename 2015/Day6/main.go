package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bluestbluex/MyAoCCalendar/utils"
)

func main() {
	input, err := utils.GetInput("2015/Day6/input.txt")
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("solution part 1: %d\n", solutionPart1(input))
	fmt.Printf("solution part 2: %d\n", solutionPart2(input))
}

func solutionPart1(input string) int {
	var grid [1000][1000]bool
	utils.PerLine(input, func(_ int, s string) int {
		s = strings.Trim(s, "\r")
		instruction := strings.Split(s, " ")
		switch instruction[0] {
		case "toggle":
			start := strings.Split(instruction[1], ",")
			end := strings.Split(instruction[3], ",")
			startX, _ := strconv.Atoi(start[0])
			startY, _ := strconv.Atoi(start[1])
			endX, _ := strconv.Atoi(end[0])
			endY, _ := strconv.Atoi(end[1])
			for x := startX; x <= endX; x++ {
				for y := startY; y <= endY; y++ {
					grid[x][y] = !grid[x][y]
				}
			}
		case "turn":
			start := strings.Split(instruction[2], ",")
			end := strings.Split(instruction[4], ",")
			switch instruction[1] {
			case "on":
				startX, _ := strconv.Atoi(start[0])
				startY, _ := strconv.Atoi(start[1])
				endX, _ := strconv.Atoi(end[0])
				endY, _ := strconv.Atoi(end[1])
				for x := startX; x <= endX; x++ {
					for y := startY; y <= endY; y++ {
						grid[x][y] = true
					}
				}
			case "off":
				startX, _ := strconv.Atoi(start[0])
				startY, _ := strconv.Atoi(start[1])
				endX, _ := strconv.Atoi(end[0])
				endY, _ := strconv.Atoi(end[1])
				for x := startX; x <= endX; x++ {
					for y := startY; y <= endY; y++ {
						grid[x][y] = false
					}
				}
			}
		}
		return 0
	})

	return checkLightsPart1(grid)
}

func solutionPart2(input string) int {
	var grid [1000][1000]int
	utils.PerLine(input, func(_ int, s string) int {
		s = strings.Trim(s, "\r")
		instruction := strings.Split(s, " ")
		switch instruction[0] {
		case "toggle":
			start := strings.Split(instruction[1], ",")
			end := strings.Split(instruction[3], ",")
			startX, _ := strconv.Atoi(start[0])
			startY, _ := strconv.Atoi(start[1])
			endX, _ := strconv.Atoi(end[0])
			endY, _ := strconv.Atoi(end[1])
			for x := startX; x <= endX; x++ {
				for y := startY; y <= endY; y++ {
					grid[x][y] += 2
				}
			}
		case "turn":
			start := strings.Split(instruction[2], ",")
			end := strings.Split(instruction[4], ",")
			switch instruction[1] {
			case "on":
				startX, _ := strconv.Atoi(start[0])
				startY, _ := strconv.Atoi(start[1])
				endX, _ := strconv.Atoi(end[0])
				endY, _ := strconv.Atoi(end[1])
				for x := startX; x <= endX; x++ {
					for y := startY; y <= endY; y++ {
						grid[x][y]++
					}
				}
			case "off":
				startX, _ := strconv.Atoi(start[0])
				startY, _ := strconv.Atoi(start[1])
				endX, _ := strconv.Atoi(end[0])
				endY, _ := strconv.Atoi(end[1])
				for x := startX; x <= endX; x++ {
					for y := startY; y <= endY; y++ {
						grid[x][y]--
						if grid[x][y] < 0 {
							grid[x][y] = 0
						}
					}
				}
			}
		}
		return 0
	})

	return checkLightsPart2(grid)
}

func checkLightsPart1(grid [1000][1000]bool) int {
	var lights int
	for _, col := range grid {
		for _, row := range col {
			if row {
				lights++
			}
		}
	}

	return lights
}

func checkLightsPart2(grid [1000][1000]int) int {
	var total int
	for _, col := range grid {
		for _, row := range col {
			total += row
		}
	}

	return total
}
