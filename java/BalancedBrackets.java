import java.util.*;

public class BalancedBrackets {

    // Fungsi untuk memeriksa apakah rangkaian tanda kurung seimbang
    public static String isBalanced(String s) {
        Stack<Character> stack = new Stack<>();
        for (char ch : s.toCharArray()) {
            if (ch == '(' || ch == '{' || ch == '[') {
                stack.push(ch);
            } else if (ch == ')' || ch == '}' || ch == ']') {
                if (stack.isEmpty()) {
                    return "NO";
                }
                char top = stack.pop();
                if (!isMatchingPair(top, ch)) {
                    return "NO";
                }
            }
        }
        return stack.isEmpty() ? "YES" : "NO";
    }

    // Fungsi pembantu untuk memeriksa apakah karakter membentuk pasangan yang cocok
    private static boolean isMatchingPair(char open, char close) {
        return (open == '(' && close == ')') ||
                (open == '{' && close == '}') ||
                (open == '[' && close == ']');
    }

    // Fungsi utama untuk menguji func isBalanced
    public static void main(String[] args) {
        String[] testCases = {
                "{ ( [ ] ) }",
                "{ [ ( ] ) }",
                "{ ( ( [ ] ) [ ] ) [ ] }"
        };

        for (String test : testCases) {
            System.out.println(isBalanced(test));
        }
    }
}