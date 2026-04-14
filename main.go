package main

import (
	asciiart "asciiartjustify/MethodsAndTesting"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var (
		input      strings.Builder
		position   strings.Builder
		formatType strings.Builder
	)

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run . [OPTION] [STRING] [BANNER]\nExample: go run . --align=right something standard")
	}

	if len(os.Args) == 2 {
		input.WriteString(os.Args[1])
		fmt.Println(asciiart.FormatPrinter(input.String()))
	} else if len(os.Args) == 4 {
		position.WriteString(os.Args[1][8:])
		input.WriteString(os.Args[2])
		formatType.WriteString(os.Args[3])
		fmt.Println(asciiart.AlignArt(position.String(), input.String(), formatType.String()))
	}
}
