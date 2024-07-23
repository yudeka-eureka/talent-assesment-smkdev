package main

import (
	"fmt"
)

func balanceBracket(a string) string {
	if len(a) > 0 && a[0] != '{' {
		a = a[1:]
	}
	stack := []rune{}
	pairs := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	for _, v := range a {
		if v == ' ' {
			continue
		}
		if closing, ok := pairs[v]; ok {
			stack = append(stack, closing)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != v {
				return "NO"
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return "YES"
	}
	fmt.Println(a)
	return "NO"
}

func highestPalindrome(a string, k int) string {
	if a[0] == '0' || a[0] == '1' {
		a = "9" + a[1:]
	}
	if len(a) == 0 || k < 0 {
		return "-1"
	}
	str := []byte(a)
	result, remaining := build(str, 0, len(str)-1, k)
	result, remaining = maximize([]byte(result), 0, len(result)-1, remaining)
	if remaining < 0 {
		fmt.Println(a, k)
		return "-1"
	}
	fmt.Println(a, k)
	return result
}

func build(s []byte, a, b, k int) (string, int) {
	if a >= b || k < 0 {
		return string(s), k
	}
	if s[a] == s[b] {
		return build(s, a+1, b-1, k)
	}
	n := 1
	if k >= n {
		mc := max(s[a], s[b])
		s[a] = mc
		s[b] = mc
		k -= n
		return build(s, a+1, b-1, k)
	}
	return "-1", -1
}

func maximize(s []byte, a, b, k int) (string, int) {
	if a >= b || k < 0 {
		return string(s), k
	}
	if s[a] != '9' {
		if s[a] == s[b] {
			if k >= 2 {
				s[a] = '9'
				s[b] = '9'
				k -= 2
			}
		} else {
			if k >= 1 {
				s[a] = '9'
				s[b] = '9'
				k -= 1
			}
		}
	}
	return maximize(s, a+1, b-1, k)
}

func max(a, b byte) byte {
	if a > b {
		return a
	}
	return b
}

func weightedStrings(a string, b []int) []string {
	weights := make(map[int]bool)
	curr := '0'
	sub := 0
	for _, char := range a {
		if curr != char {
			curr = char
			sub = 0
		}
		sub += int(char - 'a' + 1)
		weights[sub] = true
	}
	result := make([]string, len(b))
	for i, v := range b {
		if weights[v] {
			result[i] = "YES"
		} else {
			result[i] = "NO"
		}
	}
	return result
}
