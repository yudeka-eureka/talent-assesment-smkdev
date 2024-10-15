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

function balanceBracket(a){
  if (typeof a == 'undefined' || a == null || a.toString().trim() == "") {
    return "NO";
  }
  else {

      const arrStack = [];
      const bracketPairs = {
        ")":"(",
        "]":"[",
        "}":"{"
      }

      for (let char of a){
        if (char != " ") {
          if (char === "(" || char === "[" || char === "{") {
            arrStack.push(char);
          }
          else if (char === ")" || char === "]" || char === "}") {
            if (arrStack.length === 0 || arrStack.pop() !== bracketPairs[char]) {
              return "NO";
            }
          }
          else {
            return "NO";
          }
        }
      }

      return arrStack.length === 0 ? "YES" : "NO";
  }
}

module.exports = balanceBracket;