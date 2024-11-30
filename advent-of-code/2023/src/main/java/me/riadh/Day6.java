package me.riadh;

import me.riadh.Puzzle;

import java.util.Arrays;
import java.util.List;

public class Day6 extends Puzzle {
    public Day6() {
        super(6);
    }

    @Override
    public int solvePartOne(List<String> input) {
        List<Integer> time = Arrays
                .stream(input.getFirst().split(":\\s+")[1].trim().split("\\s+"))
                .map(Integer::parseInt)
                .toList();
        List<Integer> distance =Arrays
                .stream(input.get(1).split(":\\s+")[1].trim().split("\\s+"))
                .map(Integer::parseInt)
                .toList();

        int errorMargin = 1;
        for (int i = 0; i < time.size(); i++) {
            int waysToWin = 0;
            for (int j = 1; j < time.get(i); j++) {
                if ((time.get(i) - j) * j > distance.get(i)) {
                    waysToWin++;
                }
            }
            errorMargin *= waysToWin;
        }
        return errorMargin;
    }

    @Override
    public long solvePartTwo(List<String> input) {
        long time = Long.parseLong(input.getFirst().split(":\\s+")[1].replaceAll("\\s+", ""));
        long distance = Long.parseLong(input.get(1).split(":\\s+")[1].replaceAll("\\s+", ""));
        int waysToWin = 0;
        for (int i = 1; i < time; i++) {
            if ((time - i) * i > distance) {
                waysToWin++;
            }
        }
        return waysToWin;
    }
}
