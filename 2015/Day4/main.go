package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	var input string = "bgvyzdsv"

	fmt.Printf("solution part 1: %d\n", solutionPart1(input))
	fmt.Printf("solution part 2: %d\n", solutionPart2(input))
}

func solutionPart1(input string) int {
	for i := 0; ; i++ {
		hash := md5.Sum([]byte(input + strconv.Itoa(i)))
		hashString := hex.EncodeToString(hash[:])

		if hashString[:5] == "00000" {
			return i
		}
	}
}

func solutionPart2(input string) int {
	for i := 0; ; i++ {
		hash := md5.Sum([]byte(input + strconv.Itoa(i)))
		hashString := hex.EncodeToString(hash[:])

		if hashString[:6] == "000000" {
			return i
		}
	}
}
