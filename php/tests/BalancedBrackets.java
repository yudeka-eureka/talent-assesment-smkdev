import java.util.*;

public class BalancedBrackets {

    // Function to check if the string of brackets is balanced
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

    // Helper function to check if the characters form a matching pair
    private static boolean isMatchingPair(char open, char close) {
        return (open == '(' && close == ')') ||
                (open == '{' && close == '}') ||
                (open == '[' && close == ']');
    }

    // Main function to test the isBalanced function
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