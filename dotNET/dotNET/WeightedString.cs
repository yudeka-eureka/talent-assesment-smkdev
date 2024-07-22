using System.Net.NetworkInformation;

namespace dotNET
{
    public static partial class Program
    {
        /// <summary>
        /// Alfabet dari a sampai z memiliki bobot sebesar angka urutannya, 
        /// misal: a = 1, b = 2, c = 3, ..., z = 26. 
        /// Bobot sebuah string merupakan penjumlahan bobot dari 
        /// keseluruhan karakter. 
        /// Contoh: gits = 7 + 9 + 20 + 19 = 55
        /// 
        /// Aturan:
        /// Jika terdapat karakter yang berulang dan berurutan perlu dirincikan ke dalam bentuk substring dari perulangan pertama hingga n. 
        /// Contoh: bbccc => b, bb, c, cc, ccc. 
        /// 
        /// Soal: 
        /// Buat fungsi untuk menyelesaikan permasalahan Weighted Strings!
        /// </summary>
        /// <param name="a"></param>
        /// <param name="k"></param>
        /// <returns></returns>
        public static string[] WeightedString(string a, ICollection<int> k)
        {
            List<string> substrings = new List<string>();
            int weight = 0;
            int count = 1;

            for (int i = 0; i < a.Length; i++)
            {
                weight += a[i] - 'a' + 1;

                if (i < a.Length - 1 && a[i] == a[i + 1])
                {
                    count++;
                }
                else
                {
                    for (int j = 0; j < count; j++)
                    {
                        substrings.Add(a.Substring(i - count + 1, count));
                    }
                    count = 1;
                }
            }

            string[] result = new string[substrings.Count];
            for (int i = 0; i < substrings.Count; i++)
            {
                if (k.Contains(weight))
                {
                    result[i] = substrings[i];
                }
                else
                {
                    result[i] = "";
                }
            }

            return result;
        }
    }
}
