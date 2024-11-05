#include <stdlib.h>

struct ListNode {
    int val;
    struct ListNode *next;
};

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

int walk(struct ListNode *node, struct TreeNode *root)
{
    if (!node) {
        return 1;
    }

    if (!root) {
        return 0;
    }

    if (root->val == node->val) {
        return walk(node->next, root->left) || walk(node->next, root->right);
    }

    return 0;
}

int is_sub_path(struct ListNode *head, struct TreeNode *root)
{
    if (!root) {
        return 0;
    }

    return walk(head, root) || is_sub_path(head, root->left) ||
           is_sub_path(head, root->right);
}

int main(void) { return EXIT_SUCCESS; }
