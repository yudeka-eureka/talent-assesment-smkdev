package dev.smk.java.app;

public class HighestPalindrome {

    public static String highestPalindrome(String s, int k) {
        if (k == 0) return s;

        char[] chars = s.toCharArray();
        int n = s.length();
        boolean[] changed = new boolean[n];
        int left = 0, right = n - 1;

        // make the string a palindrome with existing char
        while (left < right) {
            if (chars[left] != chars[right]) {
                chars[left] = chars[right] = (char) Math.max(chars[left], chars[right]);
                changed[left] = changed[right] = true;
                k--;
            }
            left++;
            right--;
        }

        // not enough changes to make it a palindrome
        if (k < 0) return "-1";

        // maximize the palindrome by changing pairs to '9' where possible
        left = 0;
        right = n - 1;
        while (left <= right) {
            // if at the middle of an odd-length string
            if (left == right && k > 0) {
                chars[left] = '9';
            } else if (chars[left] != '9') {
                // if haven't already changed it and have changes left
                if (k >= 2 && !changed[left] && !changed[right]) {
                    chars[left] = chars[right] = '9';
                    k -= 2; // use two changes
                } else if (k >= 1 && (changed[left] || changed[right])) {
                    chars[left] = chars[right] = '9';
                    k--; // use one change if we've already changed this pair to match
                }
            }
            left++;
            right--;
        }

        return new String(chars);
    }
}