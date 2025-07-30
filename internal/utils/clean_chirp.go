package utils

import (
	"fmt"
	"strings"
)

func CleanChirp(chirp string, toReplace []string) string {
	const stars = "****"
	for _, word := range toReplace {
		lowerChirp := strings.ToLower(chirp)
		replaceIdxs := findIndices(lowerChirp, word)
		delta := len(word) - len(stars)
		for i, idx := range replaceIdxs {
			start := idx - i*delta
			end := idx - i*delta + len(word)
			chirp = chirp[:start] + stars + chirp[end:]
			fmt.Println(chirp)
		}
	}

	return chirp
}

func findIndices(str, substr string) []int {
	indices := []int{}
	index := strings.Index(str, substr)
	for index != -1 {
		indices = append(indices, index)
		index = strings.Index(str[index+1:], substr)
		if index != -1 {
			index += indices[len(indices)-1] + 1
		}
	}
	return indices
}
