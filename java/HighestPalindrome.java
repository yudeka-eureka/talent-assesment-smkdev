import java.util.*;

public class HighestPalindrome {

    // Fungsi rekursif untuk memaksimalkan palindrom
    public static String highestPalindrome(String s, int k) {
        int n = s.length();
        char[] chars = s.toCharArray();
        boolean[] changed = new boolean[n];

        // Langkah 1: Jadikan string sebagai palindrom
        int left = 0, right = n - 1;
        while (left < right) {
            if (chars[left] != chars[right]) {
                chars[left] = chars[right] = (char) Math.max(chars[left], chars[right]);
                changed[left] = changed[right] = true;
                k--;
            }
            left++;
            right--;
        }

        // Jika kita menggunakan lebih banyak perubahan daripada yang diizinkan, kembalikan "-1"
        if (k < 0) return "-1";

        // Langkah 2: Maksimalkan palindrom dengan menggunakan perubahan yang tersisa
        left = 0;
        right = n - 1;
        while (left <= right) {
            if (left == right) { // Menangani karakter tengah dalam string dengan panjang ganjil
                if (k > 0) {
                    chars[left] = '9';
                }
            } else {
                if (chars[left] < '9') {
                    if (changed[left] || changed[right]) {
                        if (k > 0) {
                            chars[left] = chars[right] = '9';
                            k--;
                        }
                    } else if (k > 1) {
                        chars[left] = chars[right] = '9';
                        k -= 2;
                    }
                }
            }
            left++;
            right--;
        }

        return new String(chars);
    }

    public static void main(String[] args) {
        // Contoh 1
        String s1 = "3943";
        int k1 = 1;
        System.out.println(highestPalindrome(s1, k1));

        // Contoh 2
        String s2 = "932239";
        int k2 = 2;
        System.out.println(highestPalindrome(s2, k2));

        // Contoh 3
        String s3 = "12321";
        int k3 = 1;
        System.out.println(highestPalindrome(s3, k3));

        // Contoh 4
        String s4 = "12345";
        int k4 = 1;
        System.out.println(highestPalindrome(s4, k4));
    }
}