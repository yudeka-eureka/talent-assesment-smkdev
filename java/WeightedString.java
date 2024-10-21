class WeightedString {
    public static void main(String[] args) {
        String data = "abccddde";
        int[] queries = {6, 1, 3, 12, 5, 9, 10};

        List<String> results = weightedStrings(data, queries);
        System.out.println(results);
    }

    public static List<String> weightedUniformStrings(String s, List<Integer> queries) {
        Set<Integer> weights = new HashSet<>();
        char lastChar = '0';
        Integer lastCharWeight = 0;
        Integer currentTotalWeight = 0;

        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == lastChar) {
                currentTotalWeight = currentTotalWeight + lastCharWeight;
            } else {
                lastChar = s.charAt(i);
                lastCharWeight = s.charAt(i) - 'a' + 1; // 'a' is equals with 97 in ASCII code
                currentTotalWeight = lastCharWeight;
            }
            weights.add(currentTotalWeight);
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