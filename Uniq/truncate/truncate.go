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

	for i, s := range str {

		if unicode.IsSpace(s) {
			if prevSymbol {

				wordCount++
				prevSymbol = false
			}

		} else {
			prevSymbol = true
		}

		if wordCount == fieldCount {
			return (str)[i:], nil
		}
	}

	return "", nil
}

func TruncateSymbols(str string, symbolCount int) (string, error) {

	if symbolCount < 0 {
		return "", errors.New("[num_chars] must be positive")
	}

	if symbolCount > len(str) {
		return "", nil
	}

	return str[symbolCount:], nil
}
