<?php

/**
 * Alfabet dari a sampai z memiliki bobot sebesar angka urutannya, 
 * misal: a = 1, b = 2, c = 3, ..., z = 26. 
 * Bobot sebuah string merupakan penjumlahan bobot dari 
 * kesuluruhan karakter. 
 * Contoh: gits = 7 + 9 + 20 + 19 = 55
 */

/**
 * Aturan:
 * Jika terdapat karakter yang berulang dan berurutan perlu dirincikan ke dalam bentuk substring dari perulangan pertama hingga n. 
 * Contoh: bbccc => b, bb, c, cc, ccc. 
 * 
 * Soal: 
 * Buat fungsi untuk menyelesaikan permasalahan Weighted Strings!
 */

 
function weightedString($a = "abbcccd", $b = [1, 3, 9, 8]) {
    // Initialize variables
    $chars = str_split($a);
    $charlen = count($chars);
    $weights= [];
    $result = [];
    $count=1;
    $currChar = $chars[0];
    $weight = function($currChar) { return ord($currChar) - 96; };

    // insert first character weight
    array_push($weights,$weight($currChar));

    for ($i=1; $i < $charlen; $i++) {
        // check if character same as previous character
        if ($chars[$i] === $currChar) {
            $count++;
        }else{
            $count = 1;
            $currChar = $chars[$i];
        }

        for ($j=1; $j <= $count; $j++) {
            array_push($weights,$weight($currChar) * $j);
        }

    }
    foreach ($b as $c => $v) {
        $result[] = in_array($v, $weights) ? "YES" : "NO";
    }

    return $result;
}
// var_dump();
?>