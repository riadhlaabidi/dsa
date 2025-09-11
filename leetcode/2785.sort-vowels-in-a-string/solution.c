#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_ASCII_DIFF 53

int is_vowel(char c)
{
    return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
           c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U';
}

char *sort_vowels(char *s)
{
    int vowels[MAX_ASCII_DIFF] = {0};

    for (int i = 0; s[i] != '\0'; i++) {
        int c = s[i];
        if (is_vowel(c)) {
            vowels[c - 'A']++;
        }
    }

    size_t n = strlen(s);
    char *result = malloc(n + 1);
    if (!result) {
        fprintf(stderr, "Failed to allocate memory");
        return NULL;
    }

    int j = 0;
    for (size_t i = 0; i < n; i++) {
        if (is_vowel(s[i])) {
            while (vowels[j] == 0) {
                j++;
            }
            result[i] = 'A' + j;
            vowels[j]--;
        } else {
            result[i] = s[i];
        }
    }

    result[n] = '\0';
    return result;
}

int main(void)
{
    char *s = "lEetcOde";
    char *expected = "lEOtcede";
    char *actual = sort_vowels(s);

    if (!actual) {
        return 1;
    }

    if (strcmp(actual, expected) != 0) {
        printf("Expected %s, but got %s\n", expected, actual);
        return 1;
    }

    free(actual);
    printf("Correct\n");
    return 0;
}
