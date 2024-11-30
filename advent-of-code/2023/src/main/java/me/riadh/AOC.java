package me.riadh;

import me.riadh.Puzzle;

public class AOC {
    public static Puzzle getPuzzle(int day) {
        Puzzle puzzle;
        switch (day) {
            case 1 -> puzzle = new Day1();
            case 2 -> puzzle = new Day2();
            case 3 -> puzzle = new Day3();
            case 4 -> puzzle = new Day4();
            case 5 -> puzzle = new Day5();
            case 6 -> puzzle = new Day6();
            case 7 -> puzzle = new Day7();
            case 8 -> puzzle = new Day8();
            case 9 -> puzzle = new Day9();
//            case 10 -> puzzle = new Day10();
//            case 11 -> puzzle = new Day11();
//            case 12 -> puzzle = new Day12();
//            case 13 -> puzzle = new Day13();
//            case 14 -> puzzle = new Day14();
//            case 15 -> puzzle = new Day15();
//            case 16 -> puzzle = new Day16();
//            case 17 -> puzzle = new Day17();
//            case 18 -> puzzle = new Day18();
//            case 19 -> puzzle = new Day19();
//            case 20 -> puzzle = new Day20();
//            case 21 -> puzzle = new Day21();
//            case 22 -> puzzle = new Day22();
//            case 23 -> puzzle = new Day23();
//            case 24 -> puzzle = new Day24();
//            case 25 -> puzzle = new Day25();
            default -> throw new IllegalArgumentException("Invalid day");
        }
        return puzzle;
    }
}
