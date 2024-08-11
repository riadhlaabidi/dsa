package me.riadh.algorithms.graph;

import java.util.List;
import java.util.Map;

public final class GraphAlgorithms {

    private GraphAlgorithms() {
    }

    public void dfs(Map<Integer, List<Integer>> adjacencyList, boolean[] visited, int vertex) {

        if (visited[vertex]) {
            return;
        }

        visited[vertex] = true;

        for (int neighbour : adjacencyList.get(vertex)) {
            dfs(adjacencyList, visited, neighbour);
        }
    }

}
