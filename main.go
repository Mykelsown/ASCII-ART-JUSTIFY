package main

import (
	asciiart "asciiartjustify/MethodsAndTesting"
	"fmt"
	"os"
	"strings"
)

// projects entry point, and where all argument passed by users are being validated for correctness
func main() {
	var (
		flag       strings.Builder
		input      strings.Builder
		position   strings.Builder
		formatType strings.Builder
	)

	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Print(usageMessage())
		return
	}

	// Base case: just a string, no flag, no banner → standard banner, no alignment
	if len(os.Args) == 2 {
		input.WriteString(os.Args[1])
		fmt.Println(asciiart.FormatPrinter(input.String()))
		return
	}

	// 3 args: could be --align=<type> + string (banner defaults to standard)
	if len(os.Args) == 3 {
		flag.WriteString(os.Args[1][:8])
		if flag.String() != "--align=" {
			fmt.Print(usageMessage())
			return
		}

		position.WriteString(os.Args[1][8:])
		input.WriteString(os.Args[2])

		switch position.String() {
		case "right", "left", "center", "justify":
			// AlignArt prints directly and blocks for resize — do NOT wrap in fmt.Println
			asciiart.AlignArt(position.String(), input.String(), "standard")
			return
		default:
			fmt.Printf("invalid alignment position: %v\n", position.String())
			return
		}
	}

	// 4 args: --align=<type> + string + banner
	if len(os.Args) == 4 {
		flag.WriteString(os.Args[1][:8])
		if flag.String() != "--align=" {
			fmt.Print(usageMessage())
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
			// AlignArt prints directly and blocks for resize — do NOT wrap in fmt.Println
			asciiart.AlignArt(position.String(), input.String(), formatType.String())
			return
		default:
			fmt.Printf("invalid alignment position: %v\n", position.String())
			return
		}
	}
}

func usageMessage() string {
	return fmt.Sprintln("Usage: go run . [OPTION] [STRING] [BANNER]\nExample: go run . --align=right something standard")
}