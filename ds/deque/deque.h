#ifndef INCLUDE_DEQUE_DEQUE_H_
#define INCLUDE_DEQUE_DEQUE_H_

#include <stddef.h>
#include <stdio.h>

typedef struct node {
    struct node *next;
    struct node *prev;
} Node;

typedef struct {
    Node *head;
    Node *tail;
} Deque;

#define DEQUE_DATA(NODE, STRUCT, MEMBER)                                       \
    ((STRUCT *)((uint8_t *)(NODE) - offsetof(STRUCT, MEMBER)))

static inline void DEQUE_INIT(Deque *const d)
{
    d->head = NULL;
    d->tail = NULL;
}

static inline void DEQUE_ADD_FRONT(Deque *const d, Node *const node)
{
    node->prev = NULL;

    if (d->head == NULL) {
        node->next = NULL;
        d->tail = node;
    } else {
        d->head->prev = node;
        node->next = d->head;
    }

    d->head = node;
}

static inline void DEQUE_ADD_BACK(Deque *const d, Node *const node)
{
    node->next = NULL;

    if (d->tail == NULL) {
        node->prev = NULL;
        d->head = node;
    } else {
        node->prev = d->tail;
        d->tail->next = node;
    }

    d->tail = node;
}

static inline int DEQUE_IS_EMPTY(const Deque *const d)
{
    return d->head == NULL && d->tail == NULL;
}

static inline void DEQUE_POP_FRONT(Deque *const d)
{
    if (DEQUE_IS_EMPTY(d)) {
        fprintf(stderr, "Deque is empty\n");
        return;
    }

    d->head = d->head->next;

    if (d->head == NULL) {
        d->tail = NULL;
    } else {
        d->head->prev = NULL;
    }
}

static inline void DEQUE_POP_BACK(Deque *const d)
{
    if (DEQUE_IS_EMPTY(d)) {
        fprintf(stderr, "Deque is empty\n");
        return;
    }

    d->tail = d->tail->prev;

    if (d->tail == NULL) {
        d->head = NULL;
    } else {
        d->tail->next = NULL;
    }
}

static inline const Node *DEQUE_GET_FRONT(const Deque *const d)
{
    if (DEQUE_IS_EMPTY(d)) {
        fprintf(stderr, "Deque is empty\n");
        return NULL;
    }

    return d->head;
}

static inline const Node *DEQUE_GET_BACK(const Deque *const d)
{
    if (DEQUE_IS_EMPTY(d)) {
        fprintf(stderr, "Deque is empty\n");
        return NULL;
    }

    return d->tail;
}

#endif // INCLUDE_DEQUE_DEQUE_H_
