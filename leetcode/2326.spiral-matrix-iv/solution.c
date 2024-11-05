#include <stdlib.h>

struct ListNode {
    int val;
    struct ListNode *next;
};

int **spiralMatrix(int m, int n, struct ListNode *head, int *returnSize,
                   int **returnColumnSizes)
{

    int **res = (int **)malloc(m * sizeof(int *));
    *returnSize = m;
    *returnColumnSizes = (int *)malloc(m * sizeof(int));

    for (int i = 0; i < m; ++i) {
        res[i] = (int *)malloc(n * sizeof(int));
        (*returnColumnSizes)[i] = n;

        for (int j = 0; j < n; ++j) {
            res[i][j] = -1;
        }
    }

    int r = 0;
    int c = 0;
    int dr = 0;
    int dc = 1;

    while (head) {
        res[r][c] = head->val;
        if (r + dr == m || c + dc == n || c + dc < 0 ||
            res[r + dr][c + dc] != -1) {
            // change direction clockwise
            int tmp = dr;
            dr = dc;
            dc = -tmp;
        }
        r += dr;
        c += dc;
        head = head->next;
    }

    return res;
}

int main(void) { return EXIT_SUCCESS; }
