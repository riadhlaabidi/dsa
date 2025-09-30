#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int triangular_sum(int *nums, int nums_size)
{
    if (nums_size == 1) {
        return nums[0];
    }

    int *res = malloc(nums_size * sizeof(int));
    memcpy(res, nums, nums_size * sizeof(int));

    int n = nums_size;

    while (n != 1) {
        for (int i = 0; i < n - 1; i++) {
            res[i] = (res[i] + res[i + 1]) % 10;
        }
        n--;
    }

    int sum = res[0];
    free(res);
    return sum;
}

int main(void)
{
    int nums[] = {1, 2, 3, 4, 5};
    int expected = 8;
    int actual = triangular_sum(nums, 5);

    if (actual != expected) {
        fprintf(stderr, "Expected %d, but got %d\n", expected, actual);
        return 1;
    }

    printf("Correct\n");
    return 0;
}
