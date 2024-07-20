<?php
// tests/SumTest.php
use PHPUnit\Framework\TestCase;

class HighestPalindromeTest extends TestCase {
    public function testHighestPalindrome() {
        require 'HighestPalindrome.php';
        $this->assertEquals("3993", highestPalindrome("3943", 1));
        $this->assertEquals("992299", highestPalindrome("092282", 3));


        $this->assertEquals("-1", highestPalindrome("5566", 1));
        $this->assertEquals("992299", highestPalindrome("932239", 2));
        $this->assertEquals("99399", highestPalindrome("11331", 4));
        $this->assertEquals("-1", highestPalindrome("A1341", 1));
        
    }
}

?>