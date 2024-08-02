package src

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// Fuction HandleFile validates the output filename and returns it as a string if valid or an error if not valid.
func HandleFile(fileName string) (string, error) {
	errMessage := "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEx: go run . --output=<filename.txt> something standard"
	if !strings.HasSuffix(fileName, ".txt") {
		return "", errors.New(errMessage)
	}

	dirPath, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for _, char := range fileName {
		if (char == '.') || unicode.IsLetter(char) || unicode.IsNumber(char) {
			continue
		} else {
			return "", fmt.Errorf("error: file name cannot consist of character \"%v\"", string(char))
		}
	}
	filePath := filepath.Join(dirPath, fileName)

	return filePath, nil
}
