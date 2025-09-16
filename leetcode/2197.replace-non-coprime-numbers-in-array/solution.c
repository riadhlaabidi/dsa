#include <stdio.h>
#include <stdlib.h>

int gcd(int x, int y)
{
    if (y == 0) {
        return x;
    }
    return gcd(y, x % y);
}

int *replace_non_coprimes(int *nums, int nums_size, int *return_size)
{
    int *stack = (int *)malloc(nums_size * sizeof(int));
    if (!stack) {
        fprintf(stderr, "Failed to allocate memory");
        *return_size = 0;
        return NULL;
    }

    int top = -1;

    for (int i = 0; i < nums_size; i++) {
        stack[++top] = nums[i];

        int top_gcd;
        while (top + 1 >= 2 &&
               (top_gcd = gcd(stack[top], stack[top - 1])) > 1) {
            int x = stack[top--];
            int y = stack[top];
            stack[top] = x / top_gcd * y;
        }
    }

    *return_size = top + 1;
    return stack;
}

int main(void)
{
    int input[] = {6, 4, 3, 2, 7, 6, 2};
    int expected[] = {12, 7, 6};
    int return_size = -1;
    int *actual = replace_non_coprimes(input, 7, &return_size);

    if (!actual) {
        return 1;
    }

    for (int i = 0; i < return_size; i++) {
        if (actual[i] != expected[i]) {
            printf("Expected %d at index %d, but got %d\n", expected[i], i,
                   actual[i]);
            free(actual);
            return 1;
        }
    }

    printf("Correct\n");
    free(actual);
    return 0;
}
