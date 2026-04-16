package asciiart

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"unsafe"
)

// ---------------------------------------------------------------------------
// Terminal Width
// ---------------------------------------------------------------------------

// winsize mirrors the kernel's winsize struct for TIOCGWINSZ.
// Defined manually because syscall.Winsize availability varies by Go version.
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

// ---------------------------------------------------------------------------
// Art Rendering Helpers
// ---------------------------------------------------------------------------

// wordToRows renders a single word (or full line) into 8 art rows using the
// pre-loaded banner lines. Returns a slice of 8 strings, one per art row.
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

// ---------------------------------------------------------------------------
// Alignment Logic
// ---------------------------------------------------------------------------

// buildAlignedOutput generates the complete aligned output string for the
// current terminal width. Called once on startup and again on every resize.
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

		words := strings.Fields(inputLine) // split on spaces, strips extras

		// ---------------------------------------------------------------
		// JUSTIFY — multiple words spread across the full terminal width
		// ---------------------------------------------------------------
		if position == "justify" && len(words) > 1 {
			// Step 1: render every word into its own 8-row art block
			wordArt := make([][]string, len(words))
			wordWidths := make([]int, len(words))

			for i, w := range words {
				wordArt[i] = wordToRows(w, bannerLines)
				wordWidths[i] = len(wordArt[i][0]) // visual width = len for ASCII
			}

			// Step 2: calculate how much gap space is available
			totalWordWidth := 0
			for _, w := range wordWidths {
				totalWordWidth += w
			}

			gapCount := len(words) - 1
			totalGapSpace := termWidth - totalWordWidth

			// Protect against negative space (word wider than terminal)
			if totalGapSpace < 0 {
				totalGapSpace = 0
			}

			baseSpace := totalGapSpace / gapCount // every gap gets at least this
			remainder := totalGapSpace % gapCount // leftover spaces → leftmost gaps first

			// Step 3: build each of the 8 art rows
			for row := 0; row < 8; row++ {
				var sb strings.Builder
				for i, art := range wordArt {
					sb.WriteString(art[row])
					// Add the gap after every word except the last
					if i < len(wordArt)-1 {
						gapSize := baseSpace
						if i < remainder { // leftmost gaps get the extra space
							gapSize++
						}
						sb.WriteString(strings.Repeat(" ", gapSize))
					}
				}
				result.WriteString(sb.String())
				result.WriteString("\n")
			}

			// ---------------------------------------------------------------
			// LEFT / RIGHT / CENTER — and JUSTIFY with a single word (= left)
			// ---------------------------------------------------------------
		} else {
			// Render the whole input line as one art block (spaces included)
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

	// Trim the trailing newline to match expected output format
	output := result.String()
	if len(output) > 0 && output[len(output)-1] == '\n' {
		output = output[:len(output)-1]
	}
	return output
}

// ---------------------------------------------------------------------------
// Entry Point — Print + Resize Loop
// ---------------------------------------------------------------------------

// AlignArt renders the ASCII art with the specified alignment and blocks,
// re-rendering whenever the terminal is resized (SIGWINCH).
// It prints directly rather than returning a string so the resize loop can
// keep running for as long as the program is alive.
func AlignArt(position, sentence, style string) {
	// Helper that clears the screen and prints a fresh render
	draw := func() {
		output := buildAlignedOutput(position, sentence, style)
		fmt.Print("\033[H\033[2J\033[3J" + output)
	}

	// Initial draw on startup
	draw()

	// Register a channel to receive SIGWINCH (terminal resize events)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGWINCH)

	// Block and redraw every time the terminal is resized.
	// The loop exits naturally when the user presses Ctrl+C (SIGINT),
	// which Go handles by terminating the program since we don't capture it.
	for range sigCh {
		draw()
	}
}
