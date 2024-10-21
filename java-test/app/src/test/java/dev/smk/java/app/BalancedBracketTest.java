package dev.smk.java.app;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class BalancedBracketTest {

    @Test
    void testBalancedBrackets() {
        assertEquals("YES", BalancedBracket.isBalanced("()"));
        assertEquals("YES", BalancedBracket.isBalanced("[]"));
        assertEquals("YES", BalancedBracket.isBalanced("{}"));
        assertEquals("YES", BalancedBracket.isBalanced("({[]})"));
    }

    @Test
    void testUnbalancedBrackets() {
        assertEquals("NO", BalancedBracket.isBalanced("(]"));
        assertEquals("NO", BalancedBracket.isBalanced("([)]"));
        assertEquals("NO", BalancedBracket.isBalanced("{[}"));
        assertEquals("NO", BalancedBracket.isBalanced("((("));
    }
    //Create test BalancedBracket from class BalancedBracket with function 
}
