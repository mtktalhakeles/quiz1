package main

import (
	"fmt"     // Used for formatting output and printing to the console.
	"sort"    // Used for sorting operations.
	"strings" // Used for string manipulation operations.
)

func sortByA(words []string) []string {
	// Initiates the sorting process using the sort.Slice function.
	sort.Slice(words, func(i, j int) bool {
		// Counts the occurrences of 'a' characters for both words.
		countA_i := strings.Count(words[i], "a")
		countA_j := strings.Count(words[j], "a")

		if countA_i == countA_j {
			// If the counts of 'a' characters are equal, we sort based on the length of the word.
			return len(words[i]) > len(words[j])
		}

		// Sorts in descending order based on the counts of 'a' characters.
		return countA_i > countA_j
	})

	return words
}

func main() {

	input := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}

	output := sortByA(input)

	fmt.Println(output)
}
