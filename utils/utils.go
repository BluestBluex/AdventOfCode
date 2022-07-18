package utils

import (
	"fmt"
	"os"
)

func GetInput(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}

	return string(data), nil
}
