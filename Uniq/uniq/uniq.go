package uniq

import (
	"errors"
	"uniq/flagchecker"
	"uniq/options"
	"uniq/utils"
)

func Uniq(input []string, options options.Options) ([]string, error) {

	result := make([]string, 0)

	if utils.BoolCount(options.Count, options.Repeated, options.Unique) > 1 {
		return nil, errors.New("Only one flag of -c, -d and -u can be used.")
	}

	// первая ли это строка ввода
	isFirstString := true

	// хранит количество встречаний строк
	var repeatCount int

	// была ли повторена строка
	var isStringRepeat bool

	var prevString string

	var prevStringToReturn string

	for _, line := range input {

		curString := line

		// учитывать регистр ?
		curString = flagchecker.CheckIgnoreCase(curString, options)

		// флаг -f
		curString, err := flagchecker.CheckSkipFields(curString, options)

		if err != nil {
			return nil, err
		}

		// флаг -s
		curString, err = flagchecker.CheckSkipChars(curString, options)

		if err != nil {
			return nil, err
		}

		// Если встретили новую строку
		if prevString != curString || isFirstString {

			newLine, updated := flagchecker.CheckUniqueFlags(line, prevStringToReturn, prevString, options, isStringRepeat, isFirstString, repeatCount)

			if updated {
				result = append(result, newLine)
			}

			repeatCount = 1
			isFirstString = false
			isStringRepeat = false

			// если встретили старую строку
		} else {

			repeatCount++
			isStringRepeat = true
		}

		prevString = curString
		prevStringToReturn = line
	}

	// нужно обработать последнюю введённую строку в соответствии с флагами
	if utils.BoolCount(options.Count, options.Repeated, options.Unique) != 0 {
		newLine, updated := flagchecker.CheckUniqueFlags("", prevStringToReturn, prevString, options, isStringRepeat, isFirstString, repeatCount)

		if updated {
			result = append(result, newLine)
		}
	}

	return result, nil
}
