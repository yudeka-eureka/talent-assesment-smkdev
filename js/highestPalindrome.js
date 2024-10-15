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

function highestPalindrome(s = '3993', k = 1) {

  if (isNaN(s)) {
    return '-1';
  }

  // Untuk deteksi perubahan
  let changes = new Array(s.length).fill(false);

  // Untuk normalisasi dan buat semua jadi palindrome dengan ambil max dan samakan di antara 2 angka yaitu left dan right
  const helper = (arr, left, right, k) => {
    
      if (left >= right) return k >= 0 ? k : '-1';

      if (arr[left] !== arr[right] && k > 0) {
        const maxAngka = Math.max(arr[left], arr[right]);
        arr[left] = arr[right] = maxAngka.toString();
        changes[left] = changes[right] = true;

        k--;

        if (k < 0) return '-1';
      }

      return helper(arr, left + 1, right - 1, k);
  }

  // fungsi untuk mengambil angka palindrome yang maksimal
  const maxPalindrome = (arr, left, right, k) => {


      if (left > right || k <= 0) return arr.join('');

      if (left === right) {
        if (arr[left] !== 9 && k > 0) {
            arr[left] = '9';
            k--;
        }
        return arr.join('');
      }
      else 
      {
        if (arr[left] !== '9'){

            if ((changes[left] || changes[right])) {
              if (k > 0) {
                arr[left] = arr[right] = '9';
                k--;
              }
            }
            else {
              if (k >= 2){
                arr[left] = arr[right] = '9';
                k-=2;
              }
            }

        }
      }


      return maxPalindrome(arr, left + 1, right - 1, k);

  }


  let arr = s.split('');
  let k_temp = helper(arr, 0, arr.length - 1, k);

  if (k_temp == -1) return k_temp;


  // Bandingkan jika tidak membentuk palindrome, maka return -1

  let result = maxPalindrome(arr, 0, arr.length - 1, k_temp);
  let arr_cmp1 = [...arr];
  let arr_cmp2 = [...arr].reverse();
  // if (arr.join('') != arr.reverse.join(''));
  if (arr_cmp1.join('') !== arr_cmp2.join('')){
    return '-1';
  }

  return result;

}

module.exports = highestPalindrome;