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


function highestPalindrome($a, $k) {
    // Step 1: Check if $a is a string, if not return -1
    if (!is_string($a)) {
        return -1;
    }
    
    // Helper function to check if a string is a palindrome
    $isPalindrome = function ($str, $start, $end) use (&$isPalindrome) {
        if ($start >= $end) {
            return true;
        }
        if ($str[$start] != $str[$end]) {
            return false;
        }
        return $isPalindrome($str, $start + 1, $end - 1);
    };

    // Helper function to find the highest palindrome by recursive replacement
    $findHighestPalindrome = function($str, $k, $start, $end) use (&$findHighestPalindrome) {
        if ($k < 0) {
            return null;
        }
        if ($start >= $end) {
            return $str;
        }

        if ($str[$start] == $str[$end]) {
            return $findHighestPalindrome($str, $k, $start + 1, $end - 1);
        } else {
            $originalStr = $str;
            // Replace with the higher of the two characters
            $str[$start] = $str[$end] = max($str[$start], $str[$end]);
            $option1 = $findHighestPalindrome($str, $k - 1, $start + 1, $end - 1);

            // Consider not replacing and just moving on
            $option2 = $findHighestPalindrome($originalStr, $k, $start + 1, $end - 1);

            if (is_null($option1)) return $option2;
            if (is_null($option2)) return $option1;

            return strcmp($option1, $option2) > 0 ? $option1 : $option2;
        }
    };

    // Step 2: Check if $a is already a palindrome
    if ($isPalindrome($a, 0, strlen($a) - 1)) {
        return $a;
    }

    // Step 3: Find the highest palindrome possible with up to $k replacements
    $result = $findHighestPalindrome($a, $k, 0, strlen($a) - 1);

    // Step 4: Return the highest palindrome number or -1 if not possible
    return $result !== null ? $result : -1;
}
?>