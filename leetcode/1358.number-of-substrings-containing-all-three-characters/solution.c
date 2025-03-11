#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int number_of_substrings(char *s)
{
    int n = strlen(s);
    int left = 0;
    int right = 0;
    int chars[3] = {0};
    int subs = 0;

    while (right < n) {
        chars[s[right] - 'a']++;

        while (chars[0] > 0 && chars[1] > 0 && chars[2] > 0) {
            // all substrings from current window to n are valid ones
            subs += n - right;

            chars[s[left] - 'a']--;
            left++;
        }

        right++;
    }

    return subs;
}

int main(void)
{
    char *s = "abcabc";
    int expected = 10;
    int actual = number_of_substrings(s);
    printf("Expected = %d\nActual = %d\n", expected, actual);
    assert(actual == expected);
    return EXIT_SUCCESS;
}
