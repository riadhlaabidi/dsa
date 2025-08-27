#include <stdio.h>
#include <string.h>

#define MAX_LENGTH 500

typedef enum {
    TOP_LEFT = 0,
    TOP_RIGHT,
    BOTTOM_RIGHT,
    BOTTOM_LEFT,
    DIRECTIONS_COUNT,
} Diagonal;

static const int directions[DIRECTIONS_COUNT][2] = {
    [TOP_LEFT] = {-1, -1},
    [TOP_RIGHT] = {-1, 1},
    [BOTTOM_RIGHT] = {1, 1},
    [BOTTOM_LEFT] = {1, -1},
};

// max diagonal length state per direction for each case
// (with making a turn 0 or not 1 (size 2)), for each target (0 or 2 (size 3))
typedef int max_per_direction[DIRECTIONS_COUNT][2][3];

static max_per_direction memo[MAX_LENGTH][MAX_LENGTH];

int dfs(int **grid, int rows, int cols, int row, int col, Diagonal direction,
        int made_turn, int target)
{

    int new_row = row + directions[direction][0];
    int new_col = col + directions[direction][1];

    if (new_row < 0 || new_row >= rows || new_col < 0 || new_col >= cols ||
        grid[new_row][new_col] != target) {
        return 0;
    }

    if (memo[new_row][new_col][direction][made_turn][2 - target] != -1) {
        return memo[new_row][new_col][direction][made_turn][2 - target];
    }

    int max_from_current = dfs(grid, rows, cols, new_row, new_col, direction,
                               made_turn, 2 - target);
    if (!made_turn) {
        int max_if_turned = dfs(grid, rows, cols, new_row, new_col,
                                (direction + 1) % DIRECTIONS_COUNT, 1,
                                2 - target);
        if (max_if_turned > max_from_current) {
            max_from_current = max_if_turned;
        }
    }

    memo[new_row][new_col][direction][made_turn][2 - target] =
        max_from_current + 1;

    return max_from_current + 1;
}

int length_of_v_diagonal(int **grid, int m, int *n)
{
    memset(memo, -1, sizeof(memo));
    int longest_diag_length = 0;

    for (int i = 0; i < m; i++) {
        for (int j = 0; j < *n; j++) {
            if (grid[i][j] == 1) {
                for (int dir = 0; dir < DIRECTIONS_COUNT; dir++) {
                    int current = dfs(grid, m, *n, i, j, dir, 0, 2) + 1;
                    if (current > longest_diag_length) {
                        longest_diag_length = current;
                    }
                }
            }
        }
    }

    return longest_diag_length;
}

int main(void)
{
    int grid[5][5] = {{2, 2, 1, 2, 2},
                      {2, 0, 2, 2, 0},
                      {2, 0, 1, 1, 0},
                      {1, 0, 2, 2, 2},
                      {2, 0, 0, 2, 2}};

    int *grid_p[] = {grid[0], grid[1], grid[2], grid[3], grid[4]};
    int grid_size = 5;

    int expected = 5;
    int actual = length_of_v_diagonal(grid_p, grid_size, &grid_size);

    if (actual != expected) {
        fprintf(stderr, "Expected %d, but got %d\n", expected, actual);
        return 1;
    }

    printf("Correct\n");
    return 0;
}
