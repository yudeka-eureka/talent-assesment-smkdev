<?php
// tests/SumTest.php
use PHPUnit\Framework\TestCase;

class BalancedBracketTest extends TestCase {
    public function testBalancedBracket() {
        require 'BalancedBracket.php';
        $this->assertEquals("NO", balanceBracket("{ [ ( ] ) }"));
        $this->assertEquals("YES", balanceBracket("{ ( ( [ ] ) [ ] ) [ ] }"));
        $this->assertEquals("YES", balanceBracket("{[()]}"));
        $this->assertEquals("NO", balanceBracket("{(([|])[])[]}"));
    }
}

?>