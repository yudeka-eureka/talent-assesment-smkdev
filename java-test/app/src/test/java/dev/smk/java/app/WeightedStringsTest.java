package dev.smk.java.app;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import java.util.Arrays;
import java.util.List;

public class WeightedStringsTest {

    @Test
    void testGetMessage() {
        assertEquals("Hello      World!", MessageUtils.getMessage());
    }

    @Test
    void testWeightedStrings() {
        List<String> result1 = WeightedStrings.weightedStrings("abccba", new int[]{1, 3, 12, 5, 9, 10});
        assertEquals(Arrays.asList("Yes", "Yes", "Yes", "No", "No", "No"), result1);

        List<String> result2 = WeightedStrings.weightedStrings("abc", new int[]{1, 2, 3, 4});
        assertEquals(Arrays.asList("Yes", "Yes", "Yes", "No"), result2);

        List<String> result3 = WeightedStrings.weightedStrings("zzz", new int[]{26, 52, 78});
        assertEquals(Arrays.asList("Yes", "Yes", "Yes"), result3);
    }
}
