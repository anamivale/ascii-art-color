package main

import (
	"ascii-art-color/src"
	"fmt"
	"os"
)

func main() {
	// Handle if no fileName is provided as argument
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
		return
	}

	inputStr, subStr,color, bannerName, err := src.HandleArgs()
	

	bannerContent,err := src.GetBanner(bannerName)
	if err != nil {
		fmt.Println(err)
		return
	}
	asciiMap := src.ReadBannerContent(bannerContent)
	words, err  := src.ReadInput(inputStr) // + "\n"
	if err != nil {
		println(err)
		return
	}

	src.PrintAsciiArt(words, subStr, color, asciiMap)
	
}
