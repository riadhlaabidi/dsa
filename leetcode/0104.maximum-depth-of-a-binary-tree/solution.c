#include <stdio.h>
#include <stdlib.h>

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

int maxDepth(struct TreeNode *root)
{
    if (root == NULL) {
        return 0;
    }

    int max_r = 1 + maxDepth(root->right);
    int max_l = 1 + maxDepth(root->left);

    if (max_r > max_l) {
        return max_r;
    }

    return max_l;
}

int main(void)
{
    struct TreeNode *root = &((struct TreeNode){3, NULL, NULL});
    root->left = &((struct TreeNode){9, NULL, NULL});
    root->right = &((struct TreeNode){20, NULL, NULL});
    root->right->left = &((struct TreeNode){15, NULL, NULL});
    root->right->right = &((struct TreeNode){7, NULL, NULL});

    int expected = 3;
    int actual = maxDepth(root);

    if (actual != expected) {
        printf("Expected %d, but got %d\n", expected, actual);
        return EXIT_FAILURE;
    }

    printf("Correct %d\n", actual);
    return EXIT_SUCCESS;
}
