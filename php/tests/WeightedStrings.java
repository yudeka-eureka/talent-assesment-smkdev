import java.util.*;

public class WeightedStrings {
    // Function to calculate the weights and check the queries
    public static List<String> weightedStrings(String s, int[] queries) {
        Set<Integer> weights = new HashSet<>();
        int length = s.length();

        // Calculate weights of all substrings with consecutive repeated characters
        for (int i = 0; i < length; ) {
            char currentChar = s.charAt(i);
            int weight = currentChar - 'a' + 1;
            int sum = 0;
            int count = 0;

            while (i < length && s.charAt(i) == currentChar) {
                sum += weight;
                count++;
                weights.add(sum);
                i++;
            }
        }

        // Check each query against the calculated weights
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

    // Main function to test the weightedStrings function
    public static void main(String[] args) {
        String data = "abbbcccd";
        int[] queries = {1, 3, 9, 8};

        List<String> results = weightedStrings(data, queries);
        System.out.println(results); // Output: [Yes, Yes, Yes, No]
    }
}
