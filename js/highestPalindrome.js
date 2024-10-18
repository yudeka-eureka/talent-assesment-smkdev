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

function highestPalindrome(a = '3993',k = 1) {
    // Helper function for recursion
    function helper(arr, left, right, k) {
        // Base case: if pointers have crossed, return
        if (left >= right) return arr.join('');

        // If the characters at the current left and right are not equal
        if (arr[left] !== arr[right]) {
            // Make the two characters equal, pick the maximum value
            const maxChar = Math.max(arr[left], arr[right]);
            arr[left] = arr[right] = maxChar;
            k--; // Decrease the allowed changes
        }

        // If we have no more allowed changes and still need to process
        if (k < 0) return '-1';

        // Recursively process the next pair
        return helper(arr, left + 1, right - 1, k);
    }

    // Start the recursive process of creating a palindrome
    let arr = s.split(''); // Convert string to array for easy manipulation
    let result = helper(arr, 0, arr.length - 1, k);

    // Check if palindrome is achieved and if changes are still available
    if (result === '-1') return result;

    // Try to maximize the palindrome by converting digits to '9' if changes are available
    arr = result.split(''); // Convert the result back to array
    function maximize(arr, left, right, k) {
        if (left >= right) return arr.join('');

        // Try to change both arr[left] and arr[right] to '9' if possible
        if (arr[left] !== '9') {
            if (arr[left] === arr[right]) {
                // If both were already equal, it will take 2 changes to set both to '9'
                if (k >= 2) {
                    arr[left] = arr[right] = '9';
                    k -= 2;
                }
            } else {
                // If they were changed to become a palindrome, just 1 more change needed
                if (k >= 1) {
                    arr[left] = arr[right] = '9';
                    k -= 1;
                }
            }
        }

        // Recursively maximize the next pair
        return maximize(arr, left + 1, right - 1, k);
    }

    return maximize(arr, 0, arr.length - 1, k);
}



module.exports = highestPalindrome;