package utils

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func GetInput(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}

	return string(data), nil
}

func PerLine(input string, function func(int, string) int) int {
	var total int
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		total += function(i, line)
	}

	return total
}

func FindMinN(slice []int, n int) []int {
	if n == 0 {
		return []int{}
	}
	sort.Ints(slice)

	return slice[:n]
}
