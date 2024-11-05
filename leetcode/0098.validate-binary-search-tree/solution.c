#include <limits.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

int is_valid_BST_helper(struct TreeNode *root, long min, long max)
{
    if (!root) {
        return 1;
    }

    if (root->val >= max || root->val <= min) {
        return 0;
    }

    return is_valid_BST_helper(root->left, min, root->val) &&
           is_valid_BST_helper(root->right, root->val, max);
}

int is_valid_BST(struct TreeNode *root)
{
    return is_valid_BST_helper(root, LONG_MIN, LONG_MAX);
}

int main(void)
{
    struct TreeNode root = {5, NULL, NULL};

    root.left = &((struct TreeNode){1, NULL, NULL});
    root.right = &((struct TreeNode){4, NULL, NULL});
    root.right->left = &((struct TreeNode){3, NULL, NULL});
    root.right->right = &((struct TreeNode){6, NULL, NULL});

    int expected = 0;
    int actual = is_valid_BST(&root);

    if (actual != expected) {
        printf("Expected %d, but got %d\n", expected, actual);
        return EXIT_FAILURE;
    }

    printf("Correct %d\n", actual);

    return EXIT_SUCCESS;
}
