package me.riadh;

import me.riadh.Puzzle;

import java.util.List;

public class Day3 extends Puzzle {
    public Day3() {
        super(3);
    }

    @Override
    public int solvePartOne(List<String> input) {
        int sum = 0;
        for (int i = 0; i < input.size(); i++) {
            int start = -1;
            boolean isAdjacent = false;
            String line = input.get(i);
            for (int j = 0; j < line.length(); j++) {
                if (!Character.isDigit(line.charAt(j))) {
                    if (start >= 0) {
                        if (isAdjacent) {
                            sum += Integer.parseInt(line.substring(start, j));
                        }
                        start = -1;
                        isAdjacent = false;
                    }
                }
                if (Character.isDigit(line.charAt(j))) {
                    if (start < 0) {
                        start = j;
                    }
                    if (isAdjacent(i, j, input)) {
                        isAdjacent = true;
                    }
                }
            }
            if (start > 0 && isAdjacent) {
                sum += Integer.parseInt(line.substring(start));
            }
        }
        return sum;
    }

    private boolean isAdjacent(int x, int y, List<String> input) {
        for (int i = x - 1; i <= x + 1; i++) {
            for (int j = y - 1; j <= y + 1; j++) {
                if (i >= 0 && i <= input.size() - 1 && j >= 0 && j <= input.getFirst().length() - 1) {
                    char c = input.get(i).charAt(j);
                    if (!Character.isDigit(c) && c != '.') {
                        return true;
                    }
                }
            }
        }
        return false;
    }

    @Override
    public long solvePartTwo(List<String> input) {
        return 0;
    }
}
