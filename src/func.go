package src

import (
	"strings"
)

// PrintAsciiArt returns a string of ASCII graphical representation of characters
func PrintAsciiArt(inputString, substr, color string, asciiMap map[int][]string) {
	// Replace escaped newlines with actual newlines in both inputString and substr
	// inputString = strings.ReplaceAll(inputString, "\\n", "\n")
	// substr = strings.ReplaceAll(substr, "\\n", "\n")

	// inputLines := strings.Split(inputString, "\n")
	// substrLines := strings.Split(substr, "\n")
	// reset := "\x1b[0m"
	// fmt.Println(substr)
	// fmt.Println(inputString)

	// for lineIndex, line := range inputLines {
	// 	if line == "" {
	// 		fmt.Println()
	// 		continue
	// 	}

	indexes := []int{}

	for i := 0; i < len(inputString); i++ {
			if strings.Contains(inputString[i:], substr) && inputString[i:i+len(substr)] == substr {
				indexes = append(indexes, i)
				i += len(substr)-1
			}
	}

	// for _, num := range indexes {
	// 	println(num)
	// }

	for i:= 0; i < len(inputString);i++{
		for j:= 0; j < len(indexes);j++{
			if indexes[j] == i {
				char := inputString[i]
				println(string(char))
				j++
				i++
			}

		}
	}
		

	// 	for i := 0; i < 8; i++ {
	// 		j := 0
	// 		for j < len(line) {
	// 			matchFound := false
	// 			for _, subLine := range substrLines {
	// 				if strings.HasPrefix(line[j:], subLine) {
	// 					// Color the matching part
	// 					for k := 0; k < len(subLine); k++ {
	// 						char := line[j+k]
	// 						figure, ok := asciiMap[int(char)]
	// 						if ok {
	// 							fmt.Printf(color + figure[i] + reset)
	// 						}
	// 					}
	// 					j += len(subLine)
	// 					matchFound = true
	// 					break
	// 				}
	// 			}
	// 			if !matchFound {

	// 				char := line[j]
	// 				figure, ok := asciiMap[int(char)]
	// 				if ok {
	// 					fmt.Print(figure[i])
	// 				}
	// 				j++
	// 			}
	// 		}
	// 		fmt.Println()
	// 	}

	// 	// Add a newline between words
	// 	if lineIndex < len(inputLines)-1 {
	// 		fmt.Println()
	// 	}
	// }
}
