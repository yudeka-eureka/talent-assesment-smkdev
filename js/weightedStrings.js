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
    const alphabetWeight = (char) => char.charCodeAt(0) - 96;

    let weights = new Set();

    let currentChar = a[0];
    let currentCount = 1;
    weights.add(alphabetWeight(currentChar));

    for (let i = 1; i < a.length; i++) {
        if (a[i] === currentChar) {
            currentCount++;
        } else {
            currentChar = a[i];
            currentCount = 1;
        }
        for (let j = 1; j <= currentCount; j++) {
            weights.add(alphabetWeight(currentChar) * j);
        }
    }

    // return ["YES", "YES", "YES", "NO"];
    return b.map(value => weights.has(value) ? "YES" : "NO");
}

module.exports = weightedString;
