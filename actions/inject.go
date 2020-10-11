package actions

import (
	"fmt"
	"log"
	"os"
)

// InjectToc injects the generated TOC to the file.
func InjectToc(lines []string, toc []string, file *os.File) {
	toc[len(toc)-1] = toc[len(toc)-1] + "\n"

	toc = append(toc, lines...)

	file.Seek(int64(0), 0)

	for _, heading := range toc {
		if _, err := file.WriteString(heading + "\n"); err != nil {
			log.Fatalf("failed writing to file: %s", err)
		}
	}
	fmt.Println("Table of contents added to file.✅")
}
