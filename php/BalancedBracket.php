<?php
/** 
 * Aturan:
 * 1. Tanda braket / karakter yang diperbolehkan sebagai berikut: ( , ) , { , } , atau [ , ]. 
 * 2. Bracket bisa dipisahkan dengan atau tanpa whitespace.
 * 3. Periksa tanda kurung yang memiliki kecocokan antara braket buka dan braket tutup dengan mengembalikan nilai string YES atau NO.
 * 
 * Soal:
 * 1. Buat fungsi untuk menemukan Balanced Bracket dengan kompleksitas paling rendah!
 * 2. Jelaskan kompleksitas dari penyelesaianmu untuk No.3 dan cantumkan di README Repo! 
*/


function balanceBracket($a)
{
    $stack = [];
    $matchingBrackets = [
        ")" => "(",
        "}" => "{",
        "]" => "[",
    ];

    for ($i = 0; $i < strlen($a); $i++) {
        $char = $a[$i];

        if ($char === "(" || $char === "{" || $char === "[") {
            array_push($stack, $char);
        } else if ($char === ")" || $char === "}" || $char === "]") {
            if (count($stack) === 0 || array_pop($stack) !== $matchingBrackets[$char]) {
                return "NO";
            }
        } else if (trim($char) !== "") {
            // Check for any non-bracket and non-whitespace character
            return "NO";
        }
    }

    return count($stack) === 0 ? "YES" : "NO";
}

?>