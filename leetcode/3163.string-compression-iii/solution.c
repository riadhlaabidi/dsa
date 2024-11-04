#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *compress_string(char *word)
{
    char *ans = malloc(4 * 10e5);
    int n = strlen(word);
    int i = 0;
    int j = 0;

    while (i < n) {
        int curr_cnt = 1;
        char curr_char = word[i++];

        while (curr_cnt < 9 && i < n && curr_char == word[i]) {
            curr_cnt++;
            i++;
        }

        ans[j++] = '0' + curr_cnt;
        ans[j++] = curr_char;
    }

    ans[j] = '\0';

    return ans;
}

int main(void)
{
    char *word = "aaaaaaaaaaaaaabb";

    char *actual = compress_string(word);
    char *expected = "9a5a2b";

    if (strcmp(actual, expected)) {
        printf("Expected %s, but got %s", expected, actual);
        return EXIT_FAILURE;
    }

    printf("Correct %s", actual);

    return EXIT_SUCCESS;
}
