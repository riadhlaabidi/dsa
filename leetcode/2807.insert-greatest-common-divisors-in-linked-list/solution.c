#include <assert.h>
#include <stdio.h>
#include <stdlib.h>

struct ListNode {
    int val;
    struct ListNode *next;
};

int gcd(int a, int b)
{
    while (b != 0) {
        int tmp = a % b;
        a = b;
        b = tmp;
    }

    return a;
}

struct ListNode *insert_gcds(struct ListNode *head)
{
    struct ListNode *tmp = head;

    while (tmp->next) {
        struct ListNode *new_node =
            (struct ListNode *)malloc(sizeof(struct ListNode));

        if (!new_node) {
            fprintf(stderr, "Memory allocation failed");
            return NULL;
        }

        new_node->val = gcd(tmp->val, tmp->next->val);
        new_node->next = tmp->next;
        tmp->next = new_node;
        tmp = new_node->next;
    }

    return head;
}

int main(void)
{
    int arr[] = {18, 6, 10, 3};
    struct ListNode *head = (struct ListNode *)malloc(sizeof(struct ListNode));

    if (!head) {
        fprintf(stderr, "Memory allocation failed");
        return EXIT_FAILURE;
    }

    head->val = arr[0];
    head->next = NULL;

    struct ListNode *tmp = head;

    for (int i = 1; i < 4; ++i) {
        struct ListNode *new_node =
            (struct ListNode *)malloc(sizeof(struct ListNode));

        if (!new_node) {
            fprintf(stderr, "Memory allocation failed");
            while (head) {
                struct ListNode *d = head;
                head = head->next;
                free(d);
            }
            return EXIT_FAILURE;
        }

        new_node->val = arr[i];
        new_node->next = NULL;
        tmp->next = new_node;
        tmp = tmp->next;
    }

    int expected[] = {18, 6, 6, 2, 10, 1, 3};

    struct ListNode *ans = insert_gcds(head);
    int i = 0;

    while (ans) {
        assert(ans->val == expected[i++]);
        struct ListNode *d = ans;
        ans = ans->next;
        free(d);
    }

    printf("Correct");

    return EXIT_SUCCESS;
}
