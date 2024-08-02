package src

import (
	"fmt"
	"strings"
)

// PrintAsciiArt returns a string of ASCII graphical representation of characters
func PrintAsciiArt(inputString, substr, color string, asciiMap map[int][]string) {
	words := strings.Split(inputString, "\n")
	reset := "\x1b[0m"


	// fmt.Println(inputString)

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
							fmt.Printf(color + figure[i] + reset)
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
