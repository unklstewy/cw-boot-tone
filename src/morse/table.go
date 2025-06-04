package morse

var (
	Dot         = "38,2" // dot: tone 38 for 2 units
	Dash        = "38,6" // dash: tone 38 for 6 units
	SilenceDot  = "0,2"  // silence for 2 units (dot space)
	SilenceDash = "0,6"  // silence for 6 units (dash/word space)
)

var morseCodeMap = map[rune]string{
	'A': ".-",
	'B': "-...",
	'C': "-.-.",
	'D': "-..",
	'E': ".",
	'F': "..-.",
	'G': "--.",
	'H': "....",
	'I': "..",
	'J': ".---",
	'K': "-.-",
	'L': ".-..",
	'M': "--",
	'N': "-.",
	'O': "---",
	'P': ".--.",
	'Q': "--.-",
	'R': ".-.",
	'S': "...",
	'T': "-",
	'U': "..-",
	'V': "...-",
	'W': ".--",
	'X': "-..-",
	'Y': "-.--",
	'Z': "--..",
	'1': ".----",
	'2': "..---",
	'3': "...--",
	'4': "....-",
	'5': ".....",
	'6': "-....",
	'7': "--...",
	'8': "---..",
	'9': "----.",
	'0': "-----",
	' ': " ", // Space between words (handled specially in encoder)
}
