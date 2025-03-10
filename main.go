package main

import (
	"ascii-art-color/src"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Handle if no fileName is provided as argument
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
		return

	}
	// Define a flag named "color"

	colorPtr := flag.String("color", "something", "letters to be colored")

	// Parse the command-line flags
	flag.Parse()

	// Access values of the flags

	colorName := *colorPtr

	// obtaining user input
	lettersToColor := os.Args[2]
	words := os.Args[3]
	bannerFileName := "banners/standard.txt"

	// obtainng RGB values for color using 'PrintColored' function
	r, g, b := src.PrintColored(colorName)

	// Concatenate everything into a string
	foregroundColor := fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
	reset_color := "\x1b[0m"

	// Regular expression pattern to match the word
	re := regexp.MustCompile(regexp.QuoteMeta(lettersToColor))

	// Find all matches
	words = src.ReadInput(words) // + "\n"
	// fmt.Println(words)
	matches := re.FindAllStringSubmatchIndex(words, -1)

	if len(matches) == 0 {
		fmt.Print(src.PrintAsciiArt(words, bannerFileName))
		// fmt.Printf("%s", result)
	} else {
		lastIndex := 0

		for _, match := range matches {
			start, end := match[0], match[1]

			if start > lastIndex {
				fmt.Print(src.PrintAsciiArt(words[lastIndex:start], bannerFileName))
			}

			fmt.Print(foregroundColor, src.PrintAsciiArt(words[start:end], bannerFileName), reset_color)

			lastIndex = end
		}
		if lastIndex < len(words) {
			fmt.Print(src.PrintAsciiArt(words[lastIndex:], bannerFileName))
		}
	}

}
