package main

import (
	"fmt"
	"strings"
)

func balanceBracket(a string) string {
	a = strings.ReplaceAll(a, " ", "")
	if len(a)%2 != 0 {
		fmt.Println("string: ", a)
		fmt.Println("len: ", len(a))
		return "NO"
	}

	stack := []rune{}
	brackets := map[rune]rune{'(': ')', '{': '}', '[': ']'}

	for _, char := range a {
		// Abaikan karakter selain tanda kurung
		if _, exists := brackets[char]; exists {
			// Jika karakter adalah opening bracket, tambahkan ke stack
			stack = append(stack, char)
		} else if char == ')' || char == '}' || char == ']' {
			// Jika stack kosong atau closing bracket tidak cocok dengan top of stack
			if len(stack) == 0 || brackets[stack[len(stack)-1]] != char {
				return "NO"
			}
			// Jika cocok, pop stack
			stack = stack[:len(stack)-1]
		}
	}

	// Jika stack kosong, semua bracket seimbang
	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}

// Fungsi untuk membuat palindrome tertinggi dengan maksimal k perubahan
func highestPalindrome(a string, k int) string {
	n := len(a)
	arr := []byte(a)
	changed := make([]bool, n)

	// Langkah 1: Coba ubah menjadi palindrome dengan minimal perubahan
	l, r := 0, n-1
	for l < r {
		if arr[l] != arr[r] {
			if k == 0 {
				return "-1" // Jika perubahan tidak cukup
			}
			if arr[l] > arr[r] {
				arr[r] = arr[l]
			} else {
				arr[l] = arr[r]
			}
			k--
			changed[l] = true
		}
		l++
		r--
	}

	// Langkah 2: Coba ubah sebanyak mungkin menjadi '9'
	l, r = 0, n-1
	for l <= r && k > 0 {
		if l == r {
			// Jika di tengah, ubah menjadi '9' jika k masih tersisa
			if arr[l] != '9' && k > 0 {
				arr[l] = '9'
				k--
			}
		} else {
			if arr[l] != '9' {
				if changed[l] && k > 0 {
					arr[l], arr[r] = '9', '9'
					k--
				} else if !changed[l] && k > 1 {
					arr[l], arr[r] = '9', '9'
					k -= 2
				}
			}
		}
		l++
		r--
	}

	return string(arr)
}

// Fungsi untuk mengecek substring dengan weighted sum
func weightedStrings(a string, b []int) []string {
	weights := map[int]bool{}
	currentWeight := 0
	prevChar := ' '

	// Hitung semua kemungkinan berat substring
	for _, char := range a {
		if char == prevChar {
			currentWeight += int(char - 'a' + 1)
		} else {
			currentWeight = int(char - 'a' + 1)
			prevChar = char
		}
		weights[currentWeight] = true
	}

	// Cek apakah nilai dari array b dapat dibentuk
	results := make([]string, len(b))
	for i, weight := range b {
		if weights[weight] {
			results[i] = "YES"
		} else {
			results[i] = "NO"
		}
	}

	return results
}
