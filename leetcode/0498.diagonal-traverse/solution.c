#include <stdio.h>
#include <stdlib.h>

int *find_diagonal_order(int **mat, int mat_size, int *mat_col_size,
                         int *return_size)
{
    if (mat == NULL || mat_size == 0) {
        *return_size = 0;
        return NULL;
    }

    *return_size = mat_size * (*mat_col_size);
    int *d_order = (int *)malloc((*return_size) * sizeof(int));
    if (!d_order) {
        fprintf(stderr, "failed to allocate memory\n");
        return NULL;
    }

    int dir = 1;
    int row = 0;
    int col = 0;
    int index = 0;

    while (row < mat_size && col < *mat_col_size) {
        d_order[index++] = mat[row][col];

        // going up   -> row - 1, col + 1;
        // going down -> row + 1, col - 1;
        int new_row = row + (dir == 1 ? -1 : 1);
        int new_col = col + (dir == 1 ? 1 : -1);

        if (new_row < 0 || new_row == mat_size || new_col < 0 ||
            new_col == *mat_col_size) {
            // have to turn
            if (dir == 1) {
                row += (col == *mat_col_size - 1 ? 1 : 0);
                col += (col < *mat_col_size - 1 ? 1 : 0);
            } else {
                col += (row == mat_size - 1 ? 1 : 0);
                row += (row < mat_size - 1 ? 1 : 0);
            }

            dir = 1 - dir;
        } else {
            row = new_row;
            col = new_col;
        }
    }

    return d_order;
}

int main(void)
{
    int status = 0;

    int mat[][3] = {
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    };

    int *mat_p[] = {mat[0], mat[1], mat[2]};

    int expected[] = {1, 2, 4, 7, 5, 3, 6, 8, 9};
    int expected_return_size = 9;

    int mat_size = 3;
    int return_size = -1;
    int *actual = find_diagonal_order(mat_p, mat_size, &mat_size, &return_size);

    if (return_size != expected_return_size) {
        fprintf(stderr, "Expected return size %d, but got %d",
                expected_return_size, return_size);
        goto defer;
    }

    for (int i = 0; i < return_size; i++) {
        if (actual[i] != expected[i]) {
            fprintf(stderr, "d_order[%d]: expected = %d, actual = %d\n", i,
                    expected[i], actual[i]);
            goto defer;
        }
    }

    status = 1;
    printf("Correct\n");

defer:
    free(actual);
    return status;
}
