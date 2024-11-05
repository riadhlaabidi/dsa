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

struct TreeNode *construct_bst(struct ListNode *start, struct ListNode *end)
{
    if (start == end) {
        return NULL;
    }

    struct ListNode *slow = start;
    struct ListNode *fast = start;

    while (fast != end && fast->next != end) {
        slow = slow->next;
        fast = fast->next->next;
    }

    struct TreeNode *root = (struct TreeNode *)malloc(sizeof(struct TreeNode));
    root->val = slow->val;
    root->left = construct_bst(start, slow);
    root->right = construct_bst(slow->next, end);

    return root;
}

struct TreeNode *sortedListToBST(struct ListNode *head)
{
    return construct_bst(head, NULL);
}

int main(void) { return EXIT_SUCCESS; }
