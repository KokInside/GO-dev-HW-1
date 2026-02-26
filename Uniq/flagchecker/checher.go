package flagchecker

import (
	"strconv"
	"strings"
	"uniq/options"
	"uniq/truncate"
	"uniq/utils"
)

func CheckIgnoreCase(curString string, options options.Options) string {

	if options.IgnoreCase {
		return strings.ToLower(curString)
	}

	return curString
}

func CheckSkipFields(curString string, options options.Options) (string, error) {

	if options.SkipFields != 0 {

		truncateStr, err := truncate.TruncateFields(curString, options.SkipFields)

		if err != nil {
			return curString, err
		}

		curString = truncateStr
	}

	return curString, nil
}

func CheckSkipChars(curString string, options options.Options) (string, error) {

	if options.SkipChars != 0 {

		truncateStr, err := truncate.TruncateSymbols(curString, options.SkipChars)

		if err != nil {
			return curString, err
		}

		curString = truncateStr
	}

	return curString, nil
}

func CheckUniqueFlags(line, prevStringToReturn, prevString string, options options.Options, isStringRepeat, isFirstString bool, repeatCount int) (string, bool) {

	if options.Repeated {

		if isStringRepeat && !isFirstString {

			return prevStringToReturn, true
		}

	} else if options.Unique {

		if !isStringRepeat && !isFirstString {

			return prevStringToReturn, true
		}

	} else if options.Count {

		if !isFirstString {
			prefix := "    " + strconv.Itoa(repeatCount) + " "
			prevStringToReturn = utils.ConcatStrings(&prefix, &prevString)

			return prevStringToReturn, true
		}

	} else {

		return line, true
	}

	return "", false
}
