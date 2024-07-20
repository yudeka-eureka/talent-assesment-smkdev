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

 
function weightedStrings($keywords, $queries) {  
    $weights = [];
    $n = strlen($keywords);

    for ($i = 0; $i < $n; $i++) {
        $char = $keywords[$i];
        $weight = ord($char) - ord('a') + 1;

        $count = 1;
        while ($i + 1 < $n && $keywords[$i + 1] == $char) {
            $i++;
            $count++;
            for ($j = 1; $j <= $count; $j++) {
                $weights[] = $weight * $j;
            }
        }
        
        if ($count == 1) {
            $weights[] = $weight;
        }
    }

    sort($weights);

    $result = [];
    foreach ($queries as $query) {
        if (in_array($query, $weights)) {
            $result[] = "YES";
        } else {
            $result[] = "NO";
        }
    }

    return $result;
}

?>