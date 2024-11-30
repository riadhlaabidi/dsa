package me.riadh;

import me.riadh.Puzzle;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Day9 extends Puzzle {
    public Day9() {
        super(9);
    }

    @Override
    public int solvePartOne(List<String> input) {
        int sum = 0;
        for (String line : input) {
            List<Integer> numbers = Arrays.stream(line.split("\\s"))
                    .map(Integer::parseInt)
                    .toList();

            while (true) {
                List<Integer> newList = new ArrayList<>();
                boolean allZeroes = true;
                sum += numbers.getLast();

                for (int j = 0; j < numbers.size() - 1; j++) {
                    int value = numbers.get(j + 1) - numbers.get(j);
                    newList.add(value);
                    if (value != 0) {
                        allZeroes = false;
                    }
                }

                if (allZeroes) {
                    break;
                }

                numbers = newList;
            }
        }
        return sum;
    }

    @Override
    public long solvePartTwo(List<String> input) {
        int sum = 0;
        for (String line : input) {
            List<Integer> numbers = Arrays.stream(line.split("\\s"))
                    .map(Integer::parseInt)
                    .toList();

            int sequences = 0;

            while (true) {
                List<Integer> newList = new ArrayList<>();
                boolean allZeroes = true;

                sum = sequences % 2 == 0
                        ? sum + numbers.getFirst()
                        : sum - numbers.getFirst();

                for (int j = 0; j < numbers.size() - 1; j++) {
                    int value = numbers.get(j + 1) - numbers.get(j);
                    newList.add(value);
                    if (value != 0) {
                        allZeroes = false;
                    }
                }

                if (allZeroes) {
                    break;
                }

                sequences++;
                numbers = newList;
            }
        }
        return sum;
    }
}
