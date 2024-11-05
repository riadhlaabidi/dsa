#include <stdio.h>
#include <stdlib.h>

struct ListNode {
    int val;
    struct ListNode *next;
};

struct ListNode **split_list_to_parts(struct ListNode *head, int k,
                                      int *return_size)
{
    int n = 0;
    struct ListNode *curr = head;

    while (curr) {
        n++;
        curr = curr->next;
    }

    int base = n / k;
    int remaining = n % k;

    struct ListNode **parts =
        (struct ListNode **)malloc(k * sizeof(struct ListNode *));

    if (!parts) {
        fprintf(stderr, "Memory allocation failed");
        return NULL;
    }

    curr = head;
    for (int i = 0; i < k; ++i) {
        if (!curr) {
            parts[i] = NULL;
            continue;
        }

        int part_size = base;

        if (remaining > 0) {
            part_size++;
            remaining--;
        }

        struct ListNode *prev = curr;
        parts[i] = prev;

        while (--part_size) {
            prev = prev->next;
        }

        curr = prev->next;
        prev->next = NULL;
    }

    *return_size = k;
    return parts;
}

int main(void) { return EXIT_SUCCESS; }
