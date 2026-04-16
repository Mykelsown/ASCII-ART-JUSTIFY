package asciiart

import (
	"log"
	"strings"
)

func FormatPrinter(input string) string {
	// FormatPrinter always uses the standard banner.
	// FileHandler now requires a style argument — pass "standard" explicitly.
	contentRead, readingStatus := FileHandler("standard")

	if !readingStatus {
		log.Fatal("there was an error while reading file")
	}

	lines := strings.Split(string(contentRead), "\n")
	words := strings.Split(input, "\\n")

	var result strings.Builder

	for _, word := range words {
		if word == "" {
			result.WriteString("\n")
			continue
		}

		for row := 0; row < 8; row++ {
			var lineBuilder strings.Builder

			for _, char := range word {
				if char < 32 || char > 126 {
					continue
				}

				charIndex := int(char) - 32
				lineIndex := charIndex*9 + 1 + row

				if lineIndex < len(lines) {
					lineBuilder.WriteString(lines[lineIndex])
				}
			}

			result.WriteString(lineBuilder.String())
			result.WriteString("\n")
		}
	}

	output := result.String()
	if len(output) > 0 && output[len(output)-1] == '\n' {
		output = output[:len(output)-1]
	}

	return output
}