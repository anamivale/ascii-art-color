package src

import (
	"fmt"
	"os"
	"strings"
)

// CheckError
func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
}

// CheckBannerName returns a string with the valid name of a banner file
func CheckBannerName(fileName string) string {
	bannerFileName := ""
	switch fileName {
	case "thinkertoy":
		bannerFileName = "thinkertoy.txt"
	case "standard":
		bannerFileName = "standard.txt"
	case "shadow":
		bannerFileName = "shadow.txt"
	default:
		bannerFileName = ""
		fmt.Println("Valid template names are :[standard][thinkertoy][shadow]")
	}
	return bannerFileName
}

// ReadInput returns a string of valid characters
func ReadInput(str string) string {
	// handling special characters
	str = strings.ReplaceAll(str, "\n", "\\n")
	str = strings.ReplaceAll(str, "\t", "    ")

	// loop to check characters are within our range (32 - 126)
	for _, value := range str {
		if value > 126 || value < 32 {
			fmt.Println("Provide characters within the ASCII range (26-126)")
			os.Exit(1)
		}
	}
	return str
}

// PrintAsciiArt returns a string of ASCII graphical representation of characters
func PrintAsciiArt(inputString string, asciiFile string) string {
	asciiMap := ReadBannerContent(asciiFile)

	var output strings.Builder

	words := strings.Split(inputString, "\\n")

	// checks for words made of entirely strings.
	for j, word := range words {
		if word == "" {
			if j < len(words)-1 {
				output.WriteString("\n")
			}
		} else {
			lines := make([]string, 8)
			// loop through character of a word, attaching each line of map for that char to a slice of length 8
			for _, char := range word {
				figure, ok := asciiMap[int(char)]
				if ok {
					for i := 0; i < 8; i++ {
						lines[i] += figure[i]
					}
				}
			}
			// write to the output string
			for i := 0; i < 8; i++ {
				output.WriteString((lines[i] + "\n"))
			}
		}
	}
	return output.String()
}

// ReadBannerContent takes a fileName as a string and returns a map with ascii number as key and each figure as it's value
func ReadBannerContent(text string) map[int][]string {
	var file1 []string
	file, err := os.ReadFile(text)
	checkError(err)
	// slice of bytes to slice of strings
	if text == "thinkertoy.txt" {
		file1 = strings.Split(string(file), "\r\n")
	} else {
		file1 = strings.Split(string(file), "\n")
	}
	// creating a map
	AsciiMap := make(map[int][]string)
	lineCount := 0
	initialAsciiValue := 32
	for i := 0; i < len(file1); i++ {
		if file1[i] == "" {
			continue
		}
		AsciiMap[initialAsciiValue] = append(AsciiMap[initialAsciiValue], file1[i])
		lineCount++

		if lineCount%8 == 0 {
			initialAsciiValue++
		}

	}
	return AsciiMap
}
