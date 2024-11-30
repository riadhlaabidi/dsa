package me.riadh;

import java.io.IOException;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public abstract class Puzzle {

    private final int day;

    public Puzzle(int day) {
        this.day = day;
    }
    public final void solve() {
        List<String> input = loadInput();
        System.out.println("Part 1: " + solvePartOne(input));
        System.out.println("Part 2: " + solvePartTwo(input));
    }

    public abstract int solvePartOne(List<String> input);

    public abstract long solvePartTwo(List<String> input);

    private List<String> loadInput() {
        String filename = String.format("day%d.txt", day);
        List<String> list = new ArrayList<>();
        try (Scanner scanner = new Scanner(Paths.get("src/main/resources/input/" + filename))) {
            while (scanner.hasNextLine()) {
                list.add(scanner.nextLine());
            }
            return list;
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

}
