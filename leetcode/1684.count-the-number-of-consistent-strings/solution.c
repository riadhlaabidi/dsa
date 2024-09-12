#include <assert.h>
#include <stdio.h>

int count_consistent_strings(char *allowed, char **words, int words_size)
{
    int allowed_flags[26] = {0};
    char *c = allowed;

    while (*c != '\0') {
        allowed_flags[*c - 'a'] = 1;
        c++;
    }

    int ans = 0;

    for (int i = 0; i < words_size; ++i) {
        char *ch = words[i];
        int consistent = 1;

        while (*ch != '\0') {
            if (!allowed_flags[*ch - 'a']) {
                consistent = 0;
                break;
            }
            ch++;
        }

        ans += consistent;
    }

    return ans;
}

int main(void)
{

    char allowed[] = "ab";
    char *words[] = {"ad", "bd", "aaab", "baa", "badab"};

    int expected = 2;
    int actual = count_consistent_strings(allowed, words, 5);

    assert(actual == expected);

    printf("Correct %d", actual);

    return 0;
}
