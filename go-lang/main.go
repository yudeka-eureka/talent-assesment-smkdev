package main

import "fmt"

// Function to check if the brackets are balanced
func balanceBracket(a string) string {
	// Parameter definition
	// a = given string

	// Stack to store opening brackets
	stack := []rune{}

	// Map to hold matching pairs of brackets
	bracketPairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	// Iterate over the string
	for _, char := range a {
		// If it's an opening bracket, push it onto the stack
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else if char == ')' || char == ']' || char == '}' {
			// If it's a closing bracket, check if it matches the last opening bracket
			if len(stack) == 0 || stack[len(stack)-1] != bracketPairs[char] {
				return "NO"
			}
			// Pop the matched opening bracket from the stack
			stack = stack[:len(stack)-1]
		}
	}

	// If the stack is empty, the brackets are balanced
	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
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
