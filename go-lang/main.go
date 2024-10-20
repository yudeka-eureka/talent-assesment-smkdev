package main

func balanceBracket(a string) string {
	stack := make([]rune, 0, len(a))
	
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range a {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != bracketMap[char] {
				return "NO"
			}
			stack = stack[:len(stack)-1]

		default:
			return "NO"
		}
	}
	
	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}


func highestPalindrome(a string, k int) string {
	chars := []rune(a)
	n := len(chars)
	changes := make([]bool, n) 
	
	var makePalindrome func(left, right int) bool
	makePalindrome = func(left, right int) bool {
		if left >= right {
			return true 
		}

		if chars[left] != chars[right] {
			if chars[left] > chars[right] {
				chars[right] = chars[left]
			} else {
				chars[left] = chars[right]
			}
			changes[left] = true
			changes[right] = true
			k-- 
		}

		if k < 0 {
			return false 
		}

		return makePalindrome(left+1, right-1)
	}

	if !makePalindrome(0, n-1) || k < 0 {
		return "-1" 
	}

	
	var maximizePalindrome func(left, right int) string
	maximizePalindrome = func(left, right int) string {
		if left > right {
			return string(chars) 
		}
		
		if left == right {
			if k > 0 {
				chars[left] = '9'
			}
			return string(chars)
		}
		
		if chars[left] != '9' {
			if changes[left] && k > 0 {
				chars[left], chars[right] = '9', '9'
				k--
			} else if !changes[left] && k >= 2 {
				chars[left], chars[right] = '9', '9'
				k -= 2
			}
		}

		return maximizePalindrome(left+1, right-1) 
	}

	return maximizePalindrome(0, n-1)
}

func weightedStrings(a string, b []int) []string {
	weights := map[int]bool{}         
	result := make([]string, len(b)) 

	var currentChar rune
	var currentWeight, charWeight, repeatCount int

	for _, char := range a {
		charWeight = int(char - 'a' + 1)

		if char != currentChar {
			currentChar = char
			repeatCount = 1
			currentWeight = charWeight
		} else {
			repeatCount++
			currentWeight += charWeight
		}
		weights[currentWeight] = true
	}
	
	for i, query := range b {
		if weights[query] {
			result[i] = "YES"
		} else {
			result[i] = "NO"
		}
	}

	return result
}
