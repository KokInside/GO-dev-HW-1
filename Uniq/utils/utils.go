package utils

import (
	"bufio"
	"io"
)

func TrueCount(a ...bool) int {
	var sum int

	for _, i := range a {
		if i {
			sum++
		}
	}

	return sum
}

func ReadLines(reader io.Reader) ([]string, error) {

	scanner := bufio.NewScanner(reader)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
