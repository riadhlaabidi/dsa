#include <stdio.h>

int minimum_total(int **triangle, int triangle_size, int *triangle_col_size)
{
    for (int i = triangle_size - 2; i >= 0; i--) {
        for (int j = 0; j < triangle_col_size[i]; j++) {
            int left = triangle[i + 1][j];
            int right = triangle[i + 1][j + 1];
            int min = left < right ? left : right;
            triangle[i][j] += min;
        }
    }

    return triangle[0][0];
}

int main(void)
{
    int triangle_col_size[] = {1, 2, 3, 4};
    int triangles[][4] = {{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}};
    int *triangle[4] = {
        triangles[0],
        triangles[1],
        triangles[2],
        triangles[3],
    };

    int expected = 11;
    int actual = minimum_total(triangle, 4, triangle_col_size);

    if (actual != expected) {
        printf("Expected %d, but got %d\n", expected, actual);
        return 1;
    }

    printf("Correct\n");
    return 0;
}
