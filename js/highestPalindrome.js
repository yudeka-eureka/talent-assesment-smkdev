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

    // inisiasi variabel
    const arrayA = [];
    let result = '';

    // recursive function untuk bikin array
    function createArrRecursive(n, current = 0) {
        if (current >= n) {
            const half = Math.floor((arrayA.length)/2);
            setPalindromeRecursive(half);
            return;
        };
        arrayA[current] = a[current];
        createArrRecursive(n, current + 1);
    };

    // recursive function untuk cek palindrome
    function setPalindromeRecursive(n, current = 0) {
        if (current >= n) {
            createResultRecursive(arrayA.length);
            return;
        };

        // inisiasi depan dan belakang dari arrayA
        const front = arrayA[current];
        const back = arrayA[arrayA.length - current - 1];

        // proses menentukan palindrome tertinggi
        if (k < 0) return '-1';
        if (front === back && front !== '9' && k > 1) {
            arrayA[current] = '9';
            arrayA[arrayA.length - current - 1] = '9';
            k -= 2;
        } else if (front !== back && k > 0) {
            if (front !== '9') {
                arrayA[current] = '9';
                k -= 1;
            }

            if (back !== '9') {
                arrayA[arrayA.length - current - 1] = '9';
                k -= 1;
            }
        };

        setPalindromeRecursive(n, current + 1)
    };

    // recursive function untuk salin jadi string
    function createResultRecursive(n, current = 0) {
        if (current >= n) {
            return;
        };

        result += arrayA[current];
        createResultRecursive(n, current + 1)
    };

    // execute function
    createArrRecursive(a.length);

    // return hasil
    return result !== '' ? result : '-1';
}

module.exports = highestPalindrome;