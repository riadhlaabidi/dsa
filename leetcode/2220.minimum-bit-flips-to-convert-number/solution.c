#include <assert.h>
#include <stdio.h>

int min_bit_flips(int start, int goal)
{
    int bits_to_flip = start ^ goal;
    int ans = 0;

    while (bits_to_flip) {
        ans += bits_to_flip & 1;
        bits_to_flip >>= 1;
    }

    return ans;
}

int main(void)
{
    int start = 10;
    int goal = 7;

    int expected = 3;
    int actual = min_bit_flips(start, goal);

    assert(actual == expected);

    printf("Correct %d", actual);

    return 0;
}
