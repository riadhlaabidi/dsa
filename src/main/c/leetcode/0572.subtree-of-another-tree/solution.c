struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

int walk(struct TreeNode *root, struct TreeNode *sub)
{
    if (!root && !sub) {
        return 1;
    }

    if (!root || !sub) {
        return 0;
    }

    if (root->val == sub->val) {
        return walk(root->left, sub->left) && walk(root->right, sub->right);
    }

    return 0;
}

int isSubtree(struct TreeNode *root, struct TreeNode *subRoot)
{
    if (!root) {
        return 0;
    }

    return walk(root, subRoot) || isSubtree(root->left, subRoot) ||
           isSubtree(root->right, subRoot);
}
