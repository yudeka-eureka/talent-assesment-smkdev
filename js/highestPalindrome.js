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

function highestPalindrome(a = '3993',k = 1) {
    // code disini
    if (!/^\d+$/.test(a)) return '-1'; 

    const n = a.length;
    let s = a.split('');
    let changes = Array(n).fill(false); 

    function makePalindrome(l, r, k) {
        if (l > r) return k;

        if (s[l] !== s[r]) {
            const maxChar = Math.max(s[l], s[r]);
            s[l] = s[r] = maxChar;
            k--;
            changes[l] = changes[r] = true;
        }

        if (k < 0) return -1;

        return makePalindrome(l + 1, r - 1, k);
    }

    function maximizePalindrome(l, r, k) {
        if (l > r || k <= 0) return;

        if (s[l] !== '9') {
            if (changes[l] || changes[r]) {
                if (k >= 1) {
                    s[l] = s[r] = '9';
                    k--;
                }
            } else {
                if (k >= 2 && l !== r) {
                    s[l] = s[r] = '9';
                    k -= 2;
                } else if (k >= 1 && l === r) {
                    s[l] = '9';
                    k--;
                }
            }
        }

        maximizePalindrome(l + 1, r - 1, k);
    }

    k = makePalindrome(0, n - 1, k);
    if (k === -1) return '-1';

    maximizePalindrome(0, n - 1, k);

    return s.join('');
}

module.exports = highestPalindrome;