package main

func balanceBracket(a string) string {
	stack := []rune{}

	bracket := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	for _, i := range a {
		if i == '{' || i == '[' || i == '(' {
			stack = append(stack, i)
		} else if len(stack) > 0 && bracket[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return "NO"
		}
	}

	return "YES"
}

func highestPalindrome(a string, k int) string {
	n := len(a)
	if n == 0 {
		return "-1"
	}

	str := []rune(a)
	changes := 0

	if !isValidNumber(str, 0, n-1) {
		return "-1"
	}

	str, k = makePalindrome(str, 0, n-1, k, &changes)

	if k < 0 {
		return "-1"
	}

	str, _ = maximizePalindrome(str, 0, n-1, k, changes)

	return string(str)
}

func weightedStrings(a string, b []int) []string {
	results := make([]string, len(b))
	weights := make(map[int]bool)

	n := len(a)
	if n == 0 {
		for i := range results {
			results[i] = "NO"
		}
		return results
	}

	currentWeight := charWeight(rune(a[0]))
	weights[currentWeight] = true

	for i := 1; i < n; i++ {
		charW := charWeight(rune(a[i]))

		if a[i] == a[i-1] {
			currentWeight += charW
		} else {
			currentWeight = charW
		}
		weights[currentWeight] = true
	}

	for i, query := range b {
		if weights[query] {
			results[i] = "YES"
		} else {
			results[i] = "NO"
		}
	}

	return results
}

func isValidNumber(s []rune, left int, right int) bool {
	if left > right {
		return true
	}

	if s[left] < '0' || s[left] > '9' || s[right] < '0' || s[right] > '9' {
		return false
	}
	return isValidNumber(s, left+1, right-1)
}

func makePalindrome(s []rune, left int, right int, k int, changes *int) ([]rune, int) {
	if left >= right {
		return s, k
	}

	if s[left] != s[right] {
		if s[left] > s[right] {
			s[right] = s[left]
		} else {
			s[left] = s[right]
		}
		*changes++
		k--
	}

	if k < 0 {
		return s, k
	}

	return makePalindrome(s, left+1, right-1, k, changes)
}

func maximizePalindrome(s []rune, left int, right int, k int, changes int) ([]rune, int) {
	if left > right || k <= 0 {
		return s, k
	}

	if s[left] != '9' {
		if left == right {
			if k > 0 {
				// s[left] = '9'
				k--
			}
		} else if s[left] != '9' {
			if changes > 0 && k > 0 {
				s[left], s[right] = '9', '9'
				k--
			} else if k > 1 {
				s[left], s[right] = '9', '9'
				k -= 2
			}
		}
	}

	return maximizePalindrome(s, left+1, right-1, k, changes)
}

func charWeight(ch rune) int {
	return int(ch - 'a' + 1)
}
