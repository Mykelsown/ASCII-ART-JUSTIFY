package asciiart

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"unsafe"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getTerminalWidth() int {
	var sz winsize
	syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		syscall.TIOCGWINSZ,
		uintptr(unsafe.Pointer(&sz)),
	)
	if sz.Col == 0 {
		return 80
	}
	return int(sz.Col)
}

func wordToRows(word string, bannerLines []string) []string {
	rows := make([]string, 8)
	for row := 0; row < 8; row++ {
		var sb strings.Builder
		for _, ch := range word {
			if ch < 32 || ch > 126 {
				continue
			}
			charIdx := int(ch) - 32
			lineIdx := charIdx*9 + 1 + row
			if lineIdx < len(bannerLines) {
				sb.WriteString(bannerLines[lineIdx])
			}
		}
		rows[row] = sb.String()
	}
	return rows
}

func buildAlignedOutput(position, sentence, style string) string {
	termWidth := getTerminalWidth()

	// Load the banner file once
	content, ok := FileHandler(style)
	if !ok {
		return "error: could not read banner file: " + style
	}
	bannerLines := strings.Split(string(content), "\n")

	// Support multi-line input via literal \n in the string
	inputLines := strings.Split(sentence, "\\n")

	var result strings.Builder

	for _, inputLine := range inputLines {
		// Empty segment from \n → emit a blank line
		if inputLine == "" {
			result.WriteString("\n")
			continue
		}

		words := strings.Fields(inputLine) 

		if position == "justify" && len(words) > 1 {
			
			wordArt := make([][]string, len(words))
			wordWidths := make([]int, len(words))

			for i, w := range words {
				wordArt[i] = wordToRows(w, bannerLines)
				wordWidths[i] = len(wordArt[i][0]) 
			}

			totalWordWidth := 0
			for _, w := range wordWidths {
				totalWordWidth += w
			}

			gapCount := len(words) - 1
			totalGapSpace := termWidth - totalWordWidth

			if totalGapSpace < 0 {
				totalGapSpace = 0
			}

			baseSpace := totalGapSpace / gapCount 
			remainder := totalGapSpace % gapCount 

			for row := 0; row < 8; row++ {
				var sb strings.Builder
				for i, art := range wordArt {
					sb.WriteString(art[row])
					
					if i < len(wordArt)-1 {
						gapSize := baseSpace
						if i < remainder {
							gapSize++
						}
						sb.WriteString(strings.Repeat(" ", gapSize))
					}
				}
				result.WriteString(sb.String())
				result.WriteString("\n")
			}

		} else {
			artRows := wordToRows(inputLine, bannerLines)

			for row := 0; row < 8; row++ {
				line := artRows[row]
				lineWidth := len(line)

				pad := 0
				switch position {
				case "right":
					pad = termWidth - lineWidth
				case "center":
					pad = (termWidth - lineWidth) / 2
				}

				if pad < 0 {
					pad = 0
				}

				result.WriteString(strings.Repeat(" ", pad))
				result.WriteString(line)
				result.WriteString("\n")
			}
		}
	}

	output := result.String()
	if len(output) > 0 && output[len(output)-1] == '\n' {
		output = output[:len(output)-1]
	}
	return output
}

func AlignArt(position, sentence, style string) {
	draw := func() {
		output := buildAlignedOutput(position, sentence, style)
		fmt.Print("\033[H\033[2J\033[3J" + output)
	}

	// Initial draw on startup
	draw()

	// Register a channel to receive SIGWINCH (terminal resize events)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGWINCH)

	for range sigCh {
		draw()
	}
}
