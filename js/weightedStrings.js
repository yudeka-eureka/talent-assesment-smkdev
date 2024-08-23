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

function weightedString(a = "abbcccd", b = [1, 3, 9, 8]) {
  const weights = new Set();
  let count = 1;

  for (let i = 0; i < a.length; i++) {
    const weight = a.charCodeAt(i) - 96; // 'a' has a weight of 1, so subtract 96
    weights.add(weight * count);

    if (i < a.length - 1 && a[i] === a[i + 1]) {
      count++;
      weights.add(weight * count);
    } else {
      count = 1;
    }
  }

  return b.map((weight) => (weights.has(weight) ? "YES" : "NO"));
}

module.exports = weightedString;
