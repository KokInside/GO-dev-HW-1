package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"uniq/options"
	"uniq/uniq"
	"uniq/utils"
)

func main() {

	// объявление флагов
	cFlag := flag.Bool("c", false, "Count the number of occurrences of a string in the input data.")

	dFlag := flag.Bool("d", false, "Output only those lines that are repeated in the input data.")

	uFlag := flag.Bool("u", false, "Output only those lines that are not repeated in the input data.")

	iFlag := flag.Bool("i", false, "Ignore letter case.")

	fFlag := flag.Int("f", 0, "Ignore the first [num_fields] fields in a row. A field in a row is a non-empty set of characters separated by a space.")

	sFlag := flag.Int("s", 0, "Ignore the first [num_chars] characters in a string. When used with the -f option, the first characters after [num_fields] fields are taken into account (ignoring the space separator after the last field).")

	helpFlag := flag.Bool("help", false, "Display this help and exit")

	flag.Parse()

	if *helpFlag {
		fmt.Println("Usage: uniq [-c | -d | -u] [-i] [-f num] [-s chars] [INPUT_FILE [OUTPUT_FILE]]")
		fmt.Println("Filter adjacent matching lines from INPUT_FILE (or standard input),\nwriting to OUTPUT_FILE (or standard output).")
		fmt.Println("")
		fmt.Println("Options:")
		flag.PrintDefaults()
		return
	}

	options := options.Options{
		Count:      *cFlag,
		Repeated:   *dFlag,
		Unique:     *uFlag,
		IgnoreCase: *iFlag,
		SkipFields: *fFlag,
		SkipChars:  *sFlag,
	}

	// поток ввода
	var reader io.Reader

	// поток вывода
	var writer io.Writer

	// Определение потока ввода-вывода данных
	switch len(flag.Args()) {
	case 0:
		reader = os.Stdin

		writer = os.Stdout

	case 1:
		inputFile, inputFileError := os.Open(flag.Arg(0))

		if inputFileError != nil {
			fmt.Println(inputFileError.Error())
			return
		}

		defer inputFile.Close()

		reader = inputFile

		writer = os.Stdout

	case 2:
		inputFile, inputFileError := os.Open(flag.Arg(0))
		outputFile, outputFileError := os.OpenFile(flag.Arg(1), os.O_WRONLY|os.O_TRUNC, 0777)

		if inputFileError != nil {
			fmt.Println(inputFileError.Error())
			return
		}

		defer inputFile.Close()

		if outputFileError != nil {

			outputFile, outputFileError = os.Create(flag.Arg(1))

			if outputFileError != nil {
				fmt.Println(outputFileError.Error())
				return
			}
		}

		defer outputFile.Close()

		writer = outputFile

		reader = inputFile

	default:
		fmt.Println("Too many arguments.")
		flag.PrintDefaults()
		return
	}

	lines, err := utils.ReadLines(reader)

	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := uniq.Uniq(lines, options)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i, uniqLine := range result {

		writer.Write([]byte(uniqLine))

		if i != len(result)-1 {

			writer.Write([]byte("\n"))
		}
	}
}
