# Balanced Bracket Function

## Deskripsi

Fungsi ini memeriksa apakah sebuah string yang berisi tanda kurung adalah seimbang. Karakter tanda kurung yang diizinkan hanya `(`, `)`, `{`, `}`, `[`, dan `]`. Fungsi ini mengembalikan `"YES"` jika tanda kurung seimbang dan `"NO"` jika tidak.

## Solusi

Fungsi `balanceBracket` bekerja dengan menggunakan struktur data stack untuk melacak tanda kurung buka. Untuk setiap karakter dalam string, jika itu adalah tanda kurung buka, maka didorong ke stack. Jika itu adalah tanda kurung tutup, fungsi akan memeriksa apakah stack kosong atau apakah bagian atas stack bukan tanda kurung buka yang sesuai. Jika salah satu dari kondisi ini terpenuhi, fungsi akan mengembalikan `"NO"`. Setelah memproses semua karakter, fungsi akan mengembalikan `"YES"` jika stack kosong dan `"NO"` jika tidak.

## Kode

```javascript
function balanceBracket(a) {
    const stack = [];
    const bracketPairs = {
        ')': '(',
        '}': '{',
        ']': '['
    };

    for (let char of a) {
        if (char === '(' || char === '{' || char === '[') {
            stack.push(char);
        } else if (char === ')' || char === '}' || char === ']') {
            if (stack.length === 0 || stack.pop() !== bracketPairs[char]) {
                return "NO";
            }
        } else {
            // Jika ditemukan karakter yang tidak valid
            return "NO";
        }
    }

    return stack.length === 0 ? "YES" : "NO";
}
```

## Kompleksitas Waktu

Kompleksitas waktu dari fungsi ini adalah **O(n)**, di mana `n` adalah panjang string input. Ini karena setiap karakter dalam string diproses tepat sekali. Operasi push dan pop pada stack keduanya membutuhkan waktu konstan, O(1).

## Kompleksitas Ruang

Kompleksitas ruang dari fungsi ini adalah **O(n)** dalam kasus terburuk, di mana `n` adalah panjang string input. Hal ini terjadi ketika semua karakter dalam string adalah tanda kurung buka dan semuanya didorong ke stack.


```javascript
  console.log(balanceBracket("{ [ ( ] ) }")); // Output: NO
  console.log(balanceBracket("{ ( ( [ ] ) [ ] ) [ ] }")); // Output: YES
  console.log(balanceBracket("{[()]}")); // Output: YES
  console.log(balanceBracket("{(([|])[])[]}")); // Output: NO
```
