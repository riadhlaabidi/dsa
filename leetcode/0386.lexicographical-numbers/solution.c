#include <assert.h>
#include <stdio.h>
#include <stdlib.h>

#include "../../ds/dynamic_array.h"

typedef struct {
    size_t capacity;
    size_t count;
    int *elements;
} Array;

int *lexical_order(int n, int *return_size)
{
    Array arr = {0};
    int current = 1;

    for (int i = 0; i < n; i++) {
        darray_append(&arr, current);

        if (current * 10 <= n) {
            current *= 10;
        } else {
            while (current % 10 == 9 || current >= n) {
                current /= 10;
            }
            current++;
        }
    }

    *return_size = arr.count;

    return arr.elements;
}

int main(void)
{
    int n = 13;
    int expected[] = {1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9};

    int expected_rsize = 13;
    int actual_rsize = 0;
    int *actual = lexical_order(n, &actual_rsize);

    assert(actual != NULL);
    assert(actual_rsize == expected_rsize);

    printf("[");
    for (int i = 0; i < expected_rsize; i++) {
        printf("%d", actual[i]);
        if (i < expected_rsize - 1) {
            printf(", ");
        }
        if (actual[i] != expected[i]) {
            printf("\nExpected %d at index %d, but got %d\n", expected[i], i,
                   actual[i]);
            return EXIT_FAILURE;
        }
    }
    printf("]\n");
    free(actual);
    printf("Correct\n");
    return EXIT_SUCCESS;
}
