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
func HandleArgs() (string, string,string, string, string, error) {
	var subStr, mainStr, color, outputFile, bannerFile string
	errMessage := "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\""

	

	flagVar := flag.String("output", "", "name of the file to contain the output")
	colorflag := flag.String("color", "", "name of the color to use in printing the output")

	flag.Parse()

	outputFile = *flagVar
	color = *colorflag
	nonFlag := flag.Args()

	if outputFile == "" {
		switch {
		case len(nonFlag) == 1:
			return nonFlag[0],nonFlag[0], color,"", "standard", nil
		case len(nonFlag) == 2:
			return nonFlag[0],nonFlag[1],color, "","standard", nil
		case len(nonFlag) == 3: 
			return nonFlag[0], nonFlag[1],color,"",nonFlag[2],nil
		}
	} else  if color == ""{
		switch  {
		case len(nonFlag) == 1:
			return nonFlag[0],nonFlag[0], "",outputFile, "standard", nil
		case len(nonFlag) != 1:
			return "", "", "","","", errors.New(errMessage)
		}

	} else if color == "" && outputFile == "" {
		switch {
		case len(nonFlag) == 1:
			return nonFlag[0],nonFlag[0], "","", "standard", nil			
		}
	} else {
		switch {
		case len(nonFlag) == 0:
			return "", "", "","","", errors.New(errMessage)
		case len(nonFlag) == 1:
			return nonFlag[0],nonFlag[0],color, outputFile, "standard", nil
		case len(nonFlag) == 2:
			return nonFlag[0],nonFlag[1],color, outputFile,"standard", nil
		case len(nonFlag) == 3:
			return nonFlag[0],nonFlag[1],color, outputFile,nonFlag[2], nil
		}
	}
	return subStr, mainStr, color, outputFile, bannerFile, errors.New(errMessage)
}
