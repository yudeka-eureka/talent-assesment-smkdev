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
String balanceBracket(String a) {
  List<String> stack = [];
  Map<String, String> matchingBracket = {')': '(', '}': '{', ']': '['};

  for (int i = 0; i < a.length; i++) {
    String char = a[i];

    if (char == '(' || char == '{' || char == '[') {
      stack.add(char); // Push opening brackets onto the stack
    } else if (char == ')' || char == '}' || char == ']') {
      if (stack.isEmpty || stack.last != matchingBracket[char]) {
        return "NO"; // Unmatched or misordered bracket
      }
      stack.removeLast(); // Pop the matched opening bracket
    }
  }

  return stack.isEmpty
      ? "YES"
      : "NO"; // If stack is empty, brackets are balanced
}

/**
 * Kamu memiliki string yang merepresentasikan angka 3943 
 * lalu diberikan sebuah variabel k untuk melakukan 
 * replacement karakter sejumlah k pada string
 * untuk mendapatkan bentuk palindrom dengan nilai tertinggi. 
 */

/* Aturan:
 * 1. Jika dari sebuah string tidak ditemukan bentuk palindrome-nya 
 * meski sudah melakukan replacement dan 
 * tidak merepresentasikan sebuah angka maka akan mengeluarkan output -1.
 * 2. Tidak boleh menggunakan fungsi bawaan/tools untuk pencarian/filter/sort.
 * 3. Tidak boleh menggunakan looping.
 * 4. Hanya diperkenankan menggunakan rekursif.
 * 

Soal:
Buat fungsi yang digunakan untuk menyelesaikan permasalahan Highest Palindrome! */
String highestPalindrome(String a, int k) {
  int len = a.length;
  List<String> chars = a.split('');
  int left = 0, right = len - 1;
  List<bool> changes = List.filled(len, false);

  // Step 1: Make the string a palindrome with minimal changes
  while (left < right) {
    if (chars[left] != chars[right]) {
      String maxChar = (chars[left].compareTo(chars[right]) > 0)
          ? chars[left]
          : chars[right];
      chars[left] = chars[right] = maxChar;
      changes[left] = changes[right] = true;
      k--;
    }
    left++;
    right--;
  }

  if (k < 0) return "-1";

  // Step 2: Maximize the palindrome value with remaining changes
  left = 0;
  right = len - 1;
  while (left < right) {
    if (chars[left] != "9") {
      if (changes[left] && k > 0) {
        chars[left] = chars[right] = "9";
        k--;
      } else if (!changes[left] && k >= 2) {
        chars[left] = chars[right] = "9";
        k -= 2;
      }
    }
    left++;
    right--;
  }

  // Handle the middle character for odd length
  if (len % 2 == 1 && k > 0) {
    chars[(len / 2).floor()] = "9";
  }

  return chars.join('');
}

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
List<String> weightedStrings(String a, List<int> b) {
  // set to store all possible wight
  Set<int> weightSet = {};

  int i = 0;
  while (i < a.length) {
    int charWeight = a.codeUnitAt(i) - 'a'.codeUnitAt(0) + 1;
    int currentWeight = charWeight;
    weightSet.add(currentWeight);

    // Check for consecutive characters
    int j = i + 1;
    while (j < a.length && a[j] == a[i]) {
      currentWeight += charWeight;
      weightSet.add(currentWeight);
      j++;
    }

    // Move to the next new character
    i = j;
  }

  // Step 3: Evaluate each query
  List<String> result = [];
  for (int query in b) {
    if (weightSet.contains(query)) {
      result.add("YES");
    } else {
      result.add("NO");
    }
  }

  return result;
}
