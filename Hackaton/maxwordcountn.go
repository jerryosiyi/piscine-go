package piscine

import "sort"

func MaxWordCountN(text string, n int) map[string]int {
	// Split 'text' into individual words using spaces as delimiters
	words := make([]string, 0)
	start := 0
	for i, r := range text {
		if r == ' ' {
			words = append(words, text[start:i])
			start = i + 1
		}
	}
	words = append(words, text[start:])
	occ := make(map[string]int)
	for _, word := range words {
		occ[word]++
	}
	// Sort the words in descending order of frequency and ascending order of word (for ties)
	sorted := make([]string, 0, len(occ))
	for word := range occ {
		sorted = append(sorted, word)
	}
	sort.Slice(sorted, func(i, j int) bool {
		if occ[sorted[i]] != occ[sorted[j]] {
			return occ[sorted[i]] > occ[sorted[j]]
		}
		return sorted[i] < sorted[j]
	})
	if n > len(sorted) {
		n = len(sorted)
	}
	tied := sorted[:n]

	// Construct a new map containing the top 'n' words and their counts
	result := make(map[string]int)
	for _, word := range tied {
		result[word] = occ[word]
	}
	return result
}
