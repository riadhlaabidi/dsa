#include <stdio.h>
#include <stdlib.h>

int can_sort_array(int *nums, int nums_size)
{

    for (int i = 0; i < nums_size; ++i) {
        for (int j = 0; j < nums_size - i - 1; ++j) {
            if (nums[j] <= nums[j + 1]) {
                continue;
            } else {
                if (__builtin_popcount(nums[j]) ==
                    __builtin_popcount(nums[j + 1])) {
                    nums[j] = nums[j] ^ nums[j + 1];
                    nums[j + 1] = nums[j] ^ nums[j + 1];
                    nums[j] = nums[j] ^ nums[j + 1];
                } else {
                    return 0;
                }
            }
        }
    }

    return 1;
}

int main(void)
{

    int nums[5] = {8, 4, 2, 30, 15};

    int expected = 1;
    int actual = can_sort_array(nums, 5);

    if (actual != expected) {
        printf("Expected %d, but got %d\n", expected, actual);
        return EXIT_FAILURE;
    }

    printf("Correct %d\n", actual);

    return EXIT_SUCCESS;
}
