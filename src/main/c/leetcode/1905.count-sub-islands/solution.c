#include <stdio.h>
#include <stdlib.h>

int is_sub_island(int row, int col, int **grid1, int **grid2, int n, int m, int **visited) 
{
    if (row < 0 || row >= n || col < 0 || col >= m
        || visited[row][col] || grid2[row][col] == 0) {
        return 1;
    }

    if (grid1[row][col] != 1) {
        return 0;
    }
    
    visited[row][col] = 1;
    int dirs[4][2] = {{0, 1}, {1, 0}, {-1, 0}, {0, -1}};
    int valid = 1;

    for (int i = 0; i < 4; ++i) {
        int next_is_sub = is_sub_island(row + dirs[i][0], col + dirs[i][1], 
                                        grid1, grid2, n, m, visited);
        valid = valid && next_is_sub;
    }

    return valid;
}

int countSubIslands(int **grid1, int grid1_size, int *grid1_col_size, 
		    int **grid2, int grid2_size, int *grid2_col_size) 
{

    int **visited = (int**) malloc(grid1_size * sizeof(int*));

    if (!visited) {
        fprintf(stderr, "Memory allocation failed");
        exit(EXIT_FAILURE);
    }

    for (int i = 0; i < grid1_size; ++i) {
        visited[i] = (int*) calloc(grid1_col_size[i], sizeof(int));

        if (!visited[i]) {
            fprintf(stderr, "Memory allocation failed");
            free(visited);
            exit(EXIT_FAILURE);
        }
    }

    int sub_count = 0;
    for (int i = 0; i < grid1_size; ++i) {
        for (int j = 0; j < grid1_col_size[i]; ++j) {
            if (!visited[i][j] 
                && grid2[i][j] == 1
                && is_sub_island(i, j, grid1, grid2, grid1_size,
                                 grid1_col_size[i], visited)
            ) {
                sub_count++;
            }
        }
    }

    free(visited);
    return sub_count;
}
