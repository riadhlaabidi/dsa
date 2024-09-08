#include <stdlib.h>
/**
 * Definition for a binary tree node.
 */
/**
 * Note: The returned array must be malloced, assume caller calls free().
 */

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

void backtrack(struct TreeNode* root, int* returnSize, char** paths, int* curr_path, int c_p_size) 
{
    if (root->left == NULL && root->right == NULL) {
	
    }
}

char** binaryTreePaths(struct TreeNode* root, int* returnSize) {
    char** res = (char **) malloc(100 * sizeof(char *));

    *returnSize = 0;
    int* curr_path = (int *) malloc(200 * sizeof(int));
    int c_p_size = 0;
    backtrack(root, returnSize, res, curr_path, c_p_size);

    return NULL; 
} 
