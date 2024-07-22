using NUnit.Framework;
using dotNET; // Pastikan namespace ini sesuai dengan namespace kelas Program

namespace Test
{
    // [SetUpFixture] untuk inisialisasi global yang dijalankan sekali untuk semua pengujian
    [SetUpFixture]
    public class GlobalSetup
    {
        [OneTimeSetUp]
        public void GlobalSetupMethod()
        {
            // Inisialisasi global sebelum semua pengujian dijalankan
        }

        [OneTimeTearDown]
        public void GlobalTeardownMethod()
        {
            // Cleanup global setelah semua pengujian selesai
        }
    }

    [TestFixture]
    public class BalanceBracketTests
    {
        [SetUp]
        public void Setup()
        {
            // Inisialisasi sebelum setiap pengujian di kelas ini dijalankan
        }

        [TestCase("{ [ ( ] ) }", "NO")]
        [TestCase("{ ( ( [ ] ) [ ] ) [ ] }", "YES")]
        [TestCase("{[()]}", "YES")]
        [TestCase("{(([|])[])[]}", "NO")]
        public void TestBalanceBracket(string input, string expected)
        {
            // Pengujian untuk fungsi BalanceBracket
            Assert.That(Program.BalanceBracket(input), Is.EqualTo(expected));
        }
    }

    [TestFixture]
    public class HighestPalindromeTests
    {
        [SetUp]
        public void Setup()
        {
            // Inisialisasi sebelum setiap pengujian di kelas ini dijalankan
        }

        [TestCase("3943", 1, "3993")]
        [TestCase("932239", 2, "992299")]
        [TestCase("092282", 3, "992299")]
        [TestCase("5566", 1, "-1")]
        [TestCase("11331", 4, "99399")]
        [TestCase("A1341", 1, "-1")]
        public void TestHighestPalindrome(string input, int k, string expected)
        {
            // Pengujian untuk fungsi HighestPalindrome
            Assert.That(Program.HighestPalindrome(input, k), Is.EqualTo(expected));
        }
    }

    [TestFixture]
    public class WeightedStringTests
    {
        [SetUp]
        public void Setup()
        {
            // Inisialisasi sebelum setiap pengujian di kelas ini dijalankan
        }
        [TestCase("abbcccd", new int[] { 1, 3, 9, 8 }, new string[] { "YES", "YES", "YES", "NO" })]
        [TestCase("aaabbbcccddd", new int[] { 5, 9, 7, 8, 12, 5 }, new string[] { "NO", "YES", "NO", "YES", "YES", "NO" })]
        [TestCase("getitssimple", new int[] { 1, 2, 8, 7, 12, 38, 10 }, new string[] { "NO", "NO", "NO", "YES", "YES", "YES", "NO" })]
        public void TestWeightedStrings(string input, int[] weights, string[] expected)
        {
            // Pengujian untuk fungsi WeightedString
            CollectionAssert.AreEqual(expected, Program.WeightedString(input, weights));
        }
    }
}