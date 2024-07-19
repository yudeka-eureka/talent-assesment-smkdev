using System.Net.NetworkInformation;

namespace dotNET
{
    public static partial class Program
    {
        /// <summary>
        /// Alfabet dari a sampai z memiliki bobot sebesar angka urutannya, 
        /// misal: a = 1, b = 2, c = 3, ..., z = 26. 
        /// Bobot sebuah string merupakan penjumlahan bobot dari 
        /// kesuluruhan karakter. 
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
            if (string.IsNullOrWhiteSpace(a))
                throw new ArgumentNullException(nameof(a));
            if (k == null || k.Count == 0)
                throw new ArgumentNullException(nameof(k));

            // rubah string menjadi 
            var charText = a.ToCharArray();

            var previousChar = -1;
            var currentChar = -1;
            var uniformStringSum = 0;
            var uniformWeightSet = new HashSet<int>();

            // loping tiap karkter
            for (int i = 0; i < charText.Length; i++)
            {
                currentChar = charText[i];
                if (currentChar != previousChar)
                    uniformStringSum = currentChar - 96;
                else
                    uniformStringSum += currentChar - 96;

                uniformWeightSet.Add(uniformStringSum);
                previousChar = currentChar;
            }

            var result = new string[k.Count];
            for (int i = 0; i < k.Count; i++)
            {
                var uniformWeight = k.ElementAt(i);
                if (uniformWeightSet.Contains(uniformWeight))
                    result[i] = "YES";
                else
                    result[i] = "NO";
            }

            return result;
        }
    }
}
