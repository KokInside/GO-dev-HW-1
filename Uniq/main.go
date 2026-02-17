package main

import (
	"fmt"

	"uniq/options"
	"uniq/uniq"
)

func main() {

	/*

		var cFlag = flag.Bool("c", false, "Count the number of occurrences of a string in the input data.")

		var dFlag = flag.Bool("d", false, "Output only those lines that are repeated in the input data.")

		var uFlag = flag.Bool("u", false, "Output only those lines that are not repeated in the input data.")

		var iFlag = flag.Bool("i", false, "Ignore letter case.")

		var fFlag = flag.Int("f", 0, "Ignore the first [num_fields] fields in a row. A field in a row is a non-empty set of characters separated by a space.")

		var sFlag = flag.Int("s", 0, "Ignore the first [num_chars] characters in a string. When used with the -f option, the first characters after [num_fields] fields are taken into account (ignoring the space separator after the last field).")

		var helpFlag = flag.Bool("help", false, "Display this help and exit")

		flag.Parse()

		if *helpFlag == true {
			flag.PrintDefaults()
			return
		}

		options := options.Options{C: *cFlag, D: *dFlag, U: *uFlag, I: *iFlag, F: *fFlag, S: *sFlag}

	*/

	///////////////////////////////////////////

	/*

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
					fmt.Println(flag.Arg(0)+":", "no such file")
					return
				}

				defer inputFile.Close()

				reader = inputFile

				writer = os.Stdout

			case 2:
				inputFile, inputFileError := os.Open(flag.Arg(0))
				outputFile, outputFileError := os.OpenFile(flag.Arg(1), os.O_WRONLY|os.O_TRUNC, 0777)

				if inputFileError != nil {
					fmt.Println(flag.Arg(0)+":", "no such file")
					return
				}

				defer inputFile.Close()

				if outputFileError != nil {

				outputFile, outputFileError = os.Create(flag.Arg(1))

				if outputFileError != nil {
					fmt.Println(flag.Arg(1)+":", "can not create file")
					return
				}
			}

			defer outputFile.Close()

			writer = outputFile

			reader = inputFile

		default:
			fmt.Println("Too many arguments")
			return
		}

	*/

	// считыватель
	// scanner := bufio.NewScanner(reader)

	/////////////////////////////////////////////////////////////////////

	/*

		lines, readErr := utils.ReadLines(reader)

		if readErr != nil {
			fmt.Println("Read error:", readErr.Error())
			return
		}

	*/

	lines := []string{
		// "I love music.",
		// "I love music.",
		// "I love music.",
		// "",
		// "I love music of Kartik.",
		// "I love music of Kartik.",
		// "Thanks.",
		// "I love music of Kartik.",
		// "I love music of Kartik.",

		// "I LOVE MUSIC.",
		// "I love music.",
		// "I LoVe MuSiC.",
		// "",
		// "I love MuSIC of Kartik.",
		// "I love music of kartik.",
		// "Thanks.",
		// "I love music of kartik.",
		// "I love MuSIC of Kartik.",

		// "We love music.",
		// "I love music.",
		// "They love music.",
		// "",
		// "I love music of Kartik.",
		// "We love music of Kartik.",
		// "Thanks.",

		"I love music.",
		"A love music.",
		"C love music.",
		"",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}

	options := options.Options{S: 1}

	result, uniqErr := uniq.Uniq(lines, options)

	if uniqErr != nil {
		fmt.Println("Uniq Error:", uniqErr.Error())
		return
	}

	for _, i := range result {
		fmt.Println(i)
		// writer.Write([]byte(i))
	}

}
