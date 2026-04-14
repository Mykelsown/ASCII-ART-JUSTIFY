package main

import (
	asciiart "asciiartjustify/MethodsAndTesting"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		flag       strings.Builder
		input      strings.Builder
		position   strings.Builder
		formatType strings.Builder
	)

	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Print(printUsageMessage())
		return
	}

	if len(os.Args) == 2 {
		input.WriteString(os.Args[1])
		fmt.Println(asciiart.FormatPrinter(input.String()))
		return
	} else if len(os.Args) == 4 {
		flag.WriteString(os.Args[1][:8])
		if flag.String() != "--align=" {
			fmt.Print(printUsageMessage())
			return
		}

		input.WriteString(os.Args[2])
		formatType.WriteString(os.Args[3])
		if formatType.String() != "standard" && formatType.String() != "thinkertoy" && formatType.String() != "shadow" {
			fmt.Printf("invalid banner type: %v", formatType.String())
			return
		}

		position.WriteString(os.Args[1][8:])
		switch position.String() {
		case "right", "left", "center", "justify":
			fmt.Println(asciiart.AlignArt(position.String(), input.String(), formatType.String()))
			return
		default:
			fmt.Printf("inavlid alignment position: %v\n", position.String())
			return
		}

	}
}

func printUsageMessage() string {
	return fmt.Sprintln("Usage: go run . [OPTION] [STRING] [BANNER]\nExample: go run . --align=right something standard")
}
