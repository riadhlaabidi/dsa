#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *string_hash(char *s, int k)
{
    int n = strlen(s);
    char *ans = (char *)malloc((n / k) * sizeof(char));

    int c = 0;
    int j = 0;

    for (int i = 0; i < n; ++i) {
        c += s[i] - 'a';

        if ((i + 1) % k == 0) {
            ans[j++] = (char)(c % 26 + 'a');
            c = 0;
        }
    }

    return ans;
}

int main(void)
{
    char s[] = "abcd";
    char expected[] = "bf";

    char *res = string_hash(s, 2);

    assert(!strcmp(res, expected));

    printf("result = %s", res);

    free(res);
    return 0;
}
