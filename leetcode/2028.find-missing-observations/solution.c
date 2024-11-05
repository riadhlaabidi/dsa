#include <stdlib.h>

int *missing_rolls(int *rolls, int rolls_size, int mean, int n,
                   int *return_size)
{
    int missing_sum = 0;

    for (int i = 0; i < rolls_size; ++i) {
        missing_sum += rolls[i];
    }

    missing_sum = mean * (rolls_size + n) - missing_sum;

    if (missing_sum > 6 * n || missing_sum < n) {
        *return_size = 0;
        return NULL;
    }

    int *ans = (int *)malloc(n * sizeof(int));
    int base = missing_sum / n;
    int remainder = missing_sum % n;

    for (int i = 0; i < n; ++i) {
        ans[i] = base + (i < remainder);
    }

    *return_size = n;
    return ans;
}

int main(void) { return EXIT_SUCCESS; }
