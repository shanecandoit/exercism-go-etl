package etl

import "fmt"

// Transform step of an Extract-Transform-Load.
func Transform(input map[int][]string) (out map[string]int) {
	// Transform(map[int][]string(tt.input)

	// The old system stored a list of letters per score:
	// 1 : ["A", "E"], 2:["D","G"], ...

	// instead stores the score per letter,
	// "a":1, "b":3

	for _, score := range input {
		fmt.Println("score", score)
		break
	}

	return out
}
