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
    try {
        return highestPalindromHandler(a,k)
    } catch (error) {
        return "-1";
    }
}

const highestPalindromHandler = (a, s) => {
        let n = a.length;

        if (n === 0) {
            return "";
        }
        if (n === 1) {
            return s >= 1 ? "9" : a;
        }

        const left = a[0];
        const right = a[n - 1];

        if (left === right) {
            let newS = s;
            if (left !== "9" && s >= 2) newS -= 2; 
            const result = highestPalindromHandler(a.substring(1, a.length - 1), newS);
            return (s >= 2) ? "9" + result + "9" : left + result + right;
        } else {
            let newS = s;
            let temp = "";

            if (newS >= 2) {
                newS -= 2;
                temp = "99";
            } else if (newS >= 1) {
                newS -= 1;
                temp = (left > right ? left : right) + (left > right ? left : right);
            } else {
                throw new Error();
            }

            const result = highestPalindromHandler(a.substring(1, n - 1), newS);
            return temp.slice(0, 1) + result + temp.slice(1);
        }
}

module.exports = highestPalindrome;