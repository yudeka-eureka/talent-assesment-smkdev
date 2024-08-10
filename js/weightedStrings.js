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

function weightedString(s, weights) {
    const weightsSet = new Set();
    for (let i = 0; i < s.length; i++) {
        let j = i;
        while(s[j] === s[i]) {
            const weight = calculateWeight(s[j]) * (j - i + 1);
            weightsSet.add(weight);
            j++;
        }
        i = j - 1;
    }

    console.log(weightsSet);

    return weights.map(weight => weightsSet.has(weight) ? "YES" : "NO");
}

function calculateWeight(char) {
    return char.toLowerCase().charCodeAt(0) - "a".charCodeAt(0) + 1;
}

module.exports = weightedString;