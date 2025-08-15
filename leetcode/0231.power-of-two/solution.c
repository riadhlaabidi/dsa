#include <assert.h>
#include <limits.h>
#include <stdlib.h>

int is_power_of_two(int n) { return n > 0 && (n & (n - 1)) == 0; }

int main(void)
{
    // no power of 2 is negative
    assert(is_power_of_two(-16) == 0);
    assert(is_power_of_two(INT_MIN) == 0);

    // powers of two have a single set bit
    assert(is_power_of_two(1) == 1);
    assert(is_power_of_two(0b0010) == 1);

    // if multiple set bits => not a power of two
    assert(is_power_of_two(INT_MAX) == 0); // 0b011111...
    assert(is_power_of_two(0b0110) == 0);

    return EXIT_SUCCESS;
}
