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
    const alphabetWeight = (char) => char.charCodeAt(0) - 'a'.charCodeAt(0) + 1;
    
    const getWeights = (str) => {
        const weights = new Set();
        let start = 0;
        while (start < str.length) {
            let weight = 0;
            for (let end = start; end < str.length; end++) {
                if (end > start && str[end] !== str[start]) break;
                weight += alphabetWeight(str[end]);
                weights.add(weight);
            }
            start++;
        }
        return weights;
    };

    const weights = getWeights(a);
    return b.map(weight => weights.has(weight) ? "YES" : "NO");
}

module.exports = weightedString;