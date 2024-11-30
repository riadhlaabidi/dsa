package me.riadh;

import me.riadh.Puzzle;

import java.util.List;

public class Day1 extends Puzzle {

    public Day1() {
        super(1);
    }

    @Override
    public int solvePartOne(List<String> input) {
        int sum = 0;
        for (String line : input) {
            sum += extractNumber(line);
        }
        return sum;
    }

    @Override
    public long solvePartTwo(List<String> input) {
        int sum = 0;
        for (String line : input) {
            sum += extractNumber(replaceFirstAndLast(line));
        }
        return sum;
    }

    private int extractNumber(String input) {
        int i = 0;
        int j = input.length() - 1;
        while (!Character.isDigit(input.charAt(i))) {
            i++;
        }
        while (!Character.isDigit(input.charAt(j))) {
            j--;
        }
        return Integer.parseInt(String.format("%c%c", input.charAt(i), input.charAt(j)));
    }

    private String replaceFirstAndLast(String input) {
        String[] numbers = {"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"};
        int firstToReplace = -1;
        int lastToReplace = -1;
        int firstIndex = input.length();
        int lastIndex = 0;

        for (int i = 0; i < numbers.length; i++) {
            int first = input.indexOf(numbers[i]);
            if (first < firstIndex && first != -1) {
                firstIndex = first;
                firstToReplace = i;
            }
            int last = input.lastIndexOf(numbers[i]);
            if (last > lastIndex) {
                lastIndex = last;
                lastToReplace = i;
            }
        }

        if (firstToReplace != -1) {
            if (firstIndex + numbers[firstToReplace].length() - 1 == lastIndex) {
                return input.substring(0, firstIndex)
                        .concat(String.valueOf(firstToReplace + 1))
                        .concat(String.valueOf(lastToReplace + 1))
                        .concat(input.substring(lastIndex + numbers[lastToReplace].length()));
            }
            input = input.replaceFirst(numbers[firstToReplace], String.valueOf(firstToReplace + 1));
        }
        if (lastToReplace != -1) {
            input = replaceLast(input, numbers[lastToReplace], String.valueOf(lastToReplace + 1));
        }
        return input;
    }

    public String replaceLast(String input, String toReplace, String replacement) {
        int pos = input.lastIndexOf(toReplace);
        if (pos == -1) {
            return input;
        }
        return input.substring(0, pos)
                .concat(replacement)
                .concat(input.substring(pos + toReplace.length()));
    }

}
