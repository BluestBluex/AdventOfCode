package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bluestbluex/MyAoCCalendar/utils"
)

func main() {
	input, err := utils.GetInput("2015/Day5/input.txt")
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("solution part 1: %d\n", solutionPart1(input))
	fmt.Printf("solution part 2: %d\n", solutionPart2(input))
}

func solutionPart1(input string) int {
	return utils.PerLine(input, func(_ int, s string) int {
		if strings.Contains(s, "ab") || strings.Contains(s, "cd") || strings.Contains(s, "pq") || strings.Contains(s, "xy") {
			return 0
		}

		var vowelCount int
		var prevLetter string
		var repeatCheck bool

		for _, c := range s {
			if prevLetter == string(c) {
				repeatCheck = true
			}
			prevLetter = string(c)

			if strings.ContainsAny(string(c), "aeiou") {
				vowelCount++
			}
		}

		if vowelCount < 3 || !repeatCheck {
			return 0
		}

		return 1
	})
}

func solutionPart2(input string) int {
	return utils.PerLine(input, func(_ int, s string) int {
		var prevLetter string
		var twoLetterBack string
		var repeatCheck bool
		var betweenCheck bool

		for i, c := range s {
			if i == 0 {
				prevLetter = string(c)
				continue
			}
			if twoLetterBack == string(c) {
				betweenCheck = true
			}
			twoLetterBack = prevLetter

			pair := prevLetter + string(c)
			if strings.Contains(s[:i-1], pair) {
				repeatCheck = true
			}
			prevLetter = string(c)
		}

		if repeatCheck && betweenCheck {
			return 1
		}

		return 0
	})
}
