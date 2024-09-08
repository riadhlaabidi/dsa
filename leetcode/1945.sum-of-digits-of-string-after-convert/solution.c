#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int get_lucky(char *s, int k)
{
    int n = strlen(s);
    int buff_size = n * 2 + 1;
    char *str_num_repr = (char *)malloc(buff_size * sizeof(char));

    int length = 0;

    for (int i = 0; i < n; ++i) {
        int pos = s[i] - 'a' + 1;

        int cx = snprintf(str_num_repr + length, buff_size - length, "%d", pos);

        if (cx < 0 || cx >= buff_size) {
            fprintf(stderr, "Failed to write to buffer");
            free(str_num_repr);
            exit(EXIT_FAILURE);
        }

        length += cx;
    }

    int sum;
    for (int i = 0; i < k; ++i) {
        sum = 0;

        for (int j = 0; j < length; ++j) {
            sum += (str_num_repr[j] - '0');
        }

        length = snprintf(str_num_repr, buff_size, "%d", sum);

        if (length < 0 || length >= buff_size) {
            fprintf(stderr, "Failed to write to buffer");
            free(str_num_repr);
            exit(EXIT_FAILURE);
        }
    }

    free(str_num_repr);
    return sum;
}
