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


function highestPalindrome($a = '3993',$k = 1) {  
    // return "3993";
    return findHighestPalindrome($a, $k, 0, strlen($a) - 1);
}

// Fungsi pembantu untuk memeriksa apakah sebuah string adalah palindrom
function isPalindrome($str) {
    return $str === strrev($str);
}

// Fungsi rekursif untuk menemukan palindrom tertinggi
function findHighestPalindrome($s, $k, $start, $end) {
    if ($k < 0 || $start >= $end) {
        return isPalindrome($s) ? $s : -1;
    }

    // Ubah string menjadi array untuk modifikasi
    $sArray = str_split($s);

    // Periksa apakah karakter saat ini cocok
    if ($sArray[$start] === $sArray[$end]) {
        return findHighestPalindrome($s, $k, $start + 1, $end - 1);
    }

    // Coba mengganti karakter untuk membuat palindrom
    $newK = $k - 1;

    // Ganti karakter di posisi awal dan akhir dengan nilai maksimum (9) untuk memaksimalkan palindrom
    $replace1 = $sArray;
    $replace2 = $sArray;
    $replace1[$start] = $replace1[$end] = '9';
    $replace2[$start] = $replace2[$end] = max($sArray[$start], $sArray[$end]);

    // Ubah array kembali menjadi string
    $replace1Str = implode('', $replace1);
    $replace2Str = implode('', $replace2);

    // Periksa kedua kemungkinan penggantian
    $result1 = findHighestPalindrome($replace1Str, $newK, $start + 1, $end - 1);
    $result2 = findHighestPalindrome($replace2Str, $newK, $start + 1, $end - 1);

    // Bandingkan hasil untuk mengembalikan nilai tertinggi
    if ($result1 == -1 && $result2 == -1) {
        return -1;
    } elseif ($result1 == -1) {
        return $result2;
    } elseif ($result2 == -1) {
        return $result1;
    } else {
        return max($result1, $result2);
    }
}

?>