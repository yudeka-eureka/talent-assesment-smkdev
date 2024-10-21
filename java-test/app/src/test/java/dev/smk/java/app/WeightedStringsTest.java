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
        /*
        As per my understanding, assert of result1 should be: Yes, Yes, No, No, No, No, since the weights of the
        characters in the string "abccba" are 1, 2, 3, 3, 2, 1. If the assert need to be true then the query is
        1, 3, 6, 5, 9, 10. Since 'c' has a weight of 3 and previous character is 'c' then the weight of the string
        is 3 + 3 = 6.
        */
        List<String> result1 = WeightedStrings.weightedStrings("abccba", new int[]{1, 3, 12, 5, 9, 10});
        assertEquals(Arrays.asList("Yes", "Yes", "Yes", "No", "No", "No"), result1);

        List<String> result2 = WeightedStrings.weightedStrings("abc", new int[]{1, 2, 3, 4});
        assertEquals(Arrays.asList("Yes", "Yes", "Yes", "No"), result2);

        List<String> result3 = WeightedStrings.weightedStrings("zzz", new int[]{26, 52, 78});
        assertEquals(Arrays.asList("Yes", "Yes", "Yes"), result3);
    }
}
