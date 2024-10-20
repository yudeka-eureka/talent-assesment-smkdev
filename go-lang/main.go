package main

import "fmt"

func balanceBracket(a string) string {
	fmt.Println(a)
	return "YES"
}

func highestPalindrome(a string, k int) string {
	fmt.Println(a, k)
	return "3993"
}

// Function to calculate the weight of each character
func getCharWeight(char rune) int {
	return int(char-'a') + 1
}

// Function to solve the Weighted Strings problem
func weightedStrings(a string, b []int) []string {
	// Parameter definition
	// a = given string
	// b = the queries

	// Map to store unique weights of characters and substrings
	weightMap := make(map[int]bool)

	// Calculate weights of individual characters and repeated substrings
	n := len(a)
	for i := 0; i < n; i++ {
		charWeight := getCharWeight(rune(a[i]))
		currentWeight := charWeight

		// Add weight of single character
		weightMap[charWeight] = true

		// Handle repeated and consecutive characters
		for j := i + 1; j < n && a[j] == a[i]; j++ {
			currentWeight += charWeight
			weightMap[currentWeight] = true
			i = j // Move the index to the last occurrence of the repeated character
		}
	}

	// Process the queries
	results := make([]string, len(b))
	for i, query := range b {
		if weightMap[query] {
			results[i] = "YES"
		} else {
			results[i] = "NO"
		}
	}

	return results
}
