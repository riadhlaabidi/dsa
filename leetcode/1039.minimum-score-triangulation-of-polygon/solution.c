#include <limits.h>
#include <stdio.h>
#include <stdlib.h>

#define UNCOMPUTED -1

int triangulate(int *values, int *dp, int row_offset, int start, int end)
{
    if (end - start < 2) {
        return 0;
    }

    int i = start * row_offset + end;
    if (dp[i] != UNCOMPUTED) {
        return dp[i];
    }

    int minScore = INT_MAX;
    for (int k = start + 1; k < end; k++) {
        int weight = values[start] * values[k] * values[end];
        int left_score = triangulate(values, dp, row_offset, start, k);
        int right_score = triangulate(values, dp, row_offset, k, end);
        int score = weight + left_score + right_score;

        if (score < minScore) {
            minScore = score;
        }
    }

    dp[i] = minScore;
    return minScore;
}

int min_score_triangulation(int *values, int values_size)
{
    int n = values_size * values_size;

    int *dp = malloc(n * sizeof(int));
    if (!dp) {
        fprintf(stderr, "Failed to allocate memory");
        return -1;
    }

    for (int i = 0; i < n; i++) {
        dp[i] = UNCOMPUTED;
    }

    int min_score = triangulate(values, dp, values_size, 0, values_size - 1);

    free(dp);
    return min_score;
}

int main(void)
{
    int values[] = {3, 7, 4, 5};
    int expected = 144;
    int actual = min_score_triangulation(values, 4);

    if (actual != expected) {
        fprintf(stderr, "Expected %d, but got %d\n", expected, actual);
        return 1;
    }

    printf("Correct\n");
    return 0;
}
