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
    function makePalindrome(s, k) {
        s = s.split('');
        let changes = new Array(s.length).fill(false);

        for (let i = 0; i < Math.floor(s.length / 2); i++) {
            const left = s[i];
            const right = s[s.length - 1 - i];

            if (left !== right) {
                const maxChar = left > right ? left : right;
                s[i] = maxChar;
                s[s.length - 1 - i] = maxChar;
                changes[i] = true;
                k--;
            }
        }

        if (k < 0) return null;

        for (let i = 0; i < Math.floor(s.length / 2); i++) {
            if (s[i] !== '9') {
                if (changes[i] && k >= 1) {
                    s[i] = '9';
                    s[s.length - 1 - i] = '9';
                    k -= 1;
                } else if (!changes[i] && k >= 2) {
                    s[i] = '9';
                    s[s.length - 1 - i] = '9';
                    k -= 2;
                }
            }
        }

        return [s.join(''), k];
    }

    const result = makePalindrome(a, k);

    if (result === null) {
        return '-1';
    }

    return result[0];
}

module.exports = highestPalindrome;