using dotNET;

namespace Test
{
    public class Tests
    {
        [SetUp]
        public void Setup()
        {
        }

        [Test]
        public void TestBalanceBracket()
        {
            Assert.Multiple(() =>
            {
                Assert.That(Program.BalanceBracket("()"), Is.EqualTo("YES"));
                Assert.That(Program.BalanceBracket("({[]})"), Is.EqualTo("YES"));
                Assert.That(Program.BalanceBracket("({[})"), Is.EqualTo("NO"));
            });
        }

        [Test]
        public void TestHighestPalindrome()
        {
            Assert.Multiple(() =>
            {
                Assert.That(Program.HighestPalindrome("3943", 1), Is.EqualTo("3993"));
                Assert.That(Program.HighestPalindrome("932239", 2), Is.EqualTo("992299"));
            });
        }

        [Test]
        public void TestWeightedStrings()
        {
            CollectionAssert.AreEqual(new[] { "YES", "YES", "YES", "NO" }, Program.WeightedString("abbcccd", [1, 3, 9, 8]));
            CollectionAssert.AreEqual(new[] { "NO", "YES", "NO", "YES", "YES", "NO" }, Program.WeightedString("aaabbbcccddd", [5, 9, 7, 8, 12, 5]));
        }
    }
}