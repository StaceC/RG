package code

import (
	"fmt"
	"strings"
)

var code = map[int32]string{
	'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".", 'F': "..-.", 'G': "--.", 'H': "....", 'I': "..", 'J': ".---", 'K': "-.-", 'L': ".-..", 'M': "--", 'N': "-.", 'O': "---", 'P': ".--.", 'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-", 'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-", 'Y': "-.--", 'Z': "--..", '0': "-----", '1': ".----", '2': "..---", '3': "...--", '4': "....-", '5': ".....", '6': "-....", '7': "--...", '8': "---..", '9': "----.", '.': ".-.-.-", ',': "--..--"}

func Encode(input string, obfuscate bool) (string, error) {

	input = strings.ToUpper(input)

	output := ""
	var last int32
	for _, char := range input {

		output += addSeparator(last, char)
		if obfuscate {
			obfuscated, _ := Obfuscate(code[char])
			output += obfuscated
		} else {
			output += code[char]
		}

		//fmt.Printf("last[%c] char[%c] output[%s]\n", last, char, output)
		last = char
	}

	return output, nil
}

func addSeparator(last, current int32) string {
	if last == 0 {
		return ""
	} else if last != ' ' && current == ' ' {
		return "/"
	} else if last == ' ' {
		return ""
	} else {
		return "|"
	}
}

func Obfuscate(input string) (string, error) {

	if len(input) == 0 {
		return "", nil
	}

	// Setup lookahead
	output := ""
	current := input[0]
	obfs, err := resetRune(input[0])
	if err != nil {
		return "", err
	}
	// Loop through input
	for i := 1; i < len(input); i++ {

		// fmt.Printf("Current[%c] Next[%c] Obfs[%c] Output[%s]\n", current, input[i], obfs, output)

		if current == input[i] { // Same rune as before?
			// Up the byte
			obfs++
		} else { // new rune ahead?
			// Add rune to output
			output += string(obfs)
			// Reset obfuscation rune
			obfs, err = resetRune(input[i])
			if err != nil {
				return "", err
			}
		}
		current = input[i]

	}
	// Catch last output
	output += string(obfs)

	return output, nil
}

func resetRune(c byte) (byte, error) {
	switch c {
	case '.':
		return '1', nil
	case '-':
		return 'A', nil
	}
	return 0, fmt.Errorf("Invalid rune %c", c)
}
