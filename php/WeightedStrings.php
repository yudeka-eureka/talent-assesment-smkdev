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

 
 function weightedStrings($string, $queries) {
    $weights = [];
    $length = strlen($string);
    
    $i = 0;
    while ($i < $length) {
        $char = $string[$i];
        $charWeight = ord($char) - ord('a') + 1;
        $repeatWeight = $charWeight;
        $count = 1;

        while ($i + 1 < $length && $string[$i] == $string[$i + 1]) {
            $count++;
            $repeatWeight += $charWeight;
            $weights[$repeatWeight] = true;
            $i++;
        }
        
        $weights[$charWeight] = true;
        
        // Move to the next character
        $i++;
    }

    // Check each query against the weights we have collected
    $results = [];
    foreach ($queries as $query) {
        if (isset($weights[$query])) {
            $results[] = "Yes";
        } else {
            $results[] = "No";
        }
    }

    return $results;
}


?>