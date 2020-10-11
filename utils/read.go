package utils

import (
	"fmt"
	"os"
)


// ReadFile reads the file passed by the user.
func ReadFile(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_RDWR, 0600)
	if err != nil {
		fmt.Printf("Error opening %s", err)
		os.Exit(0)
	}

	return file
}
