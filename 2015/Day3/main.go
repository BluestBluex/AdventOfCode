package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/bluestbluex/MyAoCCalendar/utils"
)

func main() {
	input, err := utils.GetInput("2015/Day3/input.txt")
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("solution part 1: %d\n", solutionPart1(input))
	fmt.Printf("soultion part 2: %d\n", soultionPart2(input))
}

func solutionPart1(input string) int {
	var h, v int
	var visited = make(map[string]int)

	visited["0_0"] = 1

	for _, c := range string(input) {
		switch string(c) {
		case ">":
			v += 1
		case "<":
			v -= 1
		case "^":
			h += 1
		case "v":
			h -= 1
		}
		visited[strconv.Itoa(h)+"_"+strconv.Itoa(v)] += 1
	}

	return len(visited)
}

func soultionPart2(input string) int {
	var h, v, rh, rv int
	var visited = make(map[string]int)

	visited["0_0"] = 1

	for i, c := range string(input) {
		if i%2 == 0 {
			switch string(c) {
			case ">":
				rv += 1
			case "<":
				rv -= 1
			case "^":
				rh += 1
			case "v":
				rh -= 1
			}
			visited[strconv.Itoa(rh)+"_"+strconv.Itoa(rv)] += 1
			continue
		}
		switch string(c) {
		case ">":
			v += 1
		case "<":
			v -= 1
		case "^":
			h += 1
		case "v":
			h -= 1
		}
		visited[strconv.Itoa(h)+"_"+strconv.Itoa(v)] += 1
	}

	return len(visited)
}
