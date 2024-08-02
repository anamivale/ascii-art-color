package main

import (
	"fmt"
	"os"

	"ascii-art-color/colors"
	"ascii-art-color/src"
)

func main() {
	// Handle if no fileName is provided as argument
	if len(os.Args) < 2 || len(os.Args) > 5 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
		return
	}

	subStr,inputStr, color, bannerName, err := src.HandleArgs()
	if err != nil {
		fmt.Println(err)
		return
	}

	// handle banner
	bannerContent, err := src.GetBanner(bannerName)
	if err != nil {
		fmt.Println(err)
		return
	}
	asciiMap := src.ReadBannerContent(bannerContent)

	// handle input string
	words, err := src.ReadInput(inputStr) 
	if err != nil {
		println(err)
		return
	}
	// handle color
	// Concatenate everything into a string
	r, g, b, err := colors.PrintColored(color)
	if err != nil {
		fmt.Printf("%s is not available\n",color)
		return
	}

	foregroundColor := fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
	// println(words, subStr)

	src.PrintAsciiArt(words, subStr, foregroundColor, asciiMap)
}
