# Komplektifitas
## Kompleksitas Waktu (Time Complexity):
Kode tersebut memeriksa setiap karakter dalam string a satu per satu. Karena pemeriksaan ini dilakukan sekali untuk setiap karakter, waktu yang diperlukan adalah O(n), di mana n adalah jumlah karakter dalam string.

## Kompleksitas Ruang (Space Complexity):
Kode tersebut juga menggunakan sebuah stack untuk menyimpan karakter pembuka seperti "(", "[", dan "{". Semua karakter di string bisa menjadi pembuka, sehingga stack bisa menampung hingga n karakter. Oleh karena itu, ruang yang dibutuhkan juga O(n).