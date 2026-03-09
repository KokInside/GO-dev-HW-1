package utils

import (
	"bufio"
	"io"
	"strings"
)

func ReadLines(reader io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(reader)

	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func ConcatStrings(str1, str2 *string) string {
	var builder strings.Builder

	builder.Grow(len(*str1) + len(*str2))
	builder.WriteString(*str1)
	builder.WriteString(*str2)

	return builder.String()
}

func BoolCount(bools ...bool) int {
	var counter int

	for _, b := range bools {
		if b {
			counter++
		}
	}

	return counter
}
