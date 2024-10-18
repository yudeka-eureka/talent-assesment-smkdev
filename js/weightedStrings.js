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


function weightedString(s = 'abbcccd', queries = [1, 3, 9, 8]) {
    let weights = new Set();  // Set to store all possible weights
    
    let n = s.length;
    let i = 0;

    // Calculate weights for each character and consecutive substrings
    while (i < n) {
        let currentChar = s[i];
        let weight = currentChar.charCodeAt(0) - 'a'.charCodeAt(0) + 1;  // Get weight of the character
        let sum = 0;
        let count = 0;

        // Loop to handle consecutive occurrences of the same character
        while (i < n && s[i] === currentChar) {
            count++;
            sum += weight;
            weights.add(sum);  // Add each cumulative substring weight to the set
            i++;
        }
    }

    // Process queries
    return queries.map(query => weights.has(query) ? "Yes" : "No");
}

module.exports = weightedString;