/**
 * Aturan:
 * 1. Tanda braket / karakter yang diperbolehkan sebagai berikut: ( , ) , { , } , atau [ , ].
 * 2. Bracket bisa dipisahkan dengan atau tanpa whitespace.
 * 3. Periksa tanda kurung yang memiliki kecocokan antara braket buka dan braket tutup dengan mengembalikan nilai string YES atau NO.
 *
 * Soal:
 * 1. Buat fungsi untuk menemukan Balanced Bracket dengan kompleksitas paling rendah!
 * 2. Jelaskan kompleksitas dari penyelesaianmu untuk No.3 dan cantumkan di README Repo!
 */

function balanceBracket(a) {
  let stack = [];

  const matchingBrackets = {
    ")": "(",
    "}": "{",
    "]": "[",
  };

  for (let i = 0; i < a.length; i++) {
    let char = a[i];

    if (char === "(" || char === "{" || char === "[") {
      stack.push(char);
    } else if (char === ")" || char === "}" || char === "]") {
      if (stack.length === 0 || stack.pop() !== matchingBrackets[char]) {
        return "NO";
      }
    } else if (char.trim() !== "") {
      // Check for any non-bracket and non-whitespace character
      return "NO";
    }
  }

  return stack.length === 0 ? "YES" : "NO";
}



module.exports = balanceBracket;
