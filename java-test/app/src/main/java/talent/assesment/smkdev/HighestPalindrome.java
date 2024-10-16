package talent.assesment.smkdev;
import java.util.*;

public class HighestPalindrome {

    // Fungsi rekursif untuk memaksimalkan palindrom
    public static String highestPalindrome(String s, int k) {
        return new String("00000");
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