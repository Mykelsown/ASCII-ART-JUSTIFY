package asciiart

import (
	"testing"
)

func TestFormatPrinter(t *testing.T) {
	// A table driven testting that stores the value of the rdge cases name, the desires input and what is expected of our function to output.
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:  "Simple hello",
			input: "hello",
			expected: ` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               `,
		},
		{
			name:  "Special characters",
			input: `[\]^_ 'a`,
			expected: ` ___  __       ___   /\                  _          
|  _| \ \     |_  | |/\|                ( )         
| |    \ \      | |                     |/    __ _  
| |     \ \     | |                          / _` + "`" + ` | 
| |      \ \    | |                         | (_| | 
| |_      \_\  _| |                          \__,_| 
|___|         |___|       ______                    
                         |______|                   `,
		},
		{
			name:  "With newline",
			input: `hello\n`,
			expected: ` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
`,
		},
		{
			name:  "Double newline",
			input: `hello\n\n`,
			expected: ` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               

`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatPrinter(tt.input)
			
			if got != tt.expected {
				// Show detailed comparison
				t.Errorf("\nTest: %s\nInput: %q\n\nExpected (%d chars, %d lines):\n%q\n\nGot (%d chars, %d lines):\n%q",
					tt.name,
					tt.input,
					len(tt.expected),
					countLines(tt.expected),
					tt.expected,
					len(got),
					countLines(got),
					got,
				)
			}
		})
	}
}

// Helper function to count lines
// Loops through the string it receives, when a character in the string it recieves is == \n, then it should increament count by 1.
func countLines(s string) int {
	if s == "" {
		return 0
	}
	count := 1
	for _, c := range s {
		if c == '\n' {
			count++
		}
	}
	return count
}