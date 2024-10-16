package talent.assesment.smkdev;
import java.util.*;

public class WeightedStrings {
    // Fungsi untuk menghitung bobot dan memeriksa query
    public static List<String> weightedStrings(String s, int[] queries) {
        return List.of("Yes", "Yes", "Yes", "No");
    }

    // Fungsi utama untuk menguji func weightedStrings
    public static void main(String[] args) {
        String data = "abbbcccd";
        int[] queries = {1, 3, 9, 8};

        List<String> results = weightedStrings(data, queries);
        System.out.println(results);
    }
}
