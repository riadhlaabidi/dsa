package me.riadh;

import me.riadh.Puzzle;

import java.util.*;
import java.util.stream.Collectors;

public class Day4 extends Puzzle {
    public Day4() {
        super(4);
    }

    @Override
    public int solvePartOne(List<String> input) {
        int sum = 0;
        for (String line : input) {
            String[] card = line.split(":\\s")[1]
                    .replaceFirst("\\s\\|", "")
                    .split("\\s+");
            int matches = getMatches(card);
            if (matches > 0) {
                sum += (int) Math.pow(2, matches - 1);
            }
        }
        return sum;
    }

    @Override
    public long solvePartTwo(List<String> input) {
        int sum = 0;
        Map<Integer, Integer> cardCopies = new HashMap<>();
        for (String line : input) {
            String[] card = line.split(":\\s");
            int cardNumber = Integer.parseInt(card[0].split("\\s+")[1]);
            String[] numbers = card[1]
                    .replaceFirst("\\s\\|", "")
                    .split("\\s+");
            int matches = getMatches(numbers);
            int currentCardCopies = cardCopies.getOrDefault(cardNumber, 1);
            sum += currentCardCopies;
            for (int i = 1; i <= matches; i++) {
                int copiesToAdd = cardCopies.getOrDefault(cardNumber + i, 1) + currentCardCopies;
                cardCopies.put(cardNumber + i, copiesToAdd);
            }
        }
        return sum;
    }

    private int getMatches(String[] numbers) {
        Set<String> numbersSet = Arrays.stream(numbers).collect(Collectors.toSet());
        return numbers.length - numbersSet.size();
    }
}
