package morse

import (
	"strings"
)

func EncodeToMorse(text string) string {
	var tones []string
	upperText := strings.ToUpper(strings.TrimSpace(text))
	runes := []rune(upperText)

	for i, char := range runes {
		if char == ' ' {
			// Inter-word gap: 7 dot units (already 3 from previous char, so add 4 more)
			for k := 0; k < 4; k++ {
				tones = append(tones, SilenceDot)
			}
			continue
		}
		code, exists := morseCodeMap[char]
		if !exists {
			continue
		}
		for j, symbol := range code {
			if symbol == '.' {
				tones = append(tones, Dot)
			} else if symbol == '-' {
				tones = append(tones, Dash)
			}
			// Intra-character gap (1 dot) after each symbol except the last
			if j < len(code)-1 {
				tones = append(tones, SilenceDot)
			}
		}
		// Inter-character gap (3 dot units, but 1 already after last symbol)
		// Only if not last character and next is not space
		if i < len(runes)-1 && runes[i+1] != ' ' {
			tones = append(tones, SilenceDot, SilenceDot)
		}
	}
	return strings.Join(tones, ",")
}
