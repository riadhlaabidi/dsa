#include <assert.h>
#include <stdlib.h>

int maximum_69_number(int num)
{
    int multiplier = 1;
    int diff_to_max = 0;
    int x = num;

    while (x != 0) {
        int digit = x % 10;
        x /= 10;
        if (digit != 9) {
            diff_to_max = 3 * multiplier;
        }
        multiplier *= 10;
    }

    return num + diff_to_max;
}

int main(void)
{
    assert(maximum_69_number(6) == 9);
    assert(maximum_69_number(9) == 9);
    assert(maximum_69_number(9669) == 9969);
    assert(maximum_69_number(9996) == 9999);
    assert(maximum_69_number(996999999) == 999999999);
    assert(maximum_69_number(699999999) == 999999999);
    return EXIT_SUCCESS;
}
