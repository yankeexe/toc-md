package utils

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

// ValidateFile makes sure the user input file is valid.
func ValidateFile(fileLocation string, extension string) cli.ExitCoder {
	fileInfo, err := os.Stat(fileLocation)
	if os.IsNotExist(err) {
		fmt.Printf("File does not exist %s❌", fileLocation)
		os.Exit(0)
	}

	// Check for file extension.
	if extension != ".md" {
		fmt.Println("You can only generate table of contents for markdown files having '.md' extension.❌")
		os.Exit(0)
	}

	// Check for empty file.
	if fileInfo.Size() == 0 {
		fmt.Println("File is empty.❌")
		os.Exit(0)
	}
	return nil
}
