package main

import (
	asciiart "asciiartjustify/MethodsAndTesting"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		input      strings.Builder
		position   strings.Builder
		formatType strings.Builder
	)

	if len(os.Args) == 2 {
		input.WriteString(os.Args[1])
		if input.String() == "" {
			return
		}

		fmt.Println(asciiart.FormatPrinter(input.String()))
		return
	} else if len(os.Args) == 4 {
		position.WriteString(os.Args[1][8:])
		input.WriteString(os.Args[2])
		formatType.WriteString(os.Args[3])
		fmt.Println(asciiart.AlignArt(position.String(), input.String(), formatType.String()))
	}
}
