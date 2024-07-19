using System.Net.NetworkInformation;

namespace dotNET
{
    public static partial class Program
    {
        /// <summary>
        /// Aturan:
        /// 1. Tanda braket / karakter yang diperbolehkan sebagai berikut: ( , ) , { , } , atau[,]. 
        /// 2. Bracket bisa dipisahkan dengan atau tanpa whitespace.
        /// 3. Periksa tanda kurung yang memiliki kecocokan antara braket buka dan braket tutup dengan mengembalikan nilai string YES atau NO.
        /// 
        /// Soal:
        /// 1. Buat fungsi untuk menemukan Balanced Bracket dengan kompleksitas paling rendah!
        /// 2. Jelaskan kompleksitas dari penyelesaianmu untuk No.3 dan cantumkan di README Repo! 
        /// </summary>
        /// <param name="a"></param>
        /// <returns></returns>
        public static string BalanceBracket(string a)
        {
            if (string.IsNullOrWhiteSpace(a))
                throw new ArgumentNullException(nameof(a));

            // rubah string menjadi 
            var charText = a.ToCharArray();

            Stack<char> st = new();
            // loping tiap karkter
            for (int i = 0; i < charText.Length; i++)
            {
                // jika charText[i] diawali dengan bracket maka push ke stack
                if (charText[i] == '{' || charText[i] == '(' || charText[i] == '[')
                    st.Push(charText[i]);

                // jika charText[i] di akhiri bracket maka pop dari stack dan
                // cek jika nilai yang dikeluarkan memiliki pasangan yang serasi
                if (charText[i] == '}' || charText[i] == ')' || charText[i] == ']')
                {
                    // jika stack sudah kosong maka bracket ini tidak ada pasanganya
                    if (st.Count == 0)
                        return "NO";

                    // Pop pada stack terakhir cocokan dengan charText[i]
                    else if (!IsBalanceBracket(st.Pop(), charText[i]))
                        return "NO";
                }
            }

            if (st.Count == 0)
                return "YES";
            else
                return "NO";
        }

        /// <summary>
        /// fungsi untuk mencocokan braket yan serasi dengan inputan 2 karakter
        /// </summary>
        /// <param name="character1">(char) karakter yang berisi '(', '{', atau '['</param>
        /// <param name="character2">(char) karakter yang berisi '(', '{', atau '['</param>
        /// <returns>
        /// (bool) jika karater character1 tertutup oleh character2 maka return true jika tidak maka false
        /// </returns>
        private static bool IsBalanceBracket(char character1,
                                      char character2)
        {
            if (character1 == '(' && character2 == ')')
                return true;
            else if (character1 == '{' && character2 == '}')
                return true;
            else if (character1 == '[' && character2 == ']')
                return true;
            else
                return false;
        }
    }
}
