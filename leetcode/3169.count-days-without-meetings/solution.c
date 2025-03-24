#include <assert.h>
#include <stdlib.h>

int compare_intervals(const void *a, const void *b)
{
    const int *interval1 = *(const int **)a;
    const int *interval2 = *(const int **)b;

    int start1 = interval1[0];
    int start2 = interval2[0];

    return (start1 > start2) - (start1 < start2);
}

int countDays(int days, int **meetings, int meetings_size,
              int *meetings_col_size)
{
    int free_days = 0;
    int max = 0;

    qsort(meetings, meetings_size, sizeof(int *), compare_intervals);

    for (int i = 0; i < meetings_size; i++) {
        if (meetings[i][0] > max) {
            free_days += (meetings[i][0] - max - 1);
        }

        if (meetings[i][1] > max) {
            max = meetings[i][1];
        }
    }

    return free_days + days - max;
}

int main(void)
{
    int days = 10;

    int *meetings[3];
    meetings[0] = (int[2]){5, 7};
    meetings[1] = (int[2]){1, 3};
    meetings[2] = (int[2]){9, 10};

    int meetings_col_size = 2;

    int expected = 2;
    int actual = countDays(days, meetings, 3, &meetings_col_size);

    assert(actual == expected);

    return EXIT_SUCCESS;
}
