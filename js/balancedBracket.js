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


function balanceBracket(s) {
    // Stack to store the opening brackets
    let stack = [];

    // Map to match closing brackets to opening brackets
    const bracketPairs = {
        ')': '(',
        ']': '[',
        '}': '{'
    };

    // Traverse each character in the string
    for (let char of s) {
        // Ignore whitespace
        if (char === ' ') continue;

        // If it's an opening bracket, push it onto the stack
        if (char === '(' || char === '[' || char === '{') {
            stack.push(char);
        }
        // If it's a closing bracket
        else if (char === ')' || char === ']' || char === '}') {
            // Check if stack is empty or top of the stack doesn't match the corresponding opening bracket
            if (stack.length === 0 || stack.pop() !== bracketPairs[char]) {
                return "NO";
            }
        }
    }

    // If stack is empty, all brackets were balanced
    return stack.length === 0 ? "YES" : "NO";
}


module.exports = balanceBracket;