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

function balanceBracket($str) {
    $brackets = array(
        ')' => '(',
        '}' => '{',
        ']' => '['
    );
    $stack = [];
    
    for ($i = 0; $i < strlen($str); $i++) {
        $char = $str[$i];
        // Skip non-bracket characters (e.g., spaces)
        if (!in_array($char, array_merge(array_keys($brackets), array_values($brackets)))) {
            continue;
        }
        // If it's an opening bracket, push it onto the stack
        if (in_array($char, $brackets)) {
            array_push($stack, $char);
        } else {
            // If it's a closing bracket, check if it matches the top of the stack or the stack is empty
            if (empty($stack) || array_pop($stack) !== $brackets[$char]) {
                return "NO";
            }
        }
    }

    // If the stack is empty after iterating through the string, it means all brackets are balanced
    return count($stack) == 0 ? "YES" : "NO";
}

?>