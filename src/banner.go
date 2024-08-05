package src

import (
	"fmt"
	"os"
	"strings"
)

// GetBanner validates the banner and reads it's content into a string slice.
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

// ReadBannerContent takes a fileName as a string and returns a map with ascii number as key and each figure as it's value
func ReadBannerContent(bannerContents []string) map[int][]string {
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
