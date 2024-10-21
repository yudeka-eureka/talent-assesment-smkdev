import java.util.ArrayList;
import java.util.Arrays;
import java.util.stream.Collectors;

class HighestPalindrome {
    public static ArrayList<Integer> temp = new ArrayList<Integer>();
    public static Integer highestNumber = 0;

    public static void main(String[] args) {
        String s = "3943";
        int k = 2;
        highestPalindrome(s, k);
    }

    public static boolean isPalindrome(String s) {
        if (s.isEmpty() || s.length() == 1)
            // if length =0 OR 1 then it is
            return true;
        if (s.charAt(0) == s.charAt(s.length() - 1))
            // check for first and last char of String:
            // if they are same then do the same thing for a substring
            // with first and last char removed. and carry on this
            // until you string completes or condition fails
            return isPalindrome(s.substring(1, s.length() - 1));

        // if its not the case than string is not.
        return false;
    }

    public static void highestPalindrome(String s, int k) {
        if (s.isEmpty() || k < 1) {
            System.out.println(highestNumber.toString());
        } else {
            char firstCharacter = s.charAt(0);
            char lastCharacter = s.charAt(s.length() - 1);

            if (firstCharacter != lastCharacter) {
                Integer biggest = Math.max(
                        Integer.parseInt(String.valueOf(firstCharacter)),
                        Integer.parseInt(String.valueOf(lastCharacter))
                );

                temp.add((temp.size() / 2), biggest);
                temp.add((temp.size() / 2) + 1, biggest);
                k--;

                String formattedPalindrome = temp.stream().map(Object::toString).collect(Collectors.joining(""));
                if (isPalindrome(formattedPalindrome)) {
                    highestNumber = Math.max(highestNumber, Integer.parseInt(formattedPalindrome));
                    highestPalindrome(s.substring(1, s.length() - 1), k);
                    return;
                }

                highestPalindrome(s.substring(1, s.length() - 1), k);
            } else {
                temp.add((temp.size() / 2), Integer.parseInt(String.valueOf(firstCharacter)));
                temp.add((temp.size() / 2) + 1, Integer.parseInt(String.valueOf(lastCharacter)));

                if (isPalindrome(s)) {
                    k--;
                    highestNumber = Math.max(highestNumber, Integer.parseInt(s));
                    highestPalindrome(s.substring(1, s.length() - 1), k);
                    return;
                }

                highestPalindrome(s.substring(1, s.length() - 1), k);
            }
        }
    }
}