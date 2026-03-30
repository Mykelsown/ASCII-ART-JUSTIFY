package asciiart

import (
	"fmt"
	"os"
)

func FileHandler() ([]byte, bool) {
	// Read banner file
	data, err := os.ReadFile("banners/standard.txt")
	if err != nil {
		fmt.Println("Error")
		return []byte{}, false // returns an empty slice of byte and false if there's is an error
	}


	return data, true
}