import java.util.*;

public class WeightedStrings {
    // Fungsi untuk menghitung bobot dan memeriksa query
    public static List<String> weightedStrings(String s, int[] queries) {
        Set<Integer> weights = new HashSet<>();
        int length = s.length();

        // Menghitung bobot semua substring dengan karakter berulang yang berurutan
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

        // Memeriksa setiap query terhadap bobot yang dihitung
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

    // Fungsi utama untuk menguji func weightedStrings
    public static void main(String[] args) {
        String data = "abbbcccd";
        int[] queries = {1, 3, 9, 8};

        List<String> results = weightedStrings(data, queries);
        System.out.println(results);
    }
}
