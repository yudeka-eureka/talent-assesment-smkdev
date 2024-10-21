package dev.smk.java.app;
import java.util.*;
public class WeightedStrings {

    public static List<String> weightedStrings(String s, int[] queries) {
        Set<Integer> weights = new HashSet<>();
        char lastChar = '0';
        int currentWeightSum = 0;

        for (char currentChar : s.toCharArray()) {
            /*
            Calculate the weight of the current character, since in ASCII, 'a' (the first char) is 97, we need to do
            this to get the correct weight that starts from 1.
            */
            int charWeight = currentChar - 'a' + 1;

            if (currentChar == lastChar) {
                currentWeightSum += charWeight;
            } else {
                currentWeightSum = charWeight;
            }
            weights.add(currentWeightSum);
            lastChar = currentChar;
        }

        List<String> results = new ArrayList<>();
        for (int query : queries) {
            if (weights.contains(query)) {
                results.add("Yes");
            } else {
                results.add("No");
            }
        }

        return results;
    }
}