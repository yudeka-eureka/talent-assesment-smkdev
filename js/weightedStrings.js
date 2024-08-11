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


function weightedString(a = 'abbcccd', b = [1, 3, 9, 8]) {
    // code disini
    let weights = new Set();

    for (let i = 0; i < a.length; i++) {
        let weight = 0;
        for (let j = i; j < a.length && a[i] === a[j]; j++) {
            weight += a[i].charCodeAt(0) - 96;
            weights.add(weight);
        }
    }

    return b.map((val) => (weights.has(val) ? "YES" : "NO"));
}

module.exports = weightedString;