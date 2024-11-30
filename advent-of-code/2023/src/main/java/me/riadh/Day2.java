package me.riadh;

import me.riadh.Puzzle;

import java.util.Arrays;
import java.util.List;

public class Day2 extends Puzzle {

    public Day2() {
        super(2);
    }

    @Override
    public int solvePartOne(List<String> input) {
        int sum = 0;
        for (String line : input) {
            String[] game = line.split(":\\s");
            int gameId = Integer.parseInt(game[0].split("\\s")[1]);
            if (isPossible(game[1])) {
                sum += gameId;
            }
        }
        return sum;
    }

    @Override
    public long solvePartTwo(List<String> input) {
        int sum = 0;
        for (String line : input) {
            String[] game = line.split(":\\s");
            int minRed = 0;
            int minGreen = 0;
            int minBlue = 0;
            for (String set : game[1].split(";\\s")) {
                int[] numberOfCubes = getNumberOfCubes(set);
                minRed = Math.max(minRed, numberOfCubes[0]);
                minGreen = Math.max(minGreen, numberOfCubes[1]);
                minBlue = Math.max(minBlue, numberOfCubes[2]);
            }
            int product = minRed * minGreen * minBlue;
            sum += product;
        }
        return sum;
    }

    private int[] getNumberOfCubes(String set) {
        int[] result = new int[3];
        String[] setCubes = set.split(",\\s");
        Arrays.stream(setCubes).forEach(cube -> {
            String[] line = cube.split("\\s");
            int number = Integer.parseInt(line[0]);
            switch (line[1]) {
                case "red" -> result[0] = number;
                case "green" -> result[1] = number;
                case "blue" -> result[2] = number;
            }
        });
        return result;
    }

    private boolean isPossible(String gameSets) {
        String[] sets = gameSets.split(";\\s");
        for (String set : sets) {
            int[] cubes = getNumberOfCubes(set);
            if (cubes[0] > 12 || cubes[1] > 13 || cubes[2] > 14) {
                return false;
            }
        }
        return true;
    }
}
