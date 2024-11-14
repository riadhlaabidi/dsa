#include <limits.h>
#include <stdio.h>
#include <stdlib.h>

int can_distribute(int k, int n, int *quantities, int quantities_size)
{
    int stores = 0;

    for (int i = 0; i < quantities_size; ++i) {
        stores += (quantities[i] / k) + (1 && (quantities[i] % k));

        if (stores > n) {
            return 0;
        }
    }

    return 1;
}

int minimized_maximum(int n, int *quantities, int quantities_size)
{

    int low = 1;
    int high = 0;

    for (int i = 0; i < quantities_size; ++i) {
        if (quantities[i] > high) {
            high = quantities[i];
        }
    }

    while (low < high) {
        int mid = (low + high) / 2;

        if (can_distribute(mid, n, quantities, quantities_size)) {
            high = mid;
        } else {
            low = mid + 1;
        }
    }

    return low;
}

int main(void)
{
    int quantities[] = {15, 10, 10};
    int n = 7;

    int expected = 5;
    int actual = minimized_maximum(n, quantities, 3);

    if (actual != expected) {
        printf("Expected %d, but got %d\n", expected, actual);
        return EXIT_FAILURE;
    }

    printf("Correct %d\n", actual);
    return EXIT_SUCCESS;
}
