package src

import (
	"fmt"
	"os"
	"strings"
)

// Function GetBanner validates the banner and reads it's content into a string slice.
func GetBanner(bannerName string) ([]string, error) {
	fileSize := 0
	switch bannerName {
	case "shadow", "shadow.txt":
		fileSize = 7463
		bannerName = "banners/shadow.txt"
	case "standard", "standard.txt":
		fileSize = 6623
		bannerName = "banners/standard.txt"
	case "thinkertoy", "thinkertoy.txt":
		fileSize = 5558
		bannerName = "banners/thinkertoy.txt"
	default:
		return nil, fmt.Errorf("banner file provided is not in scope: use standard; thinkertoy or shadow")
	}
	fileContent, err := os.ReadFile(bannerName)
	if err != nil {
		return nil, fmt.Errorf("error reading %v: %v", bannerName, err)
	}
	if len(fileContent) == 0 {
		return nil, fmt.Errorf("%v is an empty file", bannerName)
	}
	// Handle if file is tampered with locally.
	if len(fileContent) != fileSize {
		return nil, fmt.Errorf("local banner file %v, has been corrupted", bannerName)
	}
	var fileContents []string
	if bannerName == "banners/thinkertoy.txt" {
		fileContents = strings.Split(string(fileContent), "\r\n")
	} else {
		fileContents = strings.Split(string(fileContent), "\n")
	}
	return fileContents, nil
}

// ReadBannerContent populates the banner content in a map with ascii number as key and each figure as it's value.
func ReadBannerContent(file []string) map[int][]string {
	asciiMap := make(map[int][]string)
	lineCount := 0
	initialAsciiValue := 32
	for _, line := range file {
		if line == "" {
			continue
		}
		asciiMap[initialAsciiValue] = append(asciiMap[initialAsciiValue], line)
		lineCount++
		if lineCount%8 == 0 {
			initialAsciiValue++
		}
	}
	return asciiMap
}

// PrintAsciiArt writes the ASCII graphical representation of characters in a string.
func PrintAsciiArt(inputString string, asciiMap map[int][]string) string {
	var output strings.Builder
	words := strings.Split(inputString, "\\n")
	// checks for words made of entirely strings.
	for j, word := range words {
		if word == "" {
			if j < len(words)-1 {
				output.WriteString("\n")
			}
			continue
		}
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
	return output.String()
}
