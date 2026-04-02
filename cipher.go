package caesargo

import (
	"strings"
)

func get_alphabet() [26]string {
	return [26]string{
		"a", "b", "c", "d", "e",
		"f", "g", "h", "i", "j",
		"k", "l", "m", "n", "o",
		"p", "q", "r", "s", "t",
		"u", "v", "w", "x", "y",
		"z",
	}
}

// Caesar encode a text string
func Encode(text string, shift int) string {
	var alphabet = get_alphabet()

	chars := strings.Split(text, "")
	new_chars := make([]string, 0)

	for index := range chars {
		// handle whitespace and whatnot
		if chars[index] == " " {
			new_chars = append(new_chars, " ")
			continue
		}

		curr_pos := 0

		for position := range alphabet {
			if alphabet[position] == chars[index] {
				curr_pos = position
				break
			}
		}

		new_chars = append(new_chars, alphabet[curr_pos+shift])
	}

	return strings.Join(new_chars, "")
}

// Decode a text string
func Decode(text string, deshift int) string {
	// hacky implementation
	return Encode(text, -deshift)
}
