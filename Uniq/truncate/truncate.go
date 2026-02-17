package truncate

import (
	"errors"
	"unicode"
)

func TruncateStringFields(str string, fieldCount int) (string, error) {

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
			return str[i:], nil
		}
	}

	//return "", errors.New("[num_fields] is more than string field count. Reduce [num_fields]")
	return "", nil
}

func TruncateStringSymbols(str string, symbolCount int) (string, error) {

	if symbolCount < 0 {
		return "", errors.New("[num_chars] must be positive")
	}

	if symbolCount > len(str) {
		return "", nil
		//return "", errors.New("[num_chars] is more than string symbols count. Reduce [num_chars]")
	}

	return str[symbolCount:], nil
}
