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
            Stack<char> stack = new Stack<char>();

            foreach (char c in a)
            {
                if (c == '(' || c == '{' || c == '[')
                {
                    stack.Push(c);
                }
                else if (c == ')' || c == '}' || c == ']')
                {
                    if (stack.Count == 0)
                    {
                        return "NO";
                    }

                    char top = stack.Pop();

                    if ((c == ')' && top != '(') || (c == '}' && top != '{') || (c == ']' && top != '['))
                    {
                        return "NO";
                    }
                }
            }
           return stack.Count == 0 ?  "YES" :  "NO";
        }
    }
}
