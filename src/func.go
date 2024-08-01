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
		bannerFileName = "banners/thinkertoy.txt"
	case "standard":
		bannerFileName = "banners/standard.txt"
	case "shadow":
		bannerFileName = "banners/shadow.txt"
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
func PrintAsciiArt(inputString, substr, color, asciiFile string) {
	asciiMap := ReadBannerContent(asciiFile)

	// // runes := []rune(inputString)
	inputString = strings.ReplaceAll(inputString, "\\n", "\n")

	words := strings.Split(inputString, "\n")
	 reset := "\033[0m"
	color = "\033[0m"
	// lines := ""
	// println(inputString)
	// println(string(words))

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

				// word := inputString
				// word = strings.ReplaceAll(word, "\n", "\\n")
				// // println(i, ":", word)
				// // println(word)
				// for ind := 0; ind < len(word); ind++ {
				// 	char := word[ind]
				// 	// println(string(char))
				// 	if char == '\n' {
				// 		println("s")
				// 		continue
				// 	}
				// // println(string(char))
				// if ind <= len(word)-len(substr) && word[ind:ind+len(substr)] == substr {
				// 	for j := 0; j < len(substr); j++ { // range for
				// 		char := word[ind+j]
				// 		if char == '\n' {
				// 			println()
				// 			break
				// 		}

				// 		figure, ok := asciiMap[int(char)]
				// 		if ok {
				// 			fmt.Printf("\033[31m" + figure[i] + reset)
				// 		}
				// 	}
				// 	ind += len(substr) - 1
				// } else {
				// 	char := word[ind]
				// 	if char == '\n' {
				// 		println()
				// 		continue
				// 	}
				// 	figure, ok := asciiMap[int(char)]
				// 	if ok {
				// 		fmt.Print(figure[i])
				// 	}
				// }
			}
			fmt.Println()
		}
	}
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

func readMap(slice []string) string {
	str := ""
	for _, value := range slice {
		str += value
	}
	return str
}
