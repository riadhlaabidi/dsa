#include <stdlib.h>

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

void traverse(struct TreeNode *root, int *values, int *returnSize)
{
    if (!root) {
        return;
    }

    traverse(root->left, values, returnSize);
    traverse(root->right, values, returnSize);

    values[(*returnSize)++] = root->val;
}

int *postorderTraversal(struct TreeNode *root, int *returnSize)
{
    *returnSize = 0;
    int *values = (int *)malloc(100 * sizeof(int));
    traverse(root, values, returnSize);
    return values;
}

int main(void) { return EXIT_SUCCESS; }
