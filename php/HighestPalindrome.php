<?php

/**
 * Kamu memiliki string yang merepresentasikan angka 3943 
 * lalu diberikan sebuah variabel k untuk melakukan 
 * replacement karakter sejumlah k pada string
 * untuk mendapatkan bentuk palindrom dengan nilai tertinggi. 
 */

/* Aturan:
 * 1. Jika dari sebuah string tidak ditemukan bentuk palindrome-nya 
 * meski sudah melakukan replacement dan 
 * tidak merepresentasikan sebuah angka maka akan mengeluarkan output -1.
 * 2. Tidak boleh menggunakan fungsi bawaan/tools untuk pencarian/filter/sort.
 * 3. Tidak boleh menggunakan looping.
 * 4. Hanya diperkenankan menggunakan rekursif.
 * 

Soal:
Buat fungsi yang digunakan untuk menyelesaikan permasalahan Highest Palindrome! */


function highestPalindrome($s = '3993', $k = 1) {
    // Helper function to convert string to array and vice versa
    $n = strlen($s);
    $arr = str_split($s);
    
    // Recursive function to attempt creating the highest palindrome
    function makePalindrome(&$arr, $left, $right, &$k, $n) {
        // Base case: when pointers cross, return true (it's a palindrome)
        if ($left >= $right) {
            return true;
        }
        
        // Check if characters are equal
        if ($arr[$left] !== $arr[$right]) {
            // Reduce k for each difference and change the smaller digit to the larger one
            if ($k <= 0) return false; // If no changes left, return false
            $k--;
            $arr[$left] = $arr[$right] = max($arr[$left], $arr[$right]);
        }
        
        // Recur for the next pair
        return makePalindrome($arr, $left + 1, $right - 1, $k, $n);
    }

    // Attempt to make the string a palindrome
    if (!makePalindrome($arr, 0, $n - 1, $k, $n)) {
        return '-1'; // If we can't make it a palindrome, return -1
    }

    // Second pass: make it the highest possible palindrome
    function maximizePalindrome(&$arr, $left, $right, &$k, $n) {
        if ($left >= $right || $k <= 0) return;
        
        // If both digits are not '9', try to maximize them to '9'
        if ($arr[$left] !== '9') {
            // Check if we can make the pair 9
            if ($arr[$left] !== $arr[$right]) {
                if ($k >= 1) {
                    $arr[$left] = $arr[$right] = '9';
                    $k--;
                }
            } elseif ($k >= 2) {
                $arr[$left] = $arr[$right] = '9';
                $k -= 2;
            }
        }

        // Recur for the next pair
        maximizePalindrome($arr, $left + 1, $right - 1, $k, $n);
    }
    
    // Maximize the palindrome
    maximizePalindrome($arr, 0, $n - 1, $k, $n);
    
    // Return the highest possible palindrome as a string
    return implode('', $arr);
}

?>