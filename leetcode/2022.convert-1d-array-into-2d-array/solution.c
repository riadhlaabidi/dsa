#include <stdlib.h>

int **construct_2d_array(int *original, int original_size, int m, int n,
                         int *return_size, int **return_column_sizes)
{

    if (n * m != original_size) {
        *return_size = 0;
        return NULL;
    }

    int **ans = (int **)malloc(m * sizeof(int *));
    int *col_sizes = (int *)malloc(m * sizeof(int));

    for (int i = 0; i < m; ++i) {
        ans[i] = (int *)malloc(n * sizeof(int));
        col_sizes[i] = n;
    }

    for (int i = 0; i < original_size; ++i) {
        ans[i / n][i % n] = original[i];
    }

    *return_size = m;
    *return_column_sizes = col_sizes;

    return ans;
}

int main(void) { return EXIT_SUCCESS; }
