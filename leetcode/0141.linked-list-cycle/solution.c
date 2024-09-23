#include "../../ds/sll/sll.h"
#include <assert.h>
#include <stdio.h>

int has_cycle(ListNode *head)
{
    ListNode *slow = head;
    ListNode *fast = head;

    while (slow && fast->next) {
        slow = slow->next;
        fast = fast->next->next;
        if (fast == slow)
            return 1;
    }
    return 0;
}

int main(void)
{
    ListNode head = {.val = 3, .next = NULL};

    int expected = 0;
    int actual = has_cycle(&head);

    assert(actual == expected);

    printf("Correct %d\n", actual);

    head.next = &(ListNode){.val = 2, .next = NULL};
    head.next->next = &(ListNode){.val = 0, .next = NULL};
    head.next->next->next = &(ListNode){.val = 4, .next = NULL};
    head.next->next->next->next = head.next;

    expected = 1;
    actual = has_cycle(&head);

    assert(actual == expected);

    printf("Correct %d", actual);

    return 0;
}
