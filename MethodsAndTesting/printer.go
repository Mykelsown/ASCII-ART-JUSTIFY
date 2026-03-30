package asciiart

import (
	"log"
	"strings"
)

func FormatPrinter(input string) string {
	contentRead, readingStatus := FileHandler()

	// Error handling for failure to read file. If the file cant be read:
	// 1. Print out the message
	// 2. exit the function
	if !readingStatus {
		log.Fatal("there was an error while reading file")
	}

	// Split banner into lines
	lines := strings.Split(string(contentRead), "\n")

	// Split input by literal \n (not actual newline)
	words := strings.Split(input, "\\n")

	var result strings.Builder

	// Process each word/line
	for _, word := range words {
		// If the line is empty, add a single newline
		if word == "" {
			result.WriteString("\n")
			continue
		}

		// Each ASCII character is 8 lines tall
		for row := 0; row < 8; row++ { // Defines charcaters height
			var lineBuilder strings.Builder

			// Loop through each character
			for _, char := range word {
				// Check ASCII range (32 = space, 126 = ~) i.e chararcter is a printable ASCII character
				if char < 32 || char > 126 {
					continue
				}

				charIndex := int(char) - 32 // Find character position in banner
				lineIndex := charIndex*9 + 1 + row // Each character takes 9 lines (8 art lines + 1 empty separator)

				// what this says is that if the character line on in the file is less thn 855, the condition is valid.
				if lineIndex < len(lines) {
					lineBuilder.WriteString(lines[lineIndex]) 
				}
			}

			// Add the line to result with newline
			result.WriteString(lineBuilder.String())
			result.WriteString("\n")
		}
	}

	// Remove the very last newline to match expected output
	output := result.String()
	if len(output) > 0 && output[len(output)-1] == '\n' {
		output = output[:len(output)-1]
	}

	return output
}
