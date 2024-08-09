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
  
    // menentukan bobot a-z
    let weights = {};
    for (let i = 0; i < 26; i++) {
      weights[String.fromCharCode(97 + i)] = i + 1;
    };
    
    // memproses awal input a
    let rawA = [];
    let tempObj = {};
    
    for (let i=0; i< a.length; i++) {
      tempObj.char = a[i];
      tempObj.count = tempObj.count ?? 1;
      tempObj.weight = weights[a[i]];
      if (a[i] === a[i+1]) {
        tempObj.count++
      } else {
        rawA.push(tempObj);
        tempObj = {}
      }
    };
    
    // proses akhir input a
    let finalA = [];
    rawA.forEach((each) => {
      for(let i = 1; i <= each.count; i++){
        finalA.push(each.weight * i)
      }
    });
    
    // return hasil
    return b.map(each => finalA.some(eachA => eachA === each) ? 'YES' : 'NO');
}

module.exports = weightedString;