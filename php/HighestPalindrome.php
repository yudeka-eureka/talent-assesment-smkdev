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


function highestPalindrome($a = "3993", $k = 1)
{
    $len = strlen($a);
    $chars = str_split($a);
    $left = 0;
    $right = $len - 1;
    $changes = array_fill(0, $len, false);

    // Step 1: Make the string a palindrome with minimal changes
    while ($left < $right) {
        if ($chars[$left] !== $chars[$right]) {
            $maxChar = max($chars[$left], $chars[$right]);
            $chars[$left] = $chars[$right] = $maxChar;
            $changes[$left] = $changes[$right] = true;
            $k--;
        }
        $left++;
        $right--;
    }

    if ($k < 0) return "-1";

    // Step 2: Maximize the palindrome value with remaining changes
    $left = 0;
    $right = $len - 1;
    while ($left < $right) {
        if ($chars[$left] !== "9") {
            if ($changes[$left] && $k > 0) {
                $chars[$left] = $chars[$right] = "9";
                $k--;
            } elseif (!$changes[$left] && $k >= 2) {
                $chars[$left] = $chars[$right] = "9";
                $k -= 2;
            }
        }
        $left++;
        $right--;
    }

    // Handle the middle character for odd length
    if ($len % 2 === 1 && $k > 0) {
        $chars[floor($len / 2)] = "9";
    }

    return implode("", $chars);
}


?>