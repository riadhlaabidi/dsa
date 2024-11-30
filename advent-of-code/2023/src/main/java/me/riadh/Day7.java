package me.riadh;

import me.riadh.Puzzle;

import java.util.*;

public class Day7 extends Puzzle {
    public Day7() {
        super(7);
    }

    @Override
    public int solvePartOne(List<String> input) {
        input.sort((x, y) -> {
            String[] line1 = x.split("\\s+");
            String[] line2 = y.split("\\s+");
            int score1 = getScore(line1[0]);
            int score2 = getScore(line2[0]);
            if (score1 == score2) {
                return compareHands(line1[0], line2[0]);
            }
            return score2 - score1;
        });
        int res = 0;
        for (int i = 0; i < input.size(); i++) {
            res += (i + 1) * Integer.parseInt(input.get(i).split("\\s+")[1]);
        }
        return res;
    }

    @Override
    public long solvePartTwo(List<String> input) {
        input.forEach(i -> getScoreWithJoker(i.split("\\s+")[0]));
        input.sort((x, y) -> {
            String[] line1 = x.split("\\s+");
            String[] line2 = y.split("\\s+");
            int score1 = getScoreWithJoker(line1[0]);
            int score2 = getScoreWithJoker(line2[0]);
            if (score1 == score2) {
                return compareHandsWithJoker(line1[0], line2[0]);
            }
            return score2 - score1;
        });

        int res = 0;
        for (int i = 0; i < input.size(); i++) {
            res += (i + 1) * Integer.parseInt(input.get(i).split("\\s+")[1]);
        }
        return res;

    }

    private int getScore(String hand) {

        // five of a kind
        if (hand.matches("([AKQJT2-9])\\1{4}")) {
            return 1;
        }

        Map<Character, Integer> charsMap = getCharsMap(hand);

        // four of a kind
        if (charsMap.containsValue(4)) {
            return 2;
        }

        // full house
        if (charsMap.containsValue(2) && charsMap.containsValue(3)) {
            return 3;
        }

        // three of a kind
        if (charsMap.containsValue(3) && charsMap.size() == 3) {
            return 4;
        }

        // two pair
        if (charsMap.containsValue(2) && charsMap.size() == 3) {
            return 5;
        }

        // one pair
        if (charsMap.size() == 4) {
            return 6;
        }

        // high
        if (charsMap.size() == 5) {
            return 7;
        }

        System.out.println(hand);
        throw new IllegalArgumentException("Invalid hand");
    }

    private int getScoreWithJoker(String hand) {
        if (hand.contains("J")) {
            String s = hand.replaceAll("J", "");
            Map<Character, Integer> chars = getCharsMap(s);
            if (!chars.isEmpty()) {
                char max = s.charAt(0);
                for (Map.Entry<Character, Integer> entry : chars.entrySet()) {
                    Character key = entry.getKey();
                    Integer value = entry.getValue();
                    if (value > chars.get(max)) {
                        max = key;
                    }
                }
                return getScore(s + String.valueOf(max).repeat(hand.length() - s.length()));
            }
        }
        return getScore(hand);
    }

    private int compareHands(String hand1, String hand2) {
        String chars = "AKQJT98765432";
        for (int i = 0; i < hand1.length(); i++) {
            char c1 =  hand1.charAt(i);
            char c2 =  hand2.charAt(i);
            if (c1 == c2) {
                continue;
            }
            return chars.indexOf(c2) - chars.indexOf(c1);
        }
        throw new IllegalArgumentException();
    }

    private int compareHandsWithJoker(String hand1, String hand2) {
        String chars = "AKQT98765432J";
        for (int i = 0; i < hand1.length(); i++) {
            char c1 =  hand1.charAt(i);
            char c2 =  hand2.charAt(i);
            if (c1 == c2) {
                continue;
            }
            return chars.indexOf(c2) - chars.indexOf(c1);
        }
        throw new IllegalArgumentException();
    }

    private Map<Character, Integer> getCharsMap(String hand) {
        char[] chars = hand.toCharArray();
        Map<Character, Integer> charOccurrences = new HashMap<>();
        for (char c : chars) {
            charOccurrences.put(c, charOccurrences.getOrDefault(c, 0) + 1);
        }
        return charOccurrences;
    }
}
