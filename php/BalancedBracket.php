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

function balanceBracket($a) {  
    // code disini
    // return "YES";

     // Menghapus whitespace dari string
     $str = str_replace(' ', '', $a);

    // Stack untuk melacak braket yang belum dipasangkan
    $stack = [];

    // Map untuk mencocokkan braket buka dan tutup
    $bracketMap = [
        ')' => '(',
        '}' => '{',
        ']' => '['
    ];

    // Melakukan looping pada setiap karakter dalam string 
    foreach (str_split($str) as $char) {

        if (in_array($char, ['(', '{', '['])) {
            // Jika karakter adalah braket buka, push ke stack
            $stack[] = $char;
        } elseif (in_array($char, [')', '}', ']'])) {
            // Jika karakter adalah braket tutup
            if (empty($stack) || end($stack) !== $bracketMap[$char]) {
                // Jika stack kosong atau braket tutup tidak cocok dengan braket buka di atas stack
                return "NO";
            }
            // Jika cocok, pop braket buka dari stack
            array_pop($stack);
        }
    }

     // Jika stack kosong, semua braket buka memiliki pasangan
     return empty($stack) ? "YES" : "NO";
}

?>