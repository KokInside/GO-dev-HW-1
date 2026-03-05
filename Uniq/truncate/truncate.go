package truncate

import (
	"errors"
	"unicode"
)

func TruncateFields(str string, fieldCount int) (string, error) {
	if fieldCount < 0 {
		return "", errors.New("[num_fields] must be positive")
	}

	var prevSymbol bool
	var wordCount int

	runes := []rune(str)

	for i, r := range runes {
		if unicode.IsSpace(r) {
			if prevSymbol {
				wordCount++
				prevSymbol = false
			}

		} else {
			prevSymbol = true
		}

		if wordCount == fieldCount {
			return string(runes[i:]), nil
		}
	}

	return "", nil
}

func TruncateSymbols(str string, symbolCount int) (string, error) {
	if symbolCount < 0 {
		return "", errors.New("[num_chars] must be positive")
	}

	runes := []rune(str)

	if symbolCount > len(runes) {
		return "", nil
	}

	return string(runes[symbolCount:]), nil
}
