using System.Net.NetworkInformation;

namespace dotNET
{
    public static partial class Program
    {
        /// <summary>
        /// Kamu memiliki string yang merepresentasikan angka 3943 
        /// lalu diberikan sebuah variabel k untuk melakukan 
        /// replacement karakter sejumlah k pada string
        /// untuk mendapatkan bentuk palindrom dengan nilai tertinggi. 
        /// 
        /// Aturan:
        /// 1. Jika dari sebuah string tidak ditemukan bentuk palindrome-nya 
        /// meski sudah melakukan replacement dan 
        /// tidak merepresentasikan sebuah angka maka akan mengeluarkan output -1.
        /// 2. Tidak boleh menggunakan fungsi bawaan/tools untuk pencarian/filter/sort.
        /// 3. Tidak boleh menggunakan looping.
        /// 4. Hanya diperkenankan menggunakan rekursif.
        /// 
        /// Soal:
        /// Buat fungsi yang digunakan untuk menyelesaikan permasalahan Highest Palindrome!
        /// </summary>
        /// <param name="a"></param>
        /// <param name="k"></param>
        /// <returns></returns>
        public static string HighestPalindrome(string a, int k)
        {
            if (string.IsNullOrWhiteSpace(a))
                throw new ArgumentNullException(nameof(a));

            char[] charText = a.ToCharArray();

            // Step 1: Make the string a palindrome with minimal changes
            if (!MakePalindrome(charText, 0, a.Length - 1, ref k))
                return "-1";

            // Step 2: Maximize the palindrome to the highest possible value
            MaximizePalindrome(charText, 0, a.Length - 1, ref k);

            return new string(charText);
        }

        private static bool MakePalindrome(char[] charText, int index, int totalLength, ref int k)
        {
            if (index >= totalLength)
                return k >= 0;

            if (charText[index] != charText[totalLength])
            {
                charText[index] = charText[totalLength] = (charText[index] > charText[totalLength] ? charText[index] : charText[totalLength]);
                k--;
            }

            return MakePalindrome(charText, index + 1, totalLength - 1, ref k);
        }

        private static void MaximizePalindrome(char[] charText, int index, int totalLength, ref int k)
        {
            if (index > totalLength || k <= 0)
                return;

            if (index == totalLength)
            {
                if (k > 0)
                    charText[index] = '9';

                return;
            }

            if (charText[index] < '9')
            {
                if (k >= 2 && charText[index] == charText[index] && charText[totalLength] == charText[totalLength])
                {
                    k -= 2;
                    charText[index] = charText[totalLength] = '9';
                }
                else if (k >= 1 && (charText[index] != charText[index] || charText[totalLength] != charText[totalLength]))
                {
                    k--;
                    charText[index] = charText[totalLength] = '9';
                }
            }

            MaximizePalindrome(charText, index + 1, totalLength - 1, ref k);
        }

    }
}
