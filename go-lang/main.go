package main

func balanceBracket(a string) string {
	var data []rune

	// Map to match closing brackets with their corresponding opening brackets
	bracketPairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	// Traverse each character in the string
	for _, char := range a {
		if char == '(' || char == '[' || char == '{' {
			data = append(data, char)
		}
		if char == ')' || char == ']' || char == '}' {
			if len(data) == 0 || data[len(data)-1] != bracketPairs[char] {
				return "NO"
			}
			data = data[:len(data)-1]
		}
	}

	if len(data) == 0 {
		return "YES"
	}
	return "NO"
}

// Recursive function to form the highest palindrome with at most k changes
func highestPalindromeRecursive(s string, left, right, k int) string {
	if k < 0 {
		return "-1"
	}
	if left >= right {
		return s
	}

	runes := []rune(s)

	if runes[left] != runes[right] {
		if runes[left] > runes[right] {
			runes[right] = runes[left] // Fix by making the right side match the left side
		} else {
			runes[left] = runes[right] // Fix by making the left side match the right side
		}
		k--
	}

	// Recur for the next pair of characters (moving towards the middle)
	result := highestPalindromeRecursive(string(runes), left+1, right-1, k)
	if result == "-1" {
		return "-1"
	}
	return result
}

// Recursive function to maximize the palindrome
func maximizePalindromeRecursive(s string, left, right, k int) string {
	if k < 0 {
		return "-1"
	}
	if left >= right {
		return s
	}

	runes := []rune(s)

	if runes[left] == runes[right] {
		if runes[left] != '9' && k >= 2 {
			runes[left], runes[right] = '9', '9'
			k -= 2
		}
	} else {
		if runes[left] != '9' && runes[right] != '9' && k >= 1 {
			runes[left], runes[right] = '9', '9'
			k--
		} else if runes[left] != '9' {
			runes[left] = '9'
			k--
		} else if runes[right] != '9' {
			runes[right] = '9'
			k--
		}
	}

	// Recur for the next pair of characters (moving towards the middle)
	result := maximizePalindromeRecursive(string(runes), left+1, right-1, k)
	if result == "-1" {
		return "-1"
	}
	return result
}

// Wrapper function to check and form the highest palindrome
func highestPalindrome(s string, k int) string {
	// First, try to form the lowest possible palindrome by changing digits as needed
	lowPalindrome := highestPalindromeRecursive(s, 0, len(s)-1, k)

	if lowPalindrome == "-1" {
		return "-1"
	}

	// Then, maximize the palindrome by further changing the digits to '9' if possible
	highPalindrome := maximizePalindromeRecursive(lowPalindrome, 0, len(s)-1, k)

	return highPalindrome
}

func charWeight(char rune) int {
	return int(char-'a') + 1
}

// Function to solve the weighted strings problem
func weightedStrings(a string, b []int) []string {
	weights := make(map[int]struct{}) // Set to store unique weights
	result := make([]string, len(b))  // Slice to store the results for each query

	// Process the string and calculate weights
	n := len(a)
	for i := 0; i < n; {
		currentChar := a[i]
		weight := charWeight(rune(currentChar))

		count := 1
		for i+count < n && a[i+count] == currentChar {
			count++
		}

		// Add weights for each substring of repeated characters
		for j := 1; j <= count; j++ {
			weights[weight*j] = struct{}{}
		}

		i += count
	}

	// Check each query
	for i, query := range b {
		if _, exists := weights[query]; exists {
			result[i] = "YES"
		} else {
			result[i] = "NO"
		}
	}

	return result
}
