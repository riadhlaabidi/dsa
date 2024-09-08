#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int compare(const void *a, const void *b) {
    int arg1 = *(const int *)a;
    int arg2 = *(const int *)b;
    if (arg1 > arg2) return 1;
    if (arg1 < arg2) return -1;
    return 0;
}

void backtrack(int idx, int *candidates, int cs, int *curr, int curr_s,
               int total_left, int **combinations, int *returnSize,
               int **returnColumnSizes) {
    if (total_left < 0) {
	return;
    } else {
	if (total_left == 0) {
	    int *combi = (int *) malloc(curr_s * sizeof(int));
	    memcpy(combi, curr, curr_s * sizeof(int));
	    combinations[*returnSize] = combi;
	    (*returnColumnSizes)[*returnSize] = curr_s;
	    (*returnSize)++;
	} else {
	  for (int i = idx; i < cs && total_left >= candidates[i]; ++i) {
		if (i > idx && candidates[i] == candidates[i - 1]) continue;
	        curr[curr_s++] = candidates[i];
		backtrack(i + 1, candidates, cs, curr, curr_s,
			  total_left - candidates[i], combinations, returnSize,
			  returnColumnSizes);
	        curr_s--;
	  }
	}
    }
}

int **combinationSum2(int *candidates, int candidatesSize, int target,
                      int *returnSize, int **returnColumnSizes) {
    int **combinations = (int **) malloc(200 * sizeof(int *));
    *returnColumnSizes = (int *) malloc(200 * sizeof(int));
    *returnSize = 0;

    qsort(candidates, candidatesSize, sizeof(int), compare);

    int *curr = (int *) malloc(candidatesSize * sizeof(int));
    int curr_s = 0;
    backtrack(0, candidates, candidatesSize, curr, curr_s, target, combinations,
	      returnSize, returnColumnSizes);

    free(curr);
    return combinations;
}

int main(void) {
    int arr[7] = {10, 1, 2, 7, 6, 1, 5};
    int target = 8;
    int *returnSize = (int *) malloc(sizeof(int));
    int **returnColumnsSize = NULL;

    int **combs = combinationSum2(arr, 7, target, returnSize, returnColumnsSize);

    for (int i = 0; i < *returnSize; ++i) {
	for (int j = 0; j < *(returnColumnsSize[i]); ++j) {
	    printf("%d ", combs[i][j]);
	}
	printf("\n");
    }

    free(returnSize);
    free(returnColumnsSize);
    free(combs);

    return 0;
}
