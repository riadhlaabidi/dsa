#include <stdio.h>
#include <stdlib.h>
#include <string.h>

long long max(long long a , long long b)
{
    return a > b ? a : b;
}

long long max_points(int** points, int points_size, int* points_col_size) 
{ 
    int cols = points_col_size[0];
    long long *prev = (long long *) calloc(cols, sizeof(long long));
    long long *curr = (long long *) malloc(cols * sizeof(long long));

    for (int i = 0; i < points_size; ++i) {

        long long curr_max = 0;
        for (int j = 0; j < cols; ++j) {
            curr_max = max(curr_max - 1, prev[j]);
            curr[j] = curr_max;
        }

        curr_max = 0;
        for (int j = cols - 1; j >= 0; --j) {
            curr_max = max(curr_max - 1, prev[j]);
            curr[j] = max(curr[j], curr_max) + points[i][j];
        }

        memcpy(prev, curr, cols * sizeof(long long));
    } 

    long long max_score = 0; 
    for (int i = 0; i < cols; ++i) {
        max_score = max(prev[i], max_score);
    }

    free(prev); 
    free(curr);
    return max_score;
}

int main(void) 
{
    int grid_col_size[3] = {3, 3, 3};
    int c1[3] = {1, 2, 3};
    int c2[3] = {1, 5, 1};
    int c3[3] = {3, 1, 1};
    int *grid[3] = {c1, c2, c3};

    int expected = 9;
    int ans = max_points(grid, 3, grid_col_size);

    printf("Expected: %d\n", expected);
    printf("Result: %d\n===> ", ans);

    if (ans == expected)
	printf("Correct");
    else
	printf("Mismatch");

    return 0;
}
