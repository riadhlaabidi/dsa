#include <stdio.h>

int max_increasing_subarrays(int *nums, int numsSize)
{
    int prev_count = 0;
    int ans = 0;
    int count = 1;

    for (int i = 1; i < numsSize; i++) {
        if (nums[i] > nums[i - 1]) {
            count++;
        } else {
            prev_count = count;
            count = 1;
        }

        int min_counts = prev_count < count ? prev_count : count;
        if (min_counts > ans) {
            ans = min_counts;
        }

        int half = count >> 1;
        if (half > ans) {
            ans = half;
        }
    }

    return ans;
}

int main(void)
{
    int nums[] = {2, 5, 7, 8, 9, 2, 3, 4, 3, 1};
    int actual = max_increasing_subarrays(nums, 10);
    int expected = 3;

    if (actual != expected) {
        fprintf(stderr, "Expected %d, but got %d\n", expected, actual);
        return 1;
    }

    int nums_[] = {1, 2, 3, 4, 4, 4, 4, 5, 6, 7};
    actual = max_increasing_subarrays(nums_, 10);
    expected = 2;

    if (actual != expected) {
        fprintf(stderr, "Expected %d, but got %d\n", expected, actual);
        return 1;
    }

    printf("Correct\n");
    return 0;
}
