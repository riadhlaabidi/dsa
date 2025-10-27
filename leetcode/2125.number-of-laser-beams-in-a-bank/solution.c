#include <stdio.h>

int number_of_beams(char **bank, int bank_size)
{
    int prev = 0;
    int res = 0;

    for (int i = 0; i < bank_size; i++) {
        int count = 0;
        char *cell = bank[i];

        while (*cell) {
            if (*cell == '1') {
                res += prev;
                count++;
            }
            cell++;
        }

        if (count) {
            prev = count;
        }
    }

    return res;
}

int main(void)
{
    char *bank[] = {"011001", "000000", "010100", "001000"};
    int expected = 8;
    int actual = number_of_beams(bank, 4);

    if (actual != expected) {
        fprintf(stderr, "Expected %d, but got %d\n", expected, actual);
        return 1;
    }

    printf("Correct\n");
    return 0;
}
