#include <stdio.h>
#include <string.h>

int find_smallest_integer(int *nums, int nums_size, int value)
{
    int groups[value];
    memset(groups, 0, sizeof(groups));

    for (int i = 0; i < nums_size; i++) {
        int remainder = (nums[i] % value + value) % value;
        groups[remainder]++;
    }

    int mex = 0;
    while (groups[mex % value] > 0) {
        groups[mex % value]--;
        mex++;
    }

    return mex;
}

int main(void)
{
    int nums[] = {1, -10, 7, 13, 6, 8};
    int expected = 4;
    int actual = find_smallest_integer(nums, 6, 5);

    if (actual != expected) {
        fprintf(stderr, "Expected %d, but got %d\n", expected, actual);
        return 1;
    }

    printf("Correct\n");
    return 0;
}
