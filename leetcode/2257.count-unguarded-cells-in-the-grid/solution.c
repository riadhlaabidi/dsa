#include <stdio.h>
#include <stdlib.h>
#include <string.h>

enum { GUARD = 1, GUARDED, WALL };

int count_unguarded(int m, int n, int **guards, int guards_size,
                    int *guards_col_size, int **walls, int walls_size,
                    int *walls_col_size)
{
    int **grid = (int **)malloc(m * sizeof(int *));

    if (!grid) {
        goto error;
    }

    int i;

    for (i = 0; i < m; ++i) {
        grid[i] = (int *)malloc(n * sizeof(int));
        if (!grid[i]) {
            free(grid);
            goto error;
        }
        memset(grid[i], 0, n);
    }

    for (i = 0; i < guards_size; ++i) {
        grid[guards[i][0]][guards[i][1]] = GUARD;
    }

    for (i = 0; i < walls_size; ++i) {
        grid[walls[i][0]][walls[i][1]] = WALL;
    }

    int guarded = 0;
    int dirs[4][2] = {{0, 1}, {1, 0}, {0, -1}, {-1, 0}};

    for (i = 0; i < guards_size; ++i) {
        for (int j = 0; j < 4; ++j) {
            int r = guards[i][0] + dirs[j][0];
            int c = guards[i][1] + dirs[j][1];
            while (r >= 0 && r < m && c >= 0 && c < n && grid[r][c] != GUARD &&
                   grid[r][c] != WALL) {
                if (grid[r][c] != GUARDED) {
                    guarded++;
                }
                grid[r][c] = GUARDED;
                r += dirs[j][0];
                c += dirs[j][1];
            }
        }
    }

    for (i = 0; i < m; ++i) {
        free(grid[i]);
        grid[i] = NULL;
    }
    free(grid);

    return (m * n) - (guards_size + guarded + walls_size);

error:
    fprintf(stderr, "Failed to allocate memory");
    return -1;
}

int main(void)
{
    int m = 4;
    int n = 6;

    int guard1[] = {0, 0};
    int guard2[] = {1, 1};
    int guard3[] = {2, 3};
    int *guards[] = {guard1, guard2, guard3};

    int wall1[] = {0, 1};
    int wall2[] = {2, 2};
    int wall3[] = {1, 4};
    int *walls[] = {wall1, wall2, wall3};

    int expected = 7;
    int actual = count_unguarded(m, n, guards, 3, NULL, walls, 3, NULL);

    if (actual != expected) {
        printf("Expected %d, but got %d\n", expected, actual);
        return EXIT_FAILURE;
    }

    printf("Correct %d\n", actual);
    return EXIT_SUCCESS;
}
