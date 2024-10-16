package talent.assesment.smkdev;

import org.junit.Test;
import static org.junit.Assert.*;
import java.util.List;

public class WeightedStringsTest {

    @Test
    public void testWeightedStringsCase1() {
        List<String> expected = List.of("Yes", "Yes", "Yes", "No");
        assertEquals(expected, WeightedStrings.weightedStrings("abbcccd", new int[]{1, 3, 9, 8}));
    }

    @Test
    public void testWeightedStringsCase2() {
        List<String> expected = List.of("No", "Yes", "No", "Yes", "Yes", "No");
        assertEquals(expected, WeightedStrings.weightedStrings("aaabbbcccddd", new int[]{5, 9, 7, 8, 12, 5}));
    }

    @Test
    public void testWeightedStringsCase3() {
        List<String> expected = List.of("No", "No", "No", "Yes", "Yes", "Yes", "No");
        assertEquals(expected, WeightedStrings.weightedStrings("getitssimple", new int[]{1, 2, 8, 7, 12, 38, 10}));
    }
}
