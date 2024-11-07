#include <stdio.h>
#include <stdlib.h>

int largest_combination_size(int *candidates, int candidates_size)
{
    int max = 0;

    for (int i = 0; i < 24; ++i) {
        int count = 0;

        for (int j = 0; j < candidates_size; ++j) {
            if ((candidates[j] & (1 << i)) != 0)
                count++;
        }

        if (count > max)
            max = count;
    }

    return max;
}

int main(void)
{
    int candidates[7] = {16, 17, 71, 62, 12, 24, 14};

    int expected = 4;
    int actual = largest_combination_size(candidates, 7);

    if (actual != expected) {
        printf("Expected %d, but got %d\n", expected, actual);
        return EXIT_FAILURE;
    }

    printf("Correct %d\n", actual);

    return EXIT_SUCCESS;
}
