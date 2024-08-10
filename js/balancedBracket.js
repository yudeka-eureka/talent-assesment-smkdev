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
    const list = []
    for (const item of a) {
        if (item === " ") continue;
        if ("{([".includes(item)) list.push(item);
        else if ("})]".includes(item)) {
            const check = list.pop()
            if ((check === "{" && item === "}") ||
                (check === "(" && item === ")") ||
                (check === "[" && item === "]")) 
                {
                continue;
                }
             return "NO"
        } else return "NO"
    }
    return list.length === 0 ? "YES" : "NO";
}

module.exports = balanceBracket;