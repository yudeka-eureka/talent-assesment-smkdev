package talent.assesment.smkdev;

import org.junit.Test;
import static org.junit.Assert.*;

public class HighestPalindromeTest {
    @Test
    public void testHighestPalindrome1() {
        assertEquals("3993", HighestPalindrome.highestPalindrome("3943", 1));
    }

    @Test
    public void testHighestPalindrome2() {
        assertEquals("992299", HighestPalindrome.highestPalindrome("092282", 3));
    }

    @Test
    public void testHighestPalindrome3() {
        assertEquals("-1", HighestPalindrome.highestPalindrome("5566", 1));
    }

    @Test
    public void testHighestPalindrome4() {
        assertEquals("992299", HighestPalindrome.highestPalindrome("932239", 2));
    }

    @Test
    public void testHighestPalindrome5() {
        assertEquals("99399", HighestPalindrome.highestPalindrome("11331", 4));
    }

    @Test
    public void testHighestPalindrome6() {
        assertEquals("-1", HighestPalindrome.highestPalindrome("A1341", 1));
    }
}
