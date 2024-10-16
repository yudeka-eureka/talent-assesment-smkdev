package talent.assesment.smkdev;

public class BalancedBrackets {

    // Fungsi untuk memeriksa apakah rangkaian tanda kurung seimbang
    public static String isBalanced(String s) {
        return "YES";
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