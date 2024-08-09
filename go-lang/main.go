package main

import "unicode"

func balanceBracket(s string) string {
	stack := []rune{}
	bracetMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, varchar := range s {
		switch varchar {
		case '(', '{', '[':
			stack = append(stack, varchar)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != bracetMap[varchar] {
				return "NO"
			}
			stack = stack[:len(stack)-1]
		default:
			if varchar != ' ' && varchar != '\t' && varchar != '\n' {
				return "NO"
			}
		}
	}

	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}

func highestPalindrome(s string, k int) string {
	isValidNumber := func(s string) bool {
		for _, varchar := range s {
			if !unicode.IsDigit(varchar) {
				return false
			}
		}
		return true
	}

	if !isValidNumber(s) {
		return "-1"
	}

	n := len(s)
	runes := []rune(s)
	left, right := 0, n-1
	changes := make([]bool, n)

	for left < right {
		if runes[left] != runes[right] {
			if runes[left] > runes[right] {
				runes[right] = runes[left]
			} else {
				runes[left] = runes[right]
			}
			changes[left], changes[right] = true, true
			k--
		}
		left++
		right--
	}

	if k < 0 {
		return "-1"
	}

	left, right = 0, n-1
	for left <= right && k > 0 {
		if left == right { 
			if runes[left] != '9' && k > 0 {
				runes[left] = '9'
				k--
			}
		} else {
			if runes[left] != '9' {
				if changes[left] || changes[right] {
					runes[left], runes[right] = '9', '9'
					k--
				} else if k >= 2 {
					runes[left], runes[right] = '9', '9'
					k -= 2
				}
			}
		}
		left++
		right--
	}

	return string(runes)
}

func weightedStrings(s string, query []int) []string {
	weightMap := make(map[int]struct{})

	for i := 0; i < len(s); i++ {
		weight := 0
		for j := i; j < len(s); j++ {
			if s[j] != s[i] {
				break
			}
			weight += int(s[j] - 'a' + 1)
			weightMap[weight] = struct{}{}
		}
	}

	results := make([]string, len(query))
	for i, q := range query {
		if _, found := weightMap[q]; found {
			results[i] = "YES"
		} else {
			results[i] = "NO"
		}
	}

	return results
}
