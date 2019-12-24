package etl

import (
	"strings"
)

// Transform step of an Extract-Transform-Load.
func Transform(input map[int][]string) (out map[string]int) {
	// Transform(map[int][]string(tt.input)

	// The old system stored a list of letters per score:
	// 1 : ["A", "E"], 2:["D","G"], ...

	// instead store the score per letter,
	// "a":1, "b":3

	out = make(map[string]int, len(input)) // slightly faster
	//out = make(map[string]int)

	for score, letters := range input {
		// fmt.Println("score", score)
		// fmt.Println("letters", letters)
		for _, letter := range letters {
			// fmt.Println("letter", letter)
			letter = strings.ToLower(letter)
			out[letter] = score
		}
	}

	return out
}
