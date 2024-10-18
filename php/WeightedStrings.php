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

 
function weightedStrings($string = "abbcccd", $queries = [1, 3, 9, 8]) {
    $weights = []; // Array to store all possible weights

    // Get the length of the input string
    $n = strlen($string);

    // Calculate the weights for each character and consecutive substrings
    $i = 0;
    while ($i < $n) {
        $currentChar = $string[$i];
        $weight = ord($currentChar) - ord('a') + 1; // Calculate weight of the character
        $sum = 0;
        $count = 0;

        // Check for consecutive occurrences of the same character
        while ($i < $n && $string[$i] === $currentChar) {
            $count++;
            $sum += $weight;
            $weights[$sum] = true; // Add the weight of each consecutive substring
            $i++;
        }
    }

    // Process the queries and check if their values match any weight
    $result = [];
    foreach ($queries as $query) {
        if (isset($weights[$query])) {
            $result[] = "Yes";
        } else {
            $result[] = "No";
        }
    }

    return $result;
}

?>