package actions

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// GenerateToc generates the TOC for a given file.
func GenerateToc(lines []string, maxDepth int) []string {
	// Regular Expressions
	regTitle := regexp.MustCompile(`#+(.*)`)
	regHeading := regexp.MustCompile(`(?m)^[^\s]*#\s`)

	prevHeading := 0
	store := []string{"Table of contents:\n"}

	for _, line := range lines {
		match := regHeading.MatchString(line)

		if match {
			// Grab heading
			heading := regTitle.FindStringSubmatch(line)

			currentHeading := strings.Count(heading[0], "#")

			// if a maximum depth is set, and the current heading's depth is
			// over that maximum, just ignore it.
			if maxDepth > 0 && maxDepth < currentHeading {
				continue
			}

			lowerHeading := strings.ToLower(heading[1])
			kebabHeading := strings.ReplaceAll(strings.TrimSpace(lowerHeading), " ", "-")
			title := strings.TrimSpace(heading[1])

			// if no previous heading or title is found.
			if prevHeading == 0 || currentHeading == 1 {
				prevHeading = currentHeading
				store = append(store, fmt.Sprintf("- [%s](#%s)", title, kebabHeading))
				continue
			}

			if prevHeading == currentHeading {
				lastitem := store[len(store)-1]

				spaces := countLeadingSpaces(lastitem)
				store = append(store, fmt.Sprintf("%*s- [%s](#%s)", spaces, "", title, kebabHeading))
			}

			if prevHeading > currentHeading || prevHeading < currentHeading {
				spaces := currentHeading * 2
				store = append(store, fmt.Sprintf("%*s- [%s](#%s)", spaces, "", title, kebabHeading))
			}

		}
	}

	// Handle no headings.
	if len(store) == 1 {
		fmt.Println("No Headings found in the file.❌")
		os.Exit(0)
	}

	return store
}

// countLeadingSpaces counts the leading spaces.
func countLeadingSpaces(line string) int {
	return len(line) - len(strings.TrimLeft(line, " "))

}
