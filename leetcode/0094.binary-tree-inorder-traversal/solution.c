#include <stdlib.h>

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

void walk(struct TreeNode *root, int *values, int *return_size)
{
    if (!root) {
        return;
    }

    walk(root->left, values, return_size);
    values[(*return_size)++] = root->val;
    walk(root->right, values, return_size);
}

int *inorderTraversal(struct TreeNode *root, int *return_size)
{
    int *values = (int *)malloc(100 * sizeof(int));
    *return_size = 0;
    walk(root, values, return_size);

    return values;
}

int main(void) { return EXIT_SUCCESS; }
