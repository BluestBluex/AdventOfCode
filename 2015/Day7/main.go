package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bluestbluex/MyAoCCalendar/utils"
)

func main() {
	input, err := utils.GetInput("2015/Day7/input.txt")
	if err != nil {
		log.Panic(err)
	}

	solution(input)
}

func solution(input string) {
	var operation = make(map[string]string)
	utils.PerLine(input, func(_ int, s string) int {
		s = strings.Trim(s, "\r")
		ops := strings.Split(s, "->")
		operation[strings.Trim(ops[1], " ")] = strings.Trim(ops[0], " ")
		return 0
	})

	fmt.Printf("solution: %v\n", getWireValue("a", operation))
}

func getWireValue(key string, operation map[string]string) uint16 {
	// turn on for part2, with answer from part1
	if key == "b" {
		return uint16(956)
	}

	circuit := strings.Split(operation[key], " ")
	switch len(circuit) {
	case 1:
		value, err := strconv.ParseUint(circuit[0], 10, 16)
		if err != nil {
			result := getWireValue(circuit[0], operation)
			operation[key] = strconv.Itoa(int(result))
			return result
		}
		return uint16(value)
	case 2:
		value, err := strconv.ParseUint(circuit[1], 10, 16)
		if err != nil {
			result := ^getWireValue(circuit[1], operation)
			operation[key] = strconv.Itoa(int(result))
			return result
		}
		return ^uint16(value)
	case 3:
		valueLeft, err := strconv.ParseUint(circuit[0], 10, 16)
		if err != nil {
			valueLeft = uint64(getWireValue(circuit[0], operation))
		}
		valueRight, err := strconv.ParseUint(circuit[2], 10, 16)
		if err != nil {
			valueRight = uint64(getWireValue(circuit[2], operation))
		}
		switch circuit[1] {
		case "AND":
			result := uint16(valueLeft) & uint16(valueRight)
			operation[key] = strconv.Itoa(int(result))
			return result
		case "OR":
			result := uint16(valueLeft) | uint16(valueRight)
			operation[key] = strconv.Itoa(int(result))
			return result
		case "RSHIFT":
			result := uint16(valueLeft) >> uint16(valueRight)
			operation[key] = strconv.Itoa(int(result))
			return result
		case "LSHIFT":
			result := uint16(valueLeft) << uint16(valueRight)
			operation[key] = strconv.Itoa(int(result))
			return result
		}
	}

	return 0
}
