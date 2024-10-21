package dev.smk.java.app;
import java.util.*;
public class WeightedStrings {

    public static List<String> weightedStrings(String s, int[] queries) {
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
