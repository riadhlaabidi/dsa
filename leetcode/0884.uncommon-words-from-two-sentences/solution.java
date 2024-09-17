import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

class Solution {
    public static String[] uncommonFromSentences(String s1, String s2) {
        Map<String, Integer> count = new HashMap<>();

        String[] s1Words = s1.split("\\s+");
        String[] s2Words = s2.split("\\s+");

        for (int i = 0; i < s1Words.length; ++i) {
            count.put(s1Words[i], count.getOrDefault(s1Words[i], 0) + 1);
        }

        for (int i = 0; i < s2Words.length; ++i) {
            count.put(s2Words[i], count.getOrDefault(s2Words[i], 0) + 1);
        }

        List<String> ans = new ArrayList<>();
        for (var en : count.entrySet()) {
            if (en.getValue() == 1) {
                ans.add(en.getKey());
            }
        }

        return ans.toArray(new String[0]);
    }

    public static void main(String[] args) {
        String s1 = "this apple is sweet";
        String s2 = "this apple is sour";

        String[] expected = new String[] { "sweet", "sour" };
        String[] actual = uncommonFromSentences(s1, s2);

        assert (actual == expected);

        System.out.println("Correct " + Arrays.toString(actual));
    }
}
