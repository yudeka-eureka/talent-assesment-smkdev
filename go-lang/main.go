package main

import (
	"fmt"
	"unicode"
)

func balanceBracket(s string) string {
	stack := []rune{}
	bracetMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			// Menambahkan tanda kurung buka ke stack
			stack = append(stack, char)
		case ')', '}', ']':
			// Memeriksa tanda kurung penutup
			if len(stack) == 0 || stack[len(stack)-1] != bracetMap[char] {
				return "NO"
			}
			// Menghapus tanda kurung buka dari stack
			stack = stack[:len(stack)-1]
		default:
			// Abaikan karakter selain tanda kurung dan karakter spasi, tab, atau newline
			if char != ' ' && char != '\t' && char != '\n' {
				return "NO"
			}
		}
	}

	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}

// func highestPalindrome(a string, k int) string {
// 	fmt.Println(a, k)
// 	return "3993"
// }

// Fungsi utama untuk mencari palindrom terbesar
func highestPalindrome(s string, k int) string {
	isValidNumber := func(s string) bool {
		for _, char := range s {
			if !unicode.IsDigit(char) {
				return false
			}
		}
		return true
	}

	if !isValidNumber(s) {
		return "-1"
	}

	n := len(s)
	if k < 0 {
		return "-1"
	}

	runes := []rune(s)
	left, right := 0, n-1
	changes := 0

	// Pertama: Mengubah string menjadi palindrom dengan perubahan minimal
	for left < right {
		if runes[left] != runes[right] {
			if runes[left] > runes[right] {
				runes[right] = runes[left]
			} else {
				runes[left] = runes[right]
			}
			changes++
		}
		left++
		right--
	}

	if changes > k {
		return "-1"
	}

	// Kedua: Mengoptimalkan palindrom dengan sisa perubahan
	left, right = 0, n-1
	for left < right {
		if runes[left] != '9' {
			if runes[left] != runes[right] {
				if k-changes >= 1 {
					runes[left] = '9'
					runes[right] = '9'
					k -= 1
				}
			} else {
				if k-changes >= 2 {
					runes[left] = '9'
					runes[right] = '9'
					k -= 2
				}
			}
		}
		left++
		right--
	}

	// Untuk string dengan panjang ganjil, ubah karakter tengah menjadi '9' jika masih ada perubahan
	if n%2 == 1 && k > 0 {
		runes[n/2] = '9'
	}

	return string(runes)
}

func main() {
	// Contoh penggunaan fungsi
	fmt.Println(highestPalindrome("3943", 1))    // Output: 3993
	fmt.Println(highestPalindrome("092282", 3))  // Output: 992299
	fmt.Println(highestPalindrome("5566", 1))    // Output: -1
	fmt.Println(highestPalindrome("932239", 2))  // Output: 992299
	fmt.Println(highestPalindrome("11331", 4))   // Output: 99399
	fmt.Println(highestPalindrome("A1341", 1))   // Output: -1
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
