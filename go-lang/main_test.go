// math_test.go
package main

import "testing"

func TestBalanceBracket(t *testing.T) {
	result := balanceBracket("{ [ ( ] ) }")
	if result != "NO" {
		t.Errorf("balanceBracket('{ [ ( ] ) }') = %s; want NO", result)
	}

	result = balanceBracket("({ ( ( [ ] ) [ ] ) [ ] }")
	if result != "YES" {
		t.Errorf("balanceBracket('{ ( ( [ ] ) [ ] ) [ ] }') = %s; want YES", result)
	}

	result = balanceBracket("{[()]}")
	if result != "YES" {
		t.Errorf("balanceBracket('{[()]}') = %s; want YES", result)
	}

	result = balanceBracket("{(([|])[])[]}")
	if result != "NO" {
		t.Errorf("balanceBracket('{(([|])[])[]}') = %s; want NO", result)
	}
}

func TestHighestPalindrome(t *testing.T) {
	result := highestPalindrome("3943", 1)
	if result != "3993" {
		t.Errorf("highestPalindrome('3943', 1) = %s; want 3993", result)
	}

	result = highestPalindrome("092282", 3)
	if result != "992299" {
		t.Errorf("highestPalindrome('092282', 3) = %s; want 992299", result)
	}

	result = highestPalindrome("5566", 1)
	if result != "-1" {
		t.Errorf("highestPalindrome('5566', 1) = %s; want -1", result)
	}

	result = highestPalindrome("932239", 2)
	if result != "992299" {
		t.Errorf("highestPalindrome('932239', 2) = %s; want 992299", result)
	}

	result = highestPalindrome("11331", 4)
	if result != "99399" {
		t.Errorf("highestPalindrome('11331', 4) = %s; want 99399", result)
	}

	result = highestPalindrome("A1341", 1)
	if result != "-1" {
		t.Errorf("highestPalindrome('A1341', 1) = %s; want -1", result)
	}
}

func TestWeightedStrings(t *testing.T) {
	result := weightedStrings("abbcccd", []int{1, 3, 9, 8})
	expected := []string{"YES", "YES", "YES", "NO"}
	if !sliceEqual(result, expected) {
		t.Errorf("weightedStrings('abbcccd', [1, 3, 9, 8]) = %v; want %v", result, expected)
	}

	result = weightedStrings("aaabbbcccddd", []int{5, 9, 7, 8, 12, 5})
	expected = []string{"NO", "YES", "NO", "YES", "YES", "NO"}
	if !sliceEqual(result, expected) {
		t.Errorf("weightedStrings('aaabbbcccddd', [5, 9, 7, 8, 12, 5]) = %v; want %v", result, expected)
	}

	result = weightedStrings("getitssimple", []int{1, 2, 8, 7, 12, 38, 10})
	expected = []string{"NO", "NO", "NO", "YES", "YES", "YES", "NO"}
	if !sliceEqual(result, expected) {
		t.Errorf("weightedStrings('getitssimple', [1, 2, 8, 7, 12, 38, 10]) = %v; want %v", result, expected)
	}
}

func sliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
