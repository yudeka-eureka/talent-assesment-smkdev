<?php
// tests/SumTest.php
use PHPUnit\Framework\TestCase;

class BalancedBracketTest extends TestCase {
    public function testBalancedBracket() {
        require 'BalancedBracket.php';
        try {
            $this->assertEquals("NO", balanceBracket("{ [ ( ] ) }"));
        } catch (Exception $e) {
            $this->fail('An error occurred during the test: ' . $e->getMessage());
        }
        try {
            $this->assertEquals("YES", balanceBracket("{ ( ( [ ] ) [ ] ) [ ] }"));
        } catch (Exception $e) {
            $this->fail('An error occurred during the test: ' . $e->getMessage());
        }
        try {
            $this->assertEquals("YES", balanceBracket("{[()]}"));
        } catch (Exception $e) {
            $this->fail('An error occurred during the test: ' . $e->getMessage());
        }
        try {
            $this->assertEquals("NO", balanceBracket("{(([|])[])[]}"));
        } catch (Exception $e) {
            $this->fail('An error occurred during the test: ' . $e->getMessage());
        }
    }
}

?>