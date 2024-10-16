package main

import (
	"unicode"
)

func balanceBracket(a string) string {
	if len(a) == 0 || len(a)%2 == 1 {
		return "NO"
	}

	tbl := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []rune{}

	for _, b := range a {
		if _, ok := tbl[b]; ok {
			stack = append(stack, b)
		} else if len(stack) == 0 || tbl[stack[len(stack)-1]] != b {
			return "NO"
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}

func highestPalindrome(a string, k int) string {
	n := len(a)
	arr := []rune(a)
	left, right := 0, n-1
	changes := make([]bool, n)

	for left < right {
		if arr[left] != arr[right] {
			if !unicode.IsDigit(arr[left]) || !unicode.IsDigit(arr[right]) {
				return "-1"
			}

			if arr[left] > arr[right] {
				arr[right] = arr[left]
			} else {
				arr[left] = arr[right]
			}
			changes[left] = true
			k--
		}
		left++
		right--
	}

	if k < 0 {
		return "-1"
	}

	left, right = 0, n-1
	for left <= right {
		if left == right {
			if k > 0 && arr[left] != '9' {
				arr[left] = '9'
				k--
			}
		} else {
			if arr[left] != '9' {
				if changes[left] && k > 0 {
					arr[left] = '9'
					arr[right] = '9'
					k--
				} else if !changes[left] && k >= 2 {
					arr[left] = '9'
					arr[right] = '9'
					k -= 2
				}
			}
		}
		left++
		right--
	}

	return string(arr)
}

func weightedStrings(a string, b []int) []string {
	weights := make(map[int]struct{})
	result := make([]string, len(b))

	for i := 0; i < len(a); {
		char := a[i]
		charWeight := int(char - 'a' + 1)
		count := 0

		for i < len(a) && a[i] == char {
			count++
			weights[charWeight*count] = struct{}{}
			i++
		}
	}

	for i, q := range b {
		if _, exists := weights[q]; exists {
			result[i] = "YES"
		} else {
			result[i] = "NO"
		}
	}

	return result
}
