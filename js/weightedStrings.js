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

function weightedStrings(a = 'abbcccd', b = [1,3,9,8]) {

  let weights = new Set();

  let i = 0;
  while (i < a.length) {
    const char = a[i];
    const weight = char.charCodeAt(0) - 'a'.charCodeAt(0) + 1;

    let count = 1;
    while (((i + 1) < a.length) && a[i+1] == char) {
      count++;
      i++;
    }

    for (let j = 1; j <= count; j++){
      weights.add(weight * j);
    } 

    i++;
  }

  let result = [];

  for (let weight of b) {
    if (weights.has(weight)){
      result.push("YES");
    } else {
      result.push("NO");
    }
  }

  return result;

}

module.exports = weightedStrings;
