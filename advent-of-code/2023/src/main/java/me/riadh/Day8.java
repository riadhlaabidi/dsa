package me.riadh;

import me.riadh.Puzzle;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Day8 extends Puzzle {
    public Day8() {
        super(8);
    }

    @Override
    public int solvePartOne(List<String> input) {
        String instructions = input.getFirst();
        Map<String, Node> nodes = getNetwork(input);
        String currentPosition = "AAA";
        int i = 0;
        while (!currentPosition.equals("ZZZ")) {
            Node currentNode = nodes.get(currentPosition);
            currentPosition = instructions.charAt(i % instructions.length()) == 'L'
                    ? currentNode.left()
                    : currentNode.right();
            i++;
        }
        return i;
    }

    @Override
    public long solvePartTwo(List<String> input) {
        String instructions = input.getFirst();
        Map<String, Node> nodes = getNetwork(input);
        List<String> startingNodes = new ArrayList<>();
        nodes.forEach((key, value) -> {
            if (key.endsWith("A")) {
                startingNodes.add(key);
            }
        });

        long res = 1;
        for (String node : startingNodes) {
            int i = 0;
            String currentNode = node;
            while(!currentNode.endsWith("Z")) {
                currentNode = instructions.charAt(i % instructions.length()) == 'L'
                        ? nodes.get(currentNode).left
                        : nodes.get(currentNode).right;
                i++;
            }
            res = res / gcd(res, i) * i;
        }
        return res;
    }

    private Map<String, Node> getNetwork(List<String> input) {
        Map<String, Node> nodes = new HashMap<>();
        for (int i = 2; i < input.size(); i++) {
            String line = input.get(i);
            String[] parts = line.split("\\s=\\s");
            String[] lr = parts[1].replaceAll("[()]", "").split(",\\s");
            Node node = new Node(lr[0], lr[1]);
            nodes.put(parts[0], node);
        }
        return nodes;
    }

    private long gcd(long a, long b) {
        if (b == 0L) {
            return a;
        }
        return gcd(b, a % b);
    }
    private record Node(String left, String right) { }

}
