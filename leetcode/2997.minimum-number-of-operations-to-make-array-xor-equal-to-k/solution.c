#include <assert.h>
#include <stdio.h>

int min_operations(int *nums, int nums_size, int k)
{
    int bits_diff = 0;

    for (int i = 0; i < nums_size; ++i) {
        bits_diff ^= nums[i];
    }

    bits_diff ^= k;
    int count = 0;

    while (bits_diff) {
        count += bits_diff & 1;
        bits_diff >>= 1;
    }

    return count;
}

int main(void)
{
    int nums[] = {2, 1, 3, 4};
    int k = 1;
    int expected = 2;

    int actual = min_operations(nums, 4, k);

    assert(actual == expected);
    printf("Correct: %d", actual);

    return 0;
}
