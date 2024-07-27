package src

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

// ReadInput returns a string of Valided input characters.
func HandleInput(str string) string {
	// Handling special characters
	str = strings.ReplaceAll(str, "\n", "\\n")
	str = strings.ReplaceAll(str, "\\t", "    ")

	// Loop to check characters are within our range (32 - 126)
	for _, value := range str {
		if value < 32 || value > 126 {
			fmt.Println("Provide characters within the ASCII range (32-126)")
			os.Exit(1)
		}
	}
	return str
}

// SortArgs takes in a string of arguments and flags sorts and stores them in there respective variables.
func HandleArgs() (string, string, string, error) {
	var inputStr, outputFile, bannerFile string
	errMessage := "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEx: go run . --output=<filename.txt> something standard"

	flagVar := flag.String("output", "", "name of the file to contain the output")

	flag.Parse()

	outputFile = *flagVar
	nonFlag := flag.Args()

	if outputFile == "" {
		switch {
		case len(nonFlag) == 1:
			return nonFlag[0], "", "standard", nil
		case len(nonFlag) == 2:
			return nonFlag[0], "", nonFlag[1], nil
		}
	} else {
		switch {
		case len(nonFlag) == 0:
			return "", "", "", errors.New(errMessage)
		case len(nonFlag) == 1:
			return nonFlag[0], outputFile, "standard", nil
		case len(nonFlag) == 2:
			return nonFlag[0], outputFile, nonFlag[1], nil
		}
	}
	return inputStr, outputFile, bannerFile, errors.New(errMessage)
}
