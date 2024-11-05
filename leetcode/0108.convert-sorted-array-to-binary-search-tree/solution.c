#include <stdlib.h>

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

struct TreeNode *construct_bst(int *nums, int start, int end)
{
    if (start >= end) {
        return NULL;
    }

    struct TreeNode *root = (struct TreeNode *)malloc(sizeof(struct TreeNode));
    int middle = (start + end) / 2;
    root->val = nums[middle];
    root->left = construct_bst(nums, start, middle);
    root->right = construct_bst(nums, middle + 1, end);

    return root;
}

struct TreeNode *sortedArrayToBST(int *nums, int numsSize)
{
    return construct_bst(nums, 0, numsSize);
}

int main(void) { return EXIT_SUCCESS; }
