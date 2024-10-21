package dev.smk.java.app;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class HighestPalindromeTest {

    @Test
    void testHighestPalindromeNoChangesNeeded() {
        assertEquals("1344", HighestPalindrome.highestPalindrome("1344", 0));
    }

    @Test
    void testHighestPalindromeOneChangeNeeded() {
        assertEquals("1441", HighestPalindrome.highestPalindrome("1341", 1));
    }

    @Test
    void testHighestPalindromeMultipleChangesNeeded() {
        assertEquals("9999", HighestPalindrome.highestPalindrome("1234", 4));
    }

    @Test
    void testHighestPalindromeNotPossible() {
        assertEquals("-1", HighestPalindrome.highestPalindrome("1234", 1));
    }
}
