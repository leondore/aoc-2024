package utils

import (
	"fmt"
	"os"
	"strings"
)

func ProcessInput(path string) ([]string, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	stringify := strings.TrimSpace(string(dat))
	return strings.Split(stringify, "\n"), nil
}
