package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bluestbluex/MyAoCCalendar/utils"
)

func main() {
	input, err := utils.GetInput("2015/Day8/input.txt")
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("solution part 1: %d\n", solutionPart1(input))
	fmt.Printf("solution part 2: %d\n", solutionPart2(input))
}

func solutionPart1(input string) int {
	total := utils.PerLine(input, func(_ int, s string) int {
		s = strings.Trim(s, "\r")
		unquote, err := strconv.Unquote(s)
		if err != nil {
			log.Panic(err)
		}
		return len(s) - len(unquote)
	})

	return total
}

func solutionPart2(input string) int {
	total := utils.PerLine(input, func(_ int, s string) int {
		s = strings.Trim(s, "\r")
		quote := strconv.Quote(s)
		return len(quote) - len(s)
	})

	return total
}
