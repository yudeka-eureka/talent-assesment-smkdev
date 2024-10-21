class BalanceBracket {
    public static void main(String[] args) {
        String[] samples = {"{ ( [ ] ) }", "{ [ ( ] ) }", "{ ( ( [ ] ) [ ] ) [ ] }"};

        for (String sample : samples) {
            System.out.println(isBracketPaired(sample));
        }
    }

    public static String isBracketPaired() {
        Stack<Character> pairedStack = new Stack<>();
        Map<Character, Character> charMap = new HashMap<>();
        charMap.put('}', '{');
        charMap.put(']', '[');
        charMap.put(')', '(');

        for (char bracket : s.toCharArray()) {
            if (charMap.containsValue(bracket)) {
                pairedStack.push(bracket);
            } else if (charMap.containsKey(bracket)) {
                if (pairedStack.isEmpty() || pairedStack.pop() != charMap.get(bracket)) {
                    return "NO";
                }
            }
        }

        return pairedStack.isEmpty() ? "YES" : "NO";
    }
}