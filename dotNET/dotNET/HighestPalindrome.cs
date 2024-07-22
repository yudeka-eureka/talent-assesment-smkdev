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
            string left = a.Substring(0, a.Length / 2);
            string right = a.Substring((a.Length + 1) / 2);
            if (a.Length % 2 == 1)
            {
                right = right.Substring(1);
            }

            if (k == 0)
            {
                if (left == Reverse(right))
                {
                    return a;
                }
                return "-1";
            }
            else
            {
                if (left == Reverse(right))
                {
                    return HighestPalindrome(a, k - 1);
                }
                else
                {
                    if (left.Length > right.Length)
                    {
                        return HighestPalindrome(left + left[left.Length - 1] + Reverse(left), k - 1);
                    }
                    else
                    {
                        return HighestPalindrome(left + right[0] + Reverse(left), k - 1);
                    }
                }
            }
        }

        public static string Reverse(string s)
        {
            char[] charArray = s.ToCharArray();
            Array.Reverse(charArray);
            return new string(charArray);
        }

    }
}
