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

function isBalanced($input) {
    // Stack to keep track of opening brackets
    $stack = [];
    
    // Map of opening to closing brackets
    $brackets = [
        ')' => '(',
        ']' => '[',
        '}' => '{'
    ];

    // Traverse through the input string
    for ($i = 0; $i < strlen($input); $i++) {
        $char = $input[$i];

        // Ignore any whitespace
        if (ctype_space($char)) {
            continue;
        }

        // If the character is an opening bracket, push it onto the stack
        if (in_array($char, ['(', '[', '{'])) {
            array_push($stack, $char);
        }
        // If the character is a closing bracket, check if it matches the last opening bracket
        elseif (isset($brackets[$char])) {
            // If stack is empty or doesn't match the last opening bracket, return "NO"
            if (empty($stack) || array_pop($stack) !== $brackets[$char]) {
                return "NO";
            }
        }
    }

    // If stack is empty, all opening brackets had matching closing brackets, return "YES"
    return empty($stack) ? "YES" : "NO";
}

?>