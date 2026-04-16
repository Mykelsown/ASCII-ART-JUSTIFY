package asciiart

import (
	"fmt"
	"os"
)

// FileHandler reads the banner file for the given style (e.g. "standard", "shadow", "thinkertoy").
// Returns the file contents as bytes and a boolean indicating success.
func FileHandler(style string) ([]byte, bool) {
	data, err := os.ReadFile("banners/" + style + ".txt")
	if err != nil {
		fmt.Println("Error: could not read banner file:", style)
		return []byte{}, false
	}

	return data, true
}