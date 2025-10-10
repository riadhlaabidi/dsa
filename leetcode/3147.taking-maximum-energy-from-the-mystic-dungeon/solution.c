#include <limits.h>
#include <stdio.h>

int maximum_energy(int *energy, int energy_size, int k)
{
    int max_energy = INT_MIN;

    for (int i = energy_size - k; i < energy_size; i++) {
        int e = 0;
        for (int j = i; j >= 0; j -= k) {
            e += energy[j];
            if (e > max_energy) {
                max_energy = e;
            }
        }
    }

    return max_energy;
}

int main(void)
{
    int energy[] = {5, 2, -10, -5, 1};

    int actual = maximum_energy(energy, 5, 3);
    int expected = 3;

    if (actual != expected) {
        fprintf(stderr, "Expected %d, but got %d\n", expected, actual);
        return 1;
    }

    int energy_[] = {-2, -3, -1};
    actual = maximum_energy(energy_, 3, 2);
    expected = -1;

    if (actual != expected) {
        fprintf(stderr, "Expected %d, but got %d\n", expected, actual);
        return 1;
    }

    printf("Correct\n");
    return 0;
}
