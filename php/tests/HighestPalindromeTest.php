<?php
// tests/SumTest.php
use PHPUnit\Framework\TestCase;

class HighestPalindromeTest extends TestCase {
    public function testHighestPalindrome() {
        require 'HighestPalindrome.php';
        $this->assertEquals("3993", highestPalindrome("3943",1));
        $this->assertEquals("992299", highestPalindrome("932239",2));
    }
}

?>