package main

import (
	asciiart "asciiartjustify/MethodsAndTesting"
	"fmt"
	"os"
	"strings"
)

// projects entry point, and where all argumment passed by users are being validated for correctness
func main() {
	// This is stores all the prerequsites that will be needed in the program i.e all the arguments collected from the user
	var (
		flag       strings.Builder
		input      strings.Builder
		position   strings.Builder
		formatType strings.Builder
	)
	// checks if the corect amount of arguments program can handle is being passed in
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Print(usageMessage())
		return
	}

	// Runs for the base ascii, if the argument number satisfys the base
	if len(os.Args) == 2 {
		input.WriteString(os.Args[1])
		fmt.Println(asciiart.FormatPrinter(input.String()))
		return
	} else if len(os.Args) == 4 { //runs for the main justify, if the argument number satisfy it
		flag.WriteString(os.Args[1][:8])
		// flag type validation
		if flag.String() != "--align=" {
			fmt.Print(usageMessage())
			return
		}

		input.WriteString(os.Args[2])
		formatType.WriteString(os.Args[3])
		// validation for correct banner type
		if formatType.String() != "standard" && formatType.String() != "thinkertoy" && formatType.String() != "shadow" {
			fmt.Printf("invalid banner type: %v", formatType.String())
			return
		}

		position.WriteString(os.Args[1][8:])
		// Validation for correct banner type
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

// An helper function for the message if the argument format is wrong.
func usageMessage() string {
	return fmt.Sprintln("Usage: go run . [OPTION] [STRING] [BANNER]\nExample: go run . --align=right something standard")
}
