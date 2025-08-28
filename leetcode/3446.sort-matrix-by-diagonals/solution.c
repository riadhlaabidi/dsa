#include <assert.h>
#include <stdio.h>
#include <stdlib.h>

/**
 * Selection sort of the top-left to bottom-right diagonal starting at (row,
 * col) in decreasing 1 or non-decreasing 0 order in an n x n matrix.
 */
void sort_diagonal(int **grid, int n, int row, int col, int decreasing)
{

    for (int i = row, j = col; i < n - 1 && j < n - 1; i++, j++) {
        int keyi = i;
        int keyj = j;

        for (int k = i + 1, l = j + 1; k < n && l < n; k++, l++) {
            if ((decreasing && grid[k][l] > grid[keyi][keyj]) ||
                (!decreasing && grid[k][l] < grid[keyi][keyj])) {
                keyi = k;
                keyj = l;
            }
        }

        int tmp = grid[keyi][keyj];
        grid[keyi][keyj] = grid[i][j];
        grid[i][j] = tmp;
    }
}

void sort_matrix(int **grid, int grid_size)
{

    int i = grid_size - 2;
    while (i >= 0) {
        sort_diagonal(grid, grid_size, i, 0, 1);
        i--;
    }

    int j = 1;
    while (j < grid_size - 1) {
        sort_diagonal(grid, grid_size, 0, j, 0);
        j++;
    }
}

int main(void)
{
    int g[][3] = {
        {1, 7, 3},
        {9, 8, 2},
        {4, 5, 6},
    };

    int *grid[] = {
        g[0],
        g[1],
        g[2],
    };

    int expected[][3] = {
        {8, 2, 3},
        {9, 6, 7},
        {4, 5, 1},
    };

    int grid_size = 3;

    sort_matrix(grid, grid_size);

    for (int i = 0; i < grid_size; i++) {
        for (int j = 0; j < grid_size; j++) {
            if (grid[i][j] != expected[i][j]) {
                fprintf(stderr, "Expected grid[%d][%d] %d, but got %d\n", i, j,
                        expected[i][j], grid[i][j]);
                return EXIT_FAILURE;
            }
        }
    }

    printf("Correct");
    return EXIT_SUCCESS;
}
