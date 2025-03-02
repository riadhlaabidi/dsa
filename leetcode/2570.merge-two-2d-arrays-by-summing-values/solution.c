#include <assert.h>
#include <stdio.h>
#include <stdlib.h>

#define MAX_ID 1001

int **mergeArrays(int **nums1, int nums1Size, int **nums2, int nums2Size,
                  int *returnSize, int **returnColumnSizes)
{
    int sums[MAX_ID] = {0};

    for (int i = 0; i < nums1Size; i++) {
        int id = nums1[i][0];
        sums[id] += nums1[i][1];
    }

    int merged_size = nums1Size;

    for (int i = 0; i < nums2Size; i++) {
        int id = nums2[i][0];
        if (sums[id] == 0) {
            merged_size++;
        }
        sums[id] += nums2[i][1];
    }

    int **merged = (int **)malloc(merged_size * sizeof(int *));
    if (!merged) {
        fprintf(stderr, "Error allocating memory");
        return NULL;
    }

    *returnColumnSizes = (int *)malloc(merged_size * sizeof(int));
    if (!*returnColumnSizes) {
        fprintf(stderr, "Error allocating memory");
        free(merged);
        return NULL;
    }

    int j = 0;
    for (int i = 1; i < MAX_ID; i++) {
        if (sums[i]) {
            merged[j] = (int *)malloc(2 * sizeof(int));
            merged[j][0] = i;
            merged[j][1] = sums[i];
            (*returnColumnSizes)[j] = 2;
            j++;
        }
    }

    *returnSize = merged_size;
    return merged;
}

int main(void)
{
    int a1[] = {1, 2};
    int a2[] = {2, 3};
    int a3[] = {4, 5};
    int *nums1[] = {a1, a2, a3};

    int b1[] = {1, 4};
    int b2[] = {3, 2};
    int b3[] = {4, 1};
    int *nums2[] = {b1, b2, b3};

    int expected[4][2] = {{1, 6}, {2, 3}, {3, 2}, {4, 6}};

    int returnSize = 0;

    int **returnColumnSizes = (int **)malloc(sizeof(int *));
    if (!returnColumnSizes) {
        fprintf(stderr, "Error allocating memory");
        return EXIT_FAILURE;
    }

    int **actual =
        mergeArrays(nums1, 3, nums2, 3, &returnSize, returnColumnSizes);

    assert(actual != NULL);
    assert(returnSize == 4);

    for (int i = 0; i < returnSize; i++) {
        assert((*returnColumnSizes)[i] == 2);
        assert(expected[i][0] == actual[i][0]);
        assert(expected[i][1] == actual[i][1]);
        free(actual[i]);
    }

    printf("Correct\n");

    free(actual);
    free(*returnColumnSizes);
    free(returnColumnSizes);

    return EXIT_SUCCESS;
}
