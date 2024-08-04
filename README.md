# ASCII ART COLOR
## Overview

This project provides a Go-based solution to generate ASCII art from input text, with the capability to apply colors to specific substrings within the text. The ASCII art representation is construction is constructed using predefined templates, and supports a variety of colors models. This tool can be particularly useful for creating visually appealing command-line ouputs or text for various applications.

## Features

- **Generate ASCII Art**: Convert plain text into stylized ASCII art.
- **Color Substrings**: Apply specific colors to substrings within the ASCII art.
- **Supports Different Banners**: Standard, Shadow, and Thinkertoy.
- **Supports Different color Models**: Hex, RGB and Ansi.
- **Handles Non-ASCII Characters**: Ensures all input characters are valid ASCII.

## Project Structure
```
.
├── main.go
├── colors
│   ├── convertcolors_test.go
│   ├── convertcolors.go
├── src
│   ├── banner_test.go
│   ├── banner.go
│   ├── func_test.go
│   ├── func.go
│   ├── asciiart_test.go
│   ├── asciiart.go
├── banners
│   ├── standard.txt
│   ├── shadow.txt
│   ├── thinkertoy.txt
├── go.mod
└── README.md
```

## Setup

### Prerequisites
- Go 1.18+ installed on your machine.
- Clone the repository:
```bash
git clone https://learn.zone01kisumu.ke/git/vmuhembe/ascii-art-color.git 
cd ascii-art-color
```
### Installing Dependancies
Run the following command to download the required Go modules:
```bash
go mod tidy
```

## Usage

### Command-Line 
```bash
Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring to be colored> "something"
```
### Example

To generate ASCII art for the word "HELLO" and color the substring "LO" in red:

![Preview of the Generated Colored ASCII-Art](img/example1.png)