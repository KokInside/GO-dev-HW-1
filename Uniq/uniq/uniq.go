package uniq

import (
	"errors"
	"strconv"
	"strings"
	"uniq/options"
	"uniq/strutils"
	"uniq/truncate"
)

func Uniq(input []string, options options.Options) ([]string, error) {

	var result []string

	if options.D && options.U {
		return nil, errors.New("Flags -d and -u can't be used simultaneously.")
	}

	// первая ли это строка ввода
	var isFirstString bool = true

	// хранит количество встречаний строк
	var repeatCount int

	// была ли повторена строка
	var isStringRepeat bool

	var prevString string

	var prevStringToReturn string

	for _, line := range input {

		var curString string = line

		// учитывать регистр ?
		if options.I {
			curString = strings.ToLower(curString)
		}

		// флаг -f
		if options.F != 0 {
			f_FlagErr := truncate.TruncateFields(&curString, options.F)

			if f_FlagErr != nil {
				return nil, f_FlagErr
			}
		}

		// флаг -s
		if options.S != 0 {
			s_FlagErr := truncate.TruncateSymbols(&curString, options.S)

			if s_FlagErr != nil {
				return nil, s_FlagErr
			}
		}

		// Если встретили новую строку
		if prevString != curString || isFirstString {

			if options.C {
				var prefix string = "    " + strconv.Itoa(repeatCount) + " "
				prevStringToReturn = strutils.ConcatStrings(&prefix, &prevString)
			}

			if options.D {

				if isStringRepeat && !isFirstString {

					result = append(result, prevStringToReturn)
				}

			} else if options.U {

				if isStringRepeat == false && !isFirstString {

					result = append(result, prevStringToReturn)
				}

			} else if options.C {

				if !isFirstString {

					result = append(result, prevStringToReturn)
				}

			} else {

				result = append(result, line)
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

	if options.C {
		var prefix string = "    " + strconv.Itoa(repeatCount) + " "
		prevStringToReturn = strutils.ConcatStrings(&prefix, &prevString)
	}

	if options.D {

		if isStringRepeat && !isFirstString {
			result = append(result, prevStringToReturn)
		}

	} else if options.U && !isFirstString {

		if !isStringRepeat {
			result = append(result, prevStringToReturn)
		}

	} else if options.C {
		if !isFirstString {

			result = append(result, prevStringToReturn)
		}
	}

	return result, nil
}
