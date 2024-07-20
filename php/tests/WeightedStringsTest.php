<?php
// tests/SumTest.php
use PHPUnit\Framework\TestCase;

class WeightedStringsTest extends TestCase {
    public function testWeightedStrings() {
        require  'WeightedStrings.php';  
        $this->assertEquals(["YES", "YES", "YES", "NO"], weightedStrings("abbcccd", [1, 3, 9, 8]));
        $this->assertEquals(["NO", "YES", "NO", "YES", "YES", "NO"], weightedStrings("aaabbbcccddd", [5, 9, 7, 8, 12, 5]));
    }
}

?>