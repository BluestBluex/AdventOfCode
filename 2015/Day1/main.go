package main

import (
	"fmt"
	"log"

	"github.com/bluestbluex/MyAoCCalendar/utils"
)

func main() {
	input, err := utils.GetInput("2015/Day1/input.txt")
	if err != nil {
		log.Panic(err)
	}

	solution(input)
}

func solution(input string) {
	var currentFloor int
	var foundBasement bool

	for i, c := range string(input) {
		switch string(c) {
		case "(":
			currentFloor++
		case ")":
			currentFloor--
		}
		// Added for part 2
		if currentFloor == -1 && !foundBasement {
			fmt.Printf("First basement: %d\n", i+1)
			foundBasement = true
		}
	}

	fmt.Printf("Final floor: %d\n", currentFloor)
}
