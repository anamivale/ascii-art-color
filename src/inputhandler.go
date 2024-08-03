package src

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

// ReadInput returns a string of valid characters.
func ReadInput(text string) (string, error) {
	println(text)
	text = strings.ReplaceAll(text, "\n", "\\n")
	text = strings.ReplaceAll(text, "\\t", "   ")
	for _, char := range text {
		if char < 32 || char > 126 {
			if char == 10 || char == 13 {
				continue
			} else {
				return "", fmt.Errorf(" char is not valid")
			}
		}
	}
	println(text)
	return text, nil
}

// SortArgs takes in a string of arguments and flags sorts and stores them in there respective variables.
func HandleArgs() (string, string, string, string, error) {
	var inputStr, subStr, color, bannerFile string
	errMessage := "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\""

	args := os.Args[1:]

	knownFlags := map[string]bool{
		"color": true,
	}

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--") || strings.HasPrefix(arg, "-") {
			flagName := strings.Split(strings.TrimPrefix(arg, "--"), "=")

			if !knownFlags[flagName[0]] {
				return "", "", "", "", errors.New(errMessage)
			}
		}
		if strings.HasPrefix(arg, "--") {
			parts := strings.SplitN(arg, "=", 2)
			if len(parts) != 2 {
				return "", "", "", "", errors.New(errMessage)
			}
		}

	}

	colorFlagVar := flag.String("color", "", "name of the color")

	flag.Parse()

	color = *colorFlagVar

	nonFlag := flag.Args()

	if color == "" {
		switch {
		case len(nonFlag) == 1:
			println(nonFlag[0])
			return args[0], args[0], "white", "standard", nil
		case len(nonFlag) == 2:
			return args[0], args[0], "white", args[1], nil
		}
	} else {
		switch {
		case len(nonFlag) == 1:
			return nonFlag[0], nonFlag[0], color, "standard", nil
		case len(nonFlag) == 2:
			if (nonFlag[1] == "standard" || nonFlag[1] == "thinkertoy" || nonFlag[1] == "shadow") && !(strings.Contains(nonFlag[1], nonFlag[0])) {
				return nonFlag[0], nonFlag[0], color, nonFlag[1], nil
			} else {
				return nonFlag[0], nonFlag[1], color, "standard", nil
			}
		case len(nonFlag) == 3:
			return nonFlag[0], nonFlag[1], color, nonFlag[2], nil
		}
	}

	return subStr, inputStr, color, bannerFile, errors.New(errMessage)
}
