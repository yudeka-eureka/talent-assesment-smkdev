package main

import (
	"regexp"
	"strings"
)

func weightedStrings(a string, b []int) []string {
	var res []string

	if a == "" || len(b) == 0 {
		return res
	}

	// set letter weight mapping
	letters := "abcdefghijklmnopqrstuvwxyz"
	letterWeight := make(map[string]int)
	for i, v := range letters {
		letter := string(v)
		letterWeight[letter] = i + 1
	}

	weightMaps := make(map[int]bool)
	curr := ""
	counter := 1
	for _, v := range a {
		letter := strings.ToLower(string(v))
		if curr == letter {
			counter++
		} else {
			counter = 1
			curr = letter
		}

		// set weight in map
		weight, isExist := letterWeight[letter]
		if !isExist {
			continue
		}

		value := counter * (weight)

		weightMaps[value] = true
	}

	// compare array b with weightMap
	for _, v := range b {
		if weightMaps[v] {
			res = append(res, "YES")
			continue
		}

		res = append(res, "NO")
	}

	return res
}

func balanceBracket(a string) string {
	var (
		openBrackets    []string
		openBracketsLen int
	)

	mapBrackets := map[string]string{"(": ")", "[": "]", "{": "}"}

	// remove whitespaces
	a = strings.ReplaceAll(a, " ", "")

	// check whether length of string is zero
	if len(a) == 0 {
		return "NO"
	}

	for _, v := range a {
		curr := string(v)
		// if char is an open bracket, add it to a openBrackets and add string length
		_, isExist := mapBrackets[curr]
		if isExist {
			openBrackets = append(openBrackets, curr)
			openBracketsLen++
		} else {
			closeBracket := mapBrackets[openBrackets[openBracketsLen-1]]

			// if char is not a close bracket, then return NO
			if curr != closeBracket {
				return "NO"
			}

			// remove last char of openBrackets and substract length of openBrackets
			openBrackets = openBrackets[:openBracketsLen-1]
			openBracketsLen--
		}
	}

	if openBracketsLen != 0 {
		return "NO"
	}

	return "YES"
}

func highestPalindrome(a string, k int) string {
	if !regexp.MustCompile(`^[0-9]+$`).MatchString(a) {
		return "-1"
	}

	// set string that is not palindrome into palindrome
	a, k, changed := setPalindrome(a, k, 0, make([]bool, len(a)))

	// if there are no more steps and its not palindrome yet return -1
	if k == -1 || !isPalindrome(a, 0) {
		return "-1"
	}

	// if there are still more steps to take, try to maximize the value
	if k > 0 {
		a = maximizePalindrome(a, k, 0, changed)
	}

	return a
}

func isPalindrome(a string, el int) bool {
	lenA := len(a)
	if lenA <= 1 {
		return true
	}

	if el == (lenA / 2) {
		return true
	}

	if a[el] != a[lenA-el-1] {
		return false
	}

	return isPalindrome(a, el+1)
}

func setPalindrome(a string, b, idx int, changed []bool) (string, int, []bool) {
	if a == "" {
		return a, b, changed
	}

	if b == 0 {
		return a, b, changed
	}

	lenA := len(a)

	// when in element that is halfway the string then return string
	if idx >= lenA/2 {
		return a, b, changed
	}

	// if there are no more steps and there are char that is not the same with its counterpart, return string
	if a[idx] != a[lenA-idx-1] && b == 0 {
		return a, b, changed
	}

	// compare an element and its counterpart, change either value to which one that is bigger
	if a[idx] > a[lenA-idx-1] {
		a = a[:lenA-idx-1] + string(a[idx]) + a[lenA-idx:]
		b--
		changed[idx], changed[lenA-idx-1] = true, true
	} else if a[idx] < a[lenA-idx-1] {
		a = a[:idx] + string(a[lenA-idx-1]) + a[idx+1:]
		b--
		changed[idx], changed[lenA-idx-1] = true, true
	}

	return setPalindrome(a, b, idx+1, changed)
}

func maximizePalindrome(a string, b, idx int, changed []bool) string {
	if a == "" || b == 0 {
		return a
	}

	lenA := len(a)

	if idx >= lenA/2 {
		if (lenA%2 == 1) && b > 0 {
			a = a[:idx] + "9" + a[idx+1:]
			b--
		}
		return a
	}

	// try to change char to "9"
	if string(a[idx]) != "9" {
		// if char was already changed and there are steps available, do 1 more change
		if changed[idx] && b > 0 {
			a = a[:idx] + "9" + a[idx+1:lenA-idx-1] + "9" + a[lenA-idx:]
			b--
		} else if !changed[idx] && b > 1 { // if char wasn't changed previously, change both char and its counterpart
			a = a[:idx] + "9" + a[idx+1:lenA-idx-1] + "9" + a[lenA-idx:]
			b -= 2
		}
	}

	return maximizePalindrome(a, b, idx+1, changed)
}
