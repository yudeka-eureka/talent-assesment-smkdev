package talent.assesment.smkdev;

import org.junit.Test;
import static org.junit.Assert.*;


public class BalancedBracketsTest {

    @Test
    public void testIsBalancedCase1() {
        assertEquals("YES", BalancedBrackets.isBalanced("{ ( [ ] ) }"));
    }

    @Test
    public void testIsBalancedCase2() {
        assertEquals("NO", BalancedBrackets.isBalanced("{ [ ( ] ) }"));
    }

    @Test
    public void testIsBalancedCase3() {
        assertEquals("YES", BalancedBrackets.isBalanced("{ ( ( [ ] ) [ ] ) [ ] }"));
    }

    @Test
    public void testIsBalancedCase4() {
        assertEquals("NO", BalancedBrackets.isBalanced("{{[[(())]]}}"));
    }

    @Test
    public void testIsBalancedCase5() {
        assertEquals("NO", BalancedBrackets.isBalanced("{{[[(())]]}"));
    }
}
