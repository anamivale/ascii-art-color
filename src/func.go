package src

import (
	"fmt"
	"strings"
)

// PrintAsciiArt returns a string of ASCII graphical representation of characters
func PrintAsciiArt(inputString, substr, color string, asciiMap map[int][]string) {
	words := strings.Split(inputString, "\n")
	reset := "\033[0m"
	color = "\033[0m"

	for _, word := range words {
		if word == "" {
			fmt.Println()
			continue
		}

		for i := 0; i < 8; i++ {
			j := 0
			for j < len(word) {
				if j <= len(word)-len(substr) && word[j:j+len(substr)] == substr {
					for k := 0; k < len(substr); k++ { // range for
						char := word[j+k]
						figure, ok := asciiMap[int(char)]
						if ok {
							fmt.Printf("\033[31m" + figure[i] + reset)
						}
					}
					j += len(substr)
				} else {
					char := word[j]
					figure, ok := asciiMap[int(char)]
					if ok {
						fmt.Print(figure[i])
					}
					j++
				}
			}

			fmt.Println()
		}
	}
}

// ReadBannerContent takes a fileName as a string and returns a map with ascii number as key and each figure as it's value
func ReadBannerContent(bannerContents []string) map[int][]string {
	// var file1 []string
	// creating a map
	AsciiMap := make(map[int][]string)
	lineCount := 0
	initialAsciiValue := 32
	for i := 0; i < len(bannerContents); i++ {
		if bannerContents[i] == "" {
			continue
		}
		AsciiMap[initialAsciiValue] = append(AsciiMap[initialAsciiValue], bannerContents[i])
		lineCount++

		if lineCount%8 == 0 {
			initialAsciiValue++
		}

	}
	return AsciiMap
}

