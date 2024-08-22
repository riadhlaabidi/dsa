#include <assert.h>
#include <stdio.h>

int find_complement(int num) 
{
    unsigned int mask = ~0;

    // Shift left mask until no overlapping bits 
    while (mask & num) {
	mask = mask << 1;
    } 

    // Xor the mask with num's complement 
    // to get rid of the non used bits set to 1
    return ~num ^ mask;
} 

int main(void)
{
    int n = 5;
    int expected = 2;
    int actual = find_complement(n);

    assert(actual == expected);

    n = 1;
    expected = 0;
    actual = find_complement(n);

    assert(actual == expected);
 
    printf("Correct");
}
