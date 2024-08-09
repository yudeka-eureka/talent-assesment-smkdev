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
    // untuk push bracket pembuka
    const stack = [];
    const matchingBrackets = {
        ')': '(',
        ']': '[',
        '}': '{'
    };

    // looping untuk pemeriksaan input
    for (let char of a) {
        // cek pembuka
        if (char === '(' || char === '[' || char === '{') {
            // push kedalam stack
            stack.push(char);
        // cek penutup
        } else if (char === ')' || char === ']' || char === '}') {
          // jika tidak ada pembuka return NO
          // pop untuk menghilangkan stack terakhir sekaligus mengecek kembaliannya jika tidak match return NO juga
            if (stack.length === 0 || stack.pop() !== matchingBrackets[char]) {
                return 'NO';
            } // Jika selain bracket atau spasi, return NO
        } else if (char !== ' ') {
          return 'NO';
        }
    };

    // jika stacknya 0 return YES dan sebaliknya 
    return stack.length === 0 ? 'YES' : 'NO';
}

module.exports = balanceBracket;