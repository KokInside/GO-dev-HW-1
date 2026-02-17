package uniq

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"uniq/options"
	"uniq/truncate"
	"uniq/utils"
)

func Uniq(input []string, options options.Options) ([]string, error) {
	/////////////////////////////////////////////////////////////////////////////////////////////////

	var result []string

	if utils.TrueCount(options.C, options.D, options.U) > 1 {
		fmt.Println("Only one of flags -c, -u, -d can be used")
		fmt.Println("")
		flag.PrintDefaults()
		return nil, errors.New("Низя много флагов, ну ты чё")
	}

	// символ переноса строки
	// var endLine []byte = []byte("\n")

	// хранит предыдущую строку
	var prevString string

	// хранит количество встречаний строк
	var repeatCount int

	// была ли повторена строка
	var stringsRepeat bool = false

	// первая ли это строка ввода
	var newLine bool = true

	var prevStringToReturn string

	//for scanner.Scan() {
	for _, line := range input {

		if newLine {
			prevStringToReturn = line
		}

		var curString string = line

		// учитывать регистр ?
		if options.I {
			curString = strings.ToLower(curString)
		}

		if options.F != 0 {

			newCurString, err := truncate.TruncateStringFields(curString, options.F)

			if err != nil {
				fmt.Println(err.Error())
				return nil, errors.New("Много слов отрезал, ну ты чё")
			}

			curString = newCurString
		}

		if options.S != 0 {
			newCurString, err := truncate.TruncateStringSymbols(curString, options.S)

			if err != nil {
				fmt.Println(err.Error())
				return nil, errors.New("Много букав отрезал, ну ты чё")
			}

			curString = newCurString
		}

		////////////////////////////////////////

		// Если встрелити новую строку
		if prevString != curString || newLine {

			if options.C {

				if repeatCount != 0 {
					// writer.Write([]byte("    " + strconv.Itoa(repeatCount) + " " + prevStringToReturn))
					// writer.Write(endLine)
					result = append(result, "    "+strconv.Itoa(repeatCount)+" "+prevStringToReturn)

					repeatCount = 0

				}
				repeatCount++

			} else if options.U {

				if stringsRepeat == false && !newLine {

					// writer.Write([]byte(prevStringToReturn))
					// writer.Write(endLine)
					result = append(result, prevStringToReturn)

				}

				stringsRepeat = false

			} else if options.D {

				if stringsRepeat {

					// writer.Write([]byte(prevStringToReturn))
					// writer.Write(endLine)
					result = append(result, prevStringToReturn)

				}

				stringsRepeat = false

			} else {

				// writer.Write([]byte(prevStringToReturn))
				// writer.Write(endLine)
				// result = append(result, prevStringToReturn)
				result = append(result, line)
			}

			newLine = false

		} else {

			if options.C {
				repeatCount++

			} else if options.U {
				if stringsRepeat == false {
					stringsRepeat = true
				}

			} else if options.D {
				if stringsRepeat == false {
					stringsRepeat = true
				}
			}
		}

		prevString = curString

		prevStringToReturn = line
	}

	if options.C && repeatCount != 0 {
		result = append(result, "    "+strconv.Itoa(repeatCount)+" "+prevStringToReturn)
	}

	if options.D && stringsRepeat {
		result = append(result, prevStringToReturn)
	}

	if options.U && !stringsRepeat {
		result = append(result, prevStringToReturn)
	}

	/////////////////////////////////////////////////////////////////////////////////////////////////
	return result, nil
}
