package main

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

// Helper function to get the maximum of two byte values
func max(a, b byte) byte {
	if a > b {
		return a
	}
	return b
}

// Function to determine highest palindrome
func highestPalindrome(a string, k int) string {
	// Parameter definition
	// a = given string, representing a number
	// k = digits

	n := len(a)
	chars := []byte(a)

	// Helper recursive function to build palindrome
	var buildPalindrome func(start, end, changesLeft int) bool
	buildPalindrome = func(start, end, changesLeft int) bool {
		if start >= end { // Base case: reached the middle
			return changesLeft >= 0
		}

		// Recursively process the outer digits
		if chars[start] == chars[end] {
			return buildPalindrome(start+1, end-1, changesLeft)
		}

		// Case where digits at start and end are different
		if changesLeft <= 0 {
			return false
		}

		// Try to make both ends equal to the maximum of the two digits
		maxDigit := max(chars[start], chars[end])
		chars[start], chars[end] = maxDigit, maxDigit

		// Recurse to the next pair after making one change
		return buildPalindrome(start+1, end-1, changesLeft-1)
	}

	// Second helper to upgrade palindrome to maximum possible
	var upgradeToMaxPalindrome func(start, end, changesLeft int) bool
	upgradeToMaxPalindrome = func(start, end, changesLeft int) bool {
		if start >= end || changesLeft <= 0 {
			return true
		}

		if chars[start] == chars[end] && chars[start] != '9' && changesLeft >= 2 {
			// Upgrade both ends to '9'
			chars[start], chars[end] = '9', '9'
			return upgradeToMaxPalindrome(start+1, end-1, changesLeft-2)
		}

		return upgradeToMaxPalindrome(start+1, end-1, changesLeft)
	}

	// First step: Make the string a palindrome with at most k changes
	if !buildPalindrome(0, n-1, k) {
		return "-1" // Impossible to make a palindrome
	}

	// Second step: Upgrade the palindrome to the highest possible value
	upgradeToMaxPalindrome(0, n-1, k)

	return string(chars)
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
